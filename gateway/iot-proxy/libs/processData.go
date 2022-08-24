package libs

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	//"time"
	"os"
	"io"

	cipher "administrator/ipfs-node/libs/cipher"
	ipfsLib "administrator/ipfs-node/libs/ipfsLib"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func xor32(a [32]byte, b [32]byte) [32]byte {
	var res [32]byte
	for i := 0; i < 32 ; i++ {
		res[i] = byte(int(a[i]) ^ int(b[i]))
	}
	return res
}

// Inserts the required information to retrieve a measurement in the Blockchain
func insertDataInBlockchain(ethClient ComponentConfig, dataStruct DataBlockchain, timest string) error {
	 // Log file
	f, err := os.OpenFile("../timestampsR.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
	     log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// Check that the measurement has not already been stored
	measurement, err := ethClient.DataCon.Ledger(nil, dataStruct.Hash)
	if err != nil {
		return err
	}

	// Check that the price of the measurement has already been set
	priceTag, err := ethClient.BalanceCon.GetPriceMeasurement(nil, dataStruct.Hash)
	if err != nil {
		return err
	}

	if measurement.Uri != "" {
		str := fmt.Sprintf("%x: The measurement had already been stored in the blockchain", dataStruct.Hash[:])
		log.Printf("%s Existente\n",timest)

		// Check if the stored measurement has a price. If not, set it.
		if priceTag.Uint64() == 0 {
			auth := bind.NewKeyedTransactor(ethClient.PrivateKey)
			auth.Value = big.NewInt(0)
			auth.GasLimit = uint64(3000000)
			auth.GasPrice = big.NewInt(0)

			price := (int64)(ethClient.GeneralConfig["priceMeasurements"].(float64))
			_, err = ethClient.BalanceCon.SetPriceToMeasurement(auth, dataStruct.Hash, big.NewInt(price))
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
		return errors.New(str)
	}

	// Prepare authentication parameters
	auth := bind.NewKeyedTransactor(ethClient.PrivateKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(0)

	// Send the transaction to the data smart contract
	log.Printf("%s Enviada\n",timest)
	_, err = ethClient.DataCon.StoreInfo(auth, dataStruct.Hash, dataStruct.EncryptedURL, dataStruct.Topics)
	if err != nil {
		log.Println(err)
		return err
	}

/*	// Check if the data has been stored in the contract
	//! This is not needed since calls wait for the response before continuing
	// Wait until the value is received or the loop
	// works for more than 15 seconds
	currentTime := time.Now()
	for {
		dataBC, err := ethClient.DataCon.Ledger(nil, dataStruct.Hash)
		if err != nil {
			log.Println(err)
			return err
		}

		if dataBC.Uri != "" {
			break
		}

		secondsPassed := time.Now().Sub(currentTime)
		if secondsPassed > 15*time.Second {
			log.Println("Could not check whether the data was introduced in the Blockchain")
			return errors.New("Could not check whether the data was introduced in the Blockchain")
		}
	}
*/
	// Set the price of the product
	price := (int64)(ethClient.GeneralConfig["priceMeasurements"].(float64))
	_, err = ethClient.BalanceCon.SetPriceToMeasurement(auth, dataStruct.Hash, big.NewInt(price))
	if err != nil {
		log.Println(err)
		return err
	}

/*	// Check if the data has been stored in the contract
	//! This is not needed since calls wait for the response before continuing
	// Wait until the value is received or the loop
	// works for more than 15 seconds
	currentTime = time.Now()
	for {

		dataBC, err := ethClient.BalanceCon.GetPriceMeasurement(nil, dataStruct.Hash)
		if err != nil {
			log.Println(err)
			return err
		}

		if dataBC != big.NewInt(0) {
			break
		}

		secondsPassed := time.Now().Sub(currentTime)
		if secondsPassed > 15*time.Second {
			log.Println("Could not check whether the data was introduced in the Blockchain")
			return errors.New("Could not check whether the data was introduced in the Blockchain")
		}
	}
*/
	log.Printf("%s Guardada\n",timest)
	return nil
}

// ProcessMeasurement processes the measurement:
// 	- Signs the measurement
//	- Encrypts the measurement with a random symmetric key
//	- Stores the measurement in the IPFS node
//  - Stores the IPFS URL in the Blockchain encrypted with
//	  the public key of the administrator
func ProcessMeasurement(ethClient ComponentConfig, body map[string]interface{}) error {
	//Log file
	f, err := os.OpenFile("../timestampsR.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// Convert the body to JSON data []byte
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Sign the measurement
	signedBody, err := cipher.SignData(ethClient.PrivateKey, jsonData)
	if err != nil {
		return err
	}

	// Append the signature to the measurement
	msg := append(jsonData, signedBody...)

	/* Prepare the data that is going to be stored in the Blockchain */
	observationDate := body["timestamp"].(string)
	sensorID := body["urn"].(string)
	phenomenon := body["phenomenon"].(string)
	gatewayID := ethClient.GeneralConfig["gatewayID"].(string)
	//description := sensorID + " by " + gatewayID + " at " + observationDate

	/*hasAccess, err := ethClient.AccessCon.AllowedAccounts(nil, ethClient.Address)
	if err != nil {
		return err
	}*/

	// Prepare authentication parameters
        auth := bind.NewKeyedTransactor(ethClient.PrivateKey)
        auth.Value = big.NewInt(0)
        auth.GasLimit = uint64(3000000)
        auth.GasPrice = big.NewInt(0)

	// Comprobar si los topics están ya creados
	created1, err := ethClient.DataCon.TopicExists(nil,gatewayID)
	if err != nil {
		return err
	}

	//fmt.Println(fmt.Sprintf("topic %s: %v",gatewayID,created1))
	var rk1 [32]byte
	if !created1 {
		randomKey := make([]byte, 32)
		rand.Read(randomKey)
		copy(rk1[:], randomKey)
		_, err = ethClient.DataCon.AddTopic(auth,gatewayID,rk1)
	}

	created2, err := ethClient.DataCon.TopicExists(nil,sensorID)
	if err != nil {
		return err
	}

	//fmt.Println(fmt.Sprintf("topic %s: %v",sensorID,created2))
	var rk2 [32]byte
	if !created2 {
		randomKey := make([]byte, 32)
		rand.Read(randomKey)
		copy(rk2[:], randomKey)
                _, err = ethClient.DataCon.AddTopic(auth,sensorID,rk2)
	}

	created3, err := ethClient.DataCon.TopicExists(nil,phenomenon)
	if err != nil {
		return err
	}

	//fmt.Println(fmt.Sprintf("topic %s: %v",phenomenon,created3))
	var rk3 [32]byte
	if !created3 {
		randomKey := make([]byte, 32)
		rand.Read(randomKey)
		copy(rk3[:], randomKey)
		_, err = ethClient.DataCon.AddTopic(auth,phenomenon,rk3)
	}

	created4, err := ethClient.DataCon.TopicExists(nil,observationDate[0:10])
	if err != nil {
		return err
	}

	//fmt.Println(fmt.Sprintf("topic %s: %v",observationDate[0:10],created4))
	var rk4 [32]byte
	if !created4 {
		randomKey := make([]byte, 32)
		rand.Read(randomKey)
		copy(rk4[:], randomKey)
		_, err = ethClient.DataCon.AddTopic(auth,observationDate[0:10],rk4)
	}

	topics := []string {gatewayID,sensorID,phenomenon,observationDate[0:10]}
	measurementHashBytes := cipher.HashData(jsonData)

	auth2 := &bind.CallOpts{
		From: ethClient.Address,
	}

	//Acquire topic keys
	var keyParcial [32]byte
	if created1 {
		keyParcial, err = ethClient.DataCon.GetTopicKey(auth2,gatewayID)
		if err != nil {
			return err
		}
	} else {
		keyParcial = rk1
	}
	key := keyParcial

	if created2 {
		keyParcial, err = ethClient.DataCon.GetTopicKey(auth2,sensorID)
		if err != nil {
			return err
		}
	} else {
		keyParcial = rk2
	}
	key = xor32(key,keyParcial)

	if created3 {
		keyParcial, err = ethClient.DataCon.GetTopicKey(auth2,phenomenon)
		if err != nil {
			return err
		}
	} else {
		keyParcial = rk3
	}
	key = xor32(key,keyParcial)

	if created4 {
		keyParcial, err = ethClient.DataCon.GetTopicKey(auth2,observationDate[0:10])
		if err != nil {
			return err
		}
	} else {
		keyParcial = rk4
	}
	key = xor32(key,keyParcial)
	keyFinal := key[:]

	// Encrypt the measurement (msg) with the key
	encryptedMsg, err := cipher.SymmetricEncryption(keyFinal, msg)
	if err != nil {
		return err
	}

	/* Store the encrypted measurement in the IPFS network */
	log.Printf("%s Registrando\n",observationDate)
	// Convert bytes to files.node
	cid, err := ipfsLib.AddToIPFS(ethClient.IPFSConfig.IpfsCore, bytes.NewReader(encryptedMsg))

	// Append the cid to the symmetric key to store them in the Blockchain (BC)
	secretBC := []byte(cid)
	log.Printf("%s Registrada\n",observationDate)

	// Get the public key of the marketplace from the Blockchain
	adminPubKeyString, err := ethClient.AccessCon.AdminPublicKey(nil)
	if err != nil {
		return err
	}

	// Convert the string public key to bytes
	adminPubKeyBytes, err := hex.DecodeString(adminPubKeyString)
	if err != nil {
		return err
	}

	// Convert the public key to ecdsa.PublicKey
	adminPubKey, err := crypto.UnmarshalPubkey(adminPubKeyBytes)
	if err != nil {
		return err
	}

	// Encrypt the url with the public key of the marketplace
	encryptedURL, err := cipher.EncryptWithPublicKey(*adminPubKey, secretBC)
	if err != nil {
		return err
	}

	dataStruct := DataBlockchain{
		ByteToByte32(measurementHashBytes),
		//description,
		topics,
		fmt.Sprintf("%x", encryptedURL),
	}

	/* Introduce data in the Blockchain */
	err = insertDataInBlockchain(ethClient, dataStruct, observationDate)
	if err != nil {
		return err
	}

	//log.Printf("Medida guardada en la blockchain. Hash: 0x%x\n", measurementHashBytes)

	return nil
}
