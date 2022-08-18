package libs

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	accessControlContract "marketplace/core/contracts/accessContract"
	balanceContract "marketplace/core/contracts/balanceContract"
	dataContract "marketplace/core/contracts/dataContract"
)

// ComponentConfig is a struct that stores the parameters of the node
type ComponentConfig struct {
	EthereumClient *ethclient.Client
	Address        common.Address
	PrivateKey     *ecdsa.PrivateKey
	PublicKey      ecdsa.PublicKey
	DataCon        *dataContract.DataLedgerContract
	AccessCon      *accessControlContract.AccessControlContract
	BalanceCon     *balanceContract.BalanceContract
	IPFSPath       string
	GeneralConfig  map[string]interface{}
}

// ByteToByte32 converts []byte to [32]byte
func ByteToByte32(bytes []byte) [32]byte {
	var b [32]byte
	copy(b[:], bytes)
	return b
}

// SendTransaction sends transaction with data to the blockchain
func SendTransaction(from, to common.Address,
	privKey *ecdsa.PrivateKey,
	ethclient ComponentConfig,
	data []byte) ([]byte, error) {

	// Set the parameters of the transaction
	value := big.NewInt(0)
	gasLimit := uint64(400000)
	gasPrice := big.NewInt(0)

	nonce, err := ethclient.EthereumClient.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal(err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

	// Sign the transaction with the private key of the sender
	chainID, err := ethclient.EthereumClient.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Send the transaction
	err = ethclient.EthereumClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return signedTx.Hash().Bytes(), nil
}
