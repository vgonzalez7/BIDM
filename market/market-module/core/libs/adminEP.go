package libs

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"bytes"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cipher "marketplace/core/cipherLibs"
)

func xor32(a [32]byte, b [32]byte) [32]byte {
        var res [32]byte
        for i := 0; i < 32 ; i++ {
                res[i] = byte(int(a[i]) ^ int(b[i]))
        }
        return res
}

// AdminListener listens to the request made by the admin user
func AdminListener(ethClient ComponentConfig, body map[string]interface{}) ([]byte,error) {
	protocol := body["action"].(string)

	switch protocol {
	// Transferencia de tokens a un usuario
	case "transferencia":
		addr := body["addr"].(string)
		tokens := int64(body["tokens"].(float64))

		fmt.Println(fmt.Sprintf("+ Sending %d to %s", tokens, addr))

		// Send the tokens to the client
		auth := bind.NewKeyedTransactor(ethClient.PrivateKey)
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(400000)
		auth.GasPrice = big.NewInt(0)
		_, err := ethClient.BalanceCon.SendTokenToClient(auth, common.HexToAddress(addr), big.NewInt(tokens))
		if err != nil { return nil,err }

	case "medida":
		hash := body["hash"].(string)
		addr := body["addr"].(string)
		//callback := body["callback"].(string)
		fmt.Println(fmt.Sprintf("Sending 0x%s to 0x%s",hash,addr))

		auth2 := &bind.CallOpts{ From: ethClient.Address, }
		hashAux := common.Hex2Bytes(hash)
		var bHash [32]byte
		copy(bHash[:], hashAux)
		dataStruct, err := ethClient.DataCon.Ledger(nil, bHash)
        if err != nil { return nil,err }
        cidB, _ := cipher.DecryptWithPrivateKey(ethClient.PrivateKey, common.Hex2Bytes(dataStruct.Uri))
        topics, err := ethClient.DataCon.GetMeasurementTopics(auth2,bHash)
	if err != nil { return nil,err }
        cid := string(cidB)

        fullurl := fmt.Sprintf("%s/api/v0/cat?arg=%s",ethClient.IPFSPath,cid)
  		resp, err := http.Post(fullurl, "text/plain", nil)
  		if err != nil { return nil,err }
  		buf := new(bytes.Buffer)
	    buf.ReadFrom(resp.Body)
	    response := buf.Bytes()
	    defer resp.Body.Close()

        //Get topic keys
        var keyParcial [32]byte
        keyParcial, err = ethClient.DataCon.GetTopicKey(auth2,topics[0])
		if err != nil { return nil,err }
        key := keyParcial

        var i int
        for i = 1;i < len(topics); i++ {
        	keyParcial, err = ethClient.DataCon.GetTopicKey(auth2,topics[i])
        	key = xor32(key,keyParcial)
        }
        keyFinal := key[:]

        measurement, err := cipher.DecryptSymmetricEncryption(keyFinal, response)
        if err != nil { return nil,err }

        return measurement,nil


	default:
		return nil,errors.New("Unelegible protocol")
	}

	return nil,nil
}
