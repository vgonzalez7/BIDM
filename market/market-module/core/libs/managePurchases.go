package libs

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"io"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	cipher "marketplace/core/cipherLibs"
)

type nonceLock struct {
	mlock sync.Mutex
}

// Complete the purchases produced in the Blockchain.
// 	- Decipher the symmetric key and the URL with the private key
//	- Cipher the URL and the symmetric key with the public key of the client
//		and send them to the customer through a BC transaction
//	- Complete the transaction of tokens
func (nL *nonceLock) completePurchase(ethClient ComponentConfig, vLog types.Log) {
	// Log file
	f, err := os.OpenFile("/home/administrator/actual_blockchain/timestampsC.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
	      log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	auth := bind.NewKeyedTransactor(ethClient.PrivateKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	// Purchase ID (TxHash of the event)
	//purchaseID := vLog.Topics[0].Hex()

	// Get the address of the account who purchased the measurement
	clientAddrHex := vLog.Topics[2].Hex()
	clientAddr := common.HexToAddress(clientAddrHex[2:])

	// Get the address of the owner of the measurement
	//iotAddrHex := vLog.Topics[3].Hex()
	//iotAddr := common.HexToAddress(iotAddrHex)

	// Get the hash from the log
	hash := vLog.Topics[1].Hex()
	formattedHash := ByteToByte32(common.Hex2Bytes(hash[2:]))

	//log.Printf("+ %s: Purchase Request %s (%s ----> %s\n)", purchaseID, hash, clientAddr.Hex(), iotAddr.Hex())
	log.Printf("%s Detectada\n",hash)

	// Get the url where the measurement is stored
	dataStruct, err := ethClient.DataCon.Ledger(nil, formattedHash)
	if err != nil {
		// revoke transaction
		_, err := ethClient.BalanceCon.RevokeTransaction(auth, formattedHash, clientAddr)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	decryptedURL, err := cipher.DecryptWithPrivateKey(ethClient.PrivateKey, common.Hex2Bytes(dataStruct.Uri))
	if err != nil {
		log.Println(err)
		// revoke transaction
		_, err := ethClient.BalanceCon.RevokeTransaction(auth, formattedHash, clientAddr)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// Get the public key of the client who purchased the measurement
	clientPubKeyHex, err := ethClient.AccessCon.PubKeysKeystore(nil, clientAddr)
	if err != nil {
		// revoke transaction
		_, err := ethClient.BalanceCon.RevokeTransaction(auth, formattedHash, clientAddr)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// Convert public key to the appropiate format
	clientPubKeyBytes := common.Hex2Bytes(clientPubKeyHex)
	clientPubKeyECDSA, err := crypto.UnmarshalPubkey(clientPubKeyBytes)
	if err != nil {
		// revoke transaction
		_, err := ethClient.BalanceCon.RevokeTransaction(auth, formattedHash, clientAddr)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// cipher the URL with the public key of the client
	encryptedURL, err := cipher.EncryptWithPublicKey(*clientPubKeyECDSA, decryptedURL)
	if err != nil {
		// revoke transaction
		_, err := ethClient.BalanceCon.RevokeTransaction(auth, formattedHash, clientAddr)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	nL.mlock.Lock()
	// Send transaction through the Blockchain with the encrypted URL of the measurement
	txHash, err := SendTransaction(ethClient.Address, clientAddr, ethClient.PrivateKey, ethClient, encryptedURL)
	if err != nil {
		// revoke transaction
		_, err := ethClient.BalanceCon.RevokeTransaction(auth, formattedHash, clientAddr)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// Complete the transaction of tokens
	_, err = ethClient.BalanceCon.CompletePurchase(auth, formattedHash, clientAddr, ByteToByte32(txHash))
	if err != nil {
		// revoke transaction
		_, err := ethClient.BalanceCon.RevokeTransaction(auth, formattedHash, clientAddr)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	nL.mlock.Unlock()

	//log.Printf("\n+ %s: Purchase Completed %s (%s ----> %s)", purchaseID, hash, clientAddr.Hex(), iotAddr.Hex())
	log.Printf("%s Completada\n",hash)
}

// ManagePurchases listens to the purchases events in
// the Blockchain and processes them
func ManagePurchases(ethClient ComponentConfig) {

	nL := nonceLock{}
	// Prepare the filter to filter the Balance contract events.
	// Watch the events caused by the purchase of a measurement in the Blockhain
	topic := []byte("RequestPurchase(bytes32,address,address)")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			common.HexToAddress(ethClient.GeneralConfig["balanceContractAddr"].(string))},
		Topics: [][]common.Hash{{crypto.Keccak256Hash(topic)}},
	}

	// Make channel to receive the events
	logs := make(chan types.Log)
	sub, err := ethClient.EthereumClient.SubscribeFilterLogs(context.Background(), query, logs)
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
			go nL.completePurchase(ethClient, vLog)
		}
	}
}
