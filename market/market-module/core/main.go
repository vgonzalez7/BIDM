package main

import (
	"context"
	"log"
	"io"
	"time"
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"math"
	"math/big"
	//"math/rand"
	"net/http"
	"os"
	"flag"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/crypto"

	accessControlContract "marketplace/core/contracts/accessContract"
	balanceContract "marketplace/core/contracts/balanceContract"
	dataContract "marketplace/core/contracts/dataContract"
	libs "marketplace/core/libs"
)

// Local definition of the struct libs.ComponentConfig
type localClient libs.ComponentConfig

func readConfigFile() map[string]interface{} {
	config := make(map[string]interface{})

	// Open the configuration file
	jsonFile, err := os.Open("./config/config.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer jsonFile.Close()

	// Parse to bytes
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Load the object in the map[string]interface{} variable
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return config

}

func initialize() (localClient, localClient) {
	// Read the configuration file
	config := readConfigFile()

	// Connect to the IPC endpoint of the Ethereum Admin node
	client, err := ethclient.Dial(config["nodePath"].(string) + "geth.ipc")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Connect to the IPC endpoint of the Ethereum client node
	buyer, err := ethclient.Dial(config["node2Path"].(string) + "geth.ipc")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Get the private key of the admin
	privKey, err := libs.GetPrivateKey(config["ethAddress"].(string),
		config["ethPassword"].(string),
		config["nodePath"].(string)+"keystore/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Get the private key of the client
	privKey2, err := libs.GetPrivateKey(config["ethAddress2"].(string),
		config["ethPassword"].(string),
		config["node2Path"].(string)+"keystore/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Initialize the data contract
	dataContract, err := dataContract.NewDataLedgerContract(common.HexToAddress(config["dataContractAddr"].(string)), client)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Set the access contract address in the data contract so the data
	// contract can interact with the  access contract
	auth := bind.NewKeyedTransactor(privKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(4000000)
	auth.GasPrice = big.NewInt(0)

	_, err = dataContract.SetAddress(auth, common.HexToAddress(config["accessContractAddr"].(string)))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Initialize the accessControlContract
	accessContract, err := accessControlContract.NewAccessControlContract(common.HexToAddress(config["accessContractAddr"].(string)), client)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Set admin public address in the access smart contract
	publicKeyECDSA := privKey.PublicKey
	publicKeyBytes := elliptic.Marshal(publicKeyECDSA.Curve, publicKeyECDSA.X, publicKeyECDSA.Y)
	publicKeyString := fmt.Sprintf("%x", publicKeyBytes)
	_, err = accessContract.AddPubKey(auth, publicKeyString)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	publicKeyECDSA2 := privKey2.PublicKey

	// Initialize the balanceContract
	balanceContract, err := balanceContract.NewBalanceContract(common.HexToAddress(config["balanceContractAddr"].(string)), client)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Set the data contract address in the balance contract
	_, err = balanceContract.SetAddress(auth, common.HexToAddress(config["dataContractAddr"].(string)))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Set the total amount of tokens in the marketplace
	_, err = balanceContract.SetTotalSupply(auth, big.NewInt(int64(config["totalSupply"].(float64))))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ipfsp := config["IPFSPath"].(string)

	localEthClient := localClient{
		client,
		common.HexToAddress(config["ethAddress"].(string)),
		privKey,
		publicKeyECDSA,
		dataContract,
		accessContract,
		balanceContract,
		ipfsp,
		config,
	}

	localEthBuyer := localClient{
		buyer,
		common.HexToAddress(config["ethAddress"].(string)),
		privKey2,
		publicKeyECDSA2,
		dataContract,
		accessContract,
		balanceContract,
		ipfsp,
		config,
	}

	return localEthClient, localEthBuyer

}

// AdminEP is the end point for the admin user. Here, the admin can send
// new tokens to clients, add new producers to the access list, ...
func (localClient *localClient) AdminEP(w http.ResponseWriter, req *http.Request) {
	// Check whether the URL is correct or not
	if req.URL.Path != "/admin" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Create a map with body of the message
	//var bodyArray bodyArray
	bodyMap := make(map[string]interface{})

	// Read the body of the message
	err := json.NewDecoder(req.Body).Decode(&bodyMap)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "401 Could not process the event: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Convert the localClient to libs.ComponentConfig
	ethClient := libs.ComponentConfig{
		localClient.EthereumClient,
		localClient.Address,
		localClient.PrivateKey,
		localClient.PublicKey,
		localClient.DataCon,
		localClient.AccessCon,
		localClient.BalanceCon,
		localClient.IPFSPath,
		localClient.GeneralConfig,
	}

	// Check if the method of the request is correct
	switch req.Method {
	case "POST":
		measurement, err := libs.AdminListener(ethClient, bodyMap)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "401 Could not process the event: "+err.Error(), http.StatusBadRequest)
			return
		}
		if measurement != nil { w.Write(measurement) } else { http.Error(w, "200 OK", http.StatusOK) }

	default:
		http.Error(w, "401 Only POST method is supported", http.StatusBadRequest)
		fmt.Fprintf(w, "Only POST method is supported")
	}
}

func compulsiveBuy(ethBuyer libs.ComponentConfig) {
	// Log file
	f, err := os.OpenFile("/home/administrator/actual_blockchain/timestampsC.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
	      log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	auth := bind.NewKeyedTransactor(ethBuyer.PrivateKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	// Prepare the filter to filter the contract events
	//topic := []byte("evtStoreInfo(bytes32, string, string)")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			common.HexToAddress(ethBuyer.GeneralConfig["dataContractAddr"].(string))},
	//	Topics: [][]common.Hash{{crypto.Keccak256Hash(topic)}},
	}

	// Make channel to receive the events
	logs := make(chan types.Log)
	sub, err := ethBuyer.EthereumClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// Process the events
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			//time.Sleep(time.Second)

			// Log file
		      f, err = os.OpenFile("/home/administrator/actual_blockchain/timestampsC.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		      if err != nil {
		          log.Fatalf("error opening file: %v", err)
	              }
	              wrt = io.MultiWriter(os.Stdout, f)
	              log.SetOutput(wrt)
	              log.SetFlags(log.LstdFlags | log.Lmicroseconds)

		      // Get the hash from the log
			hash := vLog.Topics[1].Hex()
			formattedHash := libs.ByteToByte32(common.Hex2Bytes(hash[2:]))

			//log.Printf("+ %s: Purchase Request %s (%s ----> %s\n)", purchaseID, hash, clientAddr.Hex(), iotAddr.Hex())
			log.Printf("%s Comprada\n",hash)
			log.Printf("%s Enviada\n",hash)

			_, err = ethBuyer.BalanceCon.PurchaseMeasurement(auth, formattedHash)
			if err != nil {
				log.Printf("Error en el purchase\n")
				f.Close()
				continue
			}
			log.Printf("%s Registrada\n",hash)
			f.Close()
		}
	}
	return
}

func autoBuy(ethBuyer libs.ComponentConfig) {
        if len(os.Args) < 3 {
                fmt.Printf("Error en los argumentos. El comprador automatico no se ha activado. Uso: ./<nombre> -buyerA [delay en segundos]\n")
                return
        }
	delay, _ := strconv.ParseFloat(os.Args[2], 64)
	if delay < 0.2 { delay = 1.0 }
	//rand.Seed(time.Now().Unix()) //seed

	auth := bind.NewKeyedTransactor(ethBuyer.PrivateKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	lastBlock, err := ethBuyer.EthereumClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(10),
		ToBlock: lastBlock.Number,
		Addresses: []common.Address{
			common.HexToAddress(ethBuyer.GeneralConfig["dataContractAddr"].(string))},
	}
	logs, err := ethBuyer.EthereumClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Comprando %d medidas...\n\n",len(logs))
	medidas := 0

	for _, vLog := range logs {
		medidas = medidas + 1
		//if medidas <= 5000 { continue }
		hash := vLog.Topics[1].Hex()
                formattedHash := libs.ByteToByte32(common.Hex2Bytes(hash[2:]))

		f, err := os.OpenFile("/home/administrator/actual_blockchain/timestampsC.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	        if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		wrt := io.MultiWriter(os.Stdout, f)
	        log.SetOutput(wrt)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds)

		log.Printf("%s Comprada\n",hash)
		log.Printf("%s Enviada\n",hash)

		_, err = ethBuyer.BalanceCon.PurchaseMeasurement(auth, formattedHash)
		if err != nil {
			log.Printf("Error en el purchase\n")
			f.Close()
			continue
		}
		log.Printf("%s Registrada\n",hash)
		f.Close()
		//dormir := float64(time.Second) * (-delay) * math.Log(rand.Float64()) //Poisson
		dormir := float64(time.Second) * delay //Constante
		//if medidas >= 5000 { break }
		if medidas == 5000 {
			log.Printf("AVISO\n")
		} else if medidas == 10000 {
			log.Printf("AVISO\n")
		} else if medidas >= 15000 {
			log.Printf("AVISO\n")
			break
		}
		time.Sleep(time.Duration(dormir))
	}
	return
}

func main() {
	fmt.Printf("----------- Initializing Core module -----------\n\n")
	localClient, localBuyer := initialize()

	fmt.Printf("Starting server on port %s\n\n", localClient.GeneralConfig["HTTPport"].(string))
	http.HandleFunc("/admin", localClient.AdminEP)

	// Convert the localClient to libs.ComponentConfig
	ethClient := libs.ComponentConfig{
		localClient.EthereumClient,
		localClient.Address,
		localClient.PrivateKey,
		localClient.PublicKey,
		localClient.DataCon,
		localClient.AccessCon,
		localClient.BalanceCon,
		localClient.IPFSPath,
		localClient.GeneralConfig,
	}

	ethBuyer := libs.ComponentConfig{
		localBuyer.EthereumClient,
		localBuyer.Address,
		localBuyer.PrivateKey,
		localBuyer.PublicKey,
		localBuyer.DataCon,
		localBuyer.AccessCon,
		localBuyer.BalanceCon,
		localClient.IPFSPath,
		localBuyer.GeneralConfig,
	}

	// Run the purchases manager thread
	go libs.ManagePurchases(ethClient)

	hayCompBuyer := flag.Bool("buyerC",false,"Activa o desactiva el comprador automatico de medidas. Escucha las medidas registradas y las compra instantáneamente.")
	hayAutoBuyer := flag.Bool("buyerA",false,"Activa o desactiva el comprador automatico de medidas. Busca las medidas previas y las compra a una determinada velocidad.")
	flag.Parse()
	if *hayCompBuyer {
		go compulsiveBuy(ethBuyer)
	}
	if *hayAutoBuyer {
		go autoBuy(ethBuyer)
	}

	err := http.ListenAndServe(":"+localClient.GeneralConfig["HTTPport"].(string), nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
