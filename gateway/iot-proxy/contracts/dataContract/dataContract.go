// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dataContract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AccessControlContractABI is the input ABI used to generate the binding from.
const AccessControlContractABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccessControlContractFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlContractFuncSigs = map[string]string{
	"e04610ed": "allowedAccounts(address)",
}

// AccessControlContractBin is the compiled bytecode used for deploying new contracts.
var AccessControlContractBin = "0x608060405234801561001057600080fd5b50610121806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e04610ed14602d575b600080fd5b603c60383660046074565b6050565b6040516047919060a4565b60405180910390f35b60006020819052908152604090205460ff1681565b8035606e8160ca565b92915050565b600060208284031215608557600080fd5b6000608f84846065565b949350505050565b609e8160b9565b82525050565b60208101606e82846097565b6000606e8260be565b151590565b6001600160a01b031690565b60d18160b0565b811460db57600080fd5b5056fea365627a7a723158203c8214c2c9dacf231df433dada82f987058d85d6d2b01fba647fee24af2699f26c6578706572696d656e74616cf564736f6c63430005100040"

// DeployAccessControlContract deploys a new Ethereum contract, binding an instance of AccessControlContract to it.
func DeployAccessControlContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccessControlContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccessControlContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControlContract{AccessControlContractCaller: AccessControlContractCaller{contract: contract}, AccessControlContractTransactor: AccessControlContractTransactor{contract: contract}, AccessControlContractFilterer: AccessControlContractFilterer{contract: contract}}, nil
}

// AccessControlContract is an auto generated Go binding around an Ethereum contract.
type AccessControlContract struct {
	AccessControlContractCaller     // Read-only binding to the contract
	AccessControlContractTransactor // Write-only binding to the contract
	AccessControlContractFilterer   // Log filterer for contract events
}

// AccessControlContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlContractSession struct {
	Contract     *AccessControlContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AccessControlContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlContractCallerSession struct {
	Contract *AccessControlContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AccessControlContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlContractTransactorSession struct {
	Contract     *AccessControlContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AccessControlContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlContractRaw struct {
	Contract *AccessControlContract // Generic contract binding to access the raw methods on
}

// AccessControlContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlContractCallerRaw struct {
	Contract *AccessControlContractCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlContractTransactorRaw struct {
	Contract *AccessControlContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlContract creates a new instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContract(address common.Address, backend bind.ContractBackend) (*AccessControlContract, error) {
	contract, err := bindAccessControlContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlContract{AccessControlContractCaller: AccessControlContractCaller{contract: contract}, AccessControlContractTransactor: AccessControlContractTransactor{contract: contract}, AccessControlContractFilterer: AccessControlContractFilterer{contract: contract}}, nil
}

// NewAccessControlContractCaller creates a new read-only instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContractCaller(address common.Address, caller bind.ContractCaller) (*AccessControlContractCaller, error) {
	contract, err := bindAccessControlContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractCaller{contract: contract}, nil
}

// NewAccessControlContractTransactor creates a new write-only instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlContractTransactor, error) {
	contract, err := bindAccessControlContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractTransactor{contract: contract}, nil
}

// NewAccessControlContractFilterer creates a new log filterer instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlContractFilterer, error) {
	contract, err := bindAccessControlContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractFilterer{contract: contract}, nil
}

// bindAccessControlContract binds a generic wrapper to an already deployed contract.
func bindAccessControlContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlContract *AccessControlContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlContract.Contract.AccessControlContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlContract *AccessControlContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AccessControlContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlContract *AccessControlContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AccessControlContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlContract *AccessControlContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlContract *AccessControlContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlContract *AccessControlContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlContract.Contract.contract.Transact(opts, method, params...)
}

// AllowedAccounts is a free data retrieval call binding the contract method 0xe04610ed.
//
// Solidity: function allowedAccounts(address ) view returns(bool)
func (_AccessControlContract *AccessControlContractCaller) AllowedAccounts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControlContract.contract.Call(opts, &out, "allowedAccounts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedAccounts is a free data retrieval call binding the contract method 0xe04610ed.
//
// Solidity: function allowedAccounts(address ) view returns(bool)
func (_AccessControlContract *AccessControlContractSession) AllowedAccounts(arg0 common.Address) (bool, error) {
	return _AccessControlContract.Contract.AllowedAccounts(&_AccessControlContract.CallOpts, arg0)
}

// AllowedAccounts is a free data retrieval call binding the contract method 0xe04610ed.
//
// Solidity: function allowedAccounts(address ) view returns(bool)
func (_AccessControlContract *AccessControlContractCallerSession) AllowedAccounts(arg0 common.Address) (bool, error) {
	return _AccessControlContract.Contract.AllowedAccounts(&_AccessControlContract.CallOpts, arg0)
}

// DataLedgerContractABI is the input ABI used to generate the binding from.
const DataLedgerContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"deleteInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"_topics\",\"type\":\"string[]\"}],\"name\":\"evtStoreInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_topic\",\"type\":\"string\"}],\"name\":\"newTopic\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"addTopic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"}],\"name\":\"checkAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"deleteMeasurement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getIoTAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"getIoTAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getMeasurementTopics\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"getTopicKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTopics\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ledger\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"topic\",\"type\":\"string[]\"}],\"name\":\"storeInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"topicExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataLedgerContractFuncSigs maps the 4-byte function signature to its string representation.
var DataLedgerContractFuncSigs = map[string]string{
	"3eeb446e": "addTopic(string,bytes32)",
	"466a0146": "checkAccess(address)",
	"77ad95ca": "deleteMeasurement(bytes32)",
	"6ade0219": "getIoTAddress(bytes32)",
	"ab680e02": "getIoTAddress(string)",
	"dda4997c": "getMeasurementTopics(bytes32)",
	"8a746949": "getTopicKey(string)",
	"fee01177": "getTopics()",
	"15977d45": "ledger(bytes32)",
	"e30081a0": "setAddress(address)",
	"0335e024": "storeInfo(bytes32,string,string[])",
	"78b5749d": "topicExists(string)",
}

// DataLedgerContractBin is the compiled bytecode used for deploying new contracts.
var DataLedgerContractBin = "0x6080604052600180546001600160a01b031916737fff1f93d22e9b2aa4c6a536531950ed386bd95e17905534801561003657600080fd5b506112c8806100466000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806378b5749d1161007157806378b5749d1461015e5780638a74694914610171578063ab680e0214610191578063dda4997c146101a4578063e30081a0146101c4578063fee01177146101d7576100b4565b80630335e024146100b957806315977d45146100ce5780633eeb446e146100f8578063466a01461461010b5780636ade02191461012b57806377ad95ca1461014b575b600080fd5b6100cc6100c7366004610c9a565b6101df565b005b6100e16100dc366004610c7c565b610324565b6040516100ef9291906110fe565b60405180910390f35b6100cc610106366004610d4c565b6103ce565b61011e610119366004610c38565b61050a565b6040516100ef91906110bd565b61013e610139366004610c7c565b6105b0565b6040516100ef919061109e565b6100cc610159366004610c7c565b6105cf565b61011e61016c366004610d17565b610660565b61018461017f366004610d17565b6106f4565b6040516100ef91906110cb565b61013e61019f366004610d17565b61076e565b6101b76101b2366004610c7c565b6107a0565b6040516100ef91906110ac565b6100cc6101d2366004610c38565b61088e565b6101b76108da565b6101e83361050a565b15156001146102125760405162461bcd60e51b81526004016102099061115e565b60405180910390fd5b61021a6109b3565b815160005b8181101561026c5761024384828151811061023657fe5b6020026020010151610660565b15156001146102645760405162461bcd60e51b81526004016102099061111e565b60010161021f565b50838252602080830184905233604080850191909152600087815260028352208351805185936102a09284929101906109dd565b5060208281015180516102b99260018501920190610a5b565b5060409182015160029190910180546001600160a01b0319166001600160a01b039092169190911790555185907f717b8d30e5c9a0e4112f8c62e289d73e44dc76c0e1254192334f3735ea4a86559061031590879087906110d9565b60405180910390a25050505050565b600260208181526000928352604092839020805484516001821615610100026000190190911693909304601f81018390048302840183019094528383529283918301828280156103b55780601f1061038a576101008083540402835291602001916103b5565b820191906000526020600020905b81548152906001019060200180831161039857829003601f168201915b505050600290930154919250506001600160a01b031682565b6103d73361050a565b15156001146103f85760405162461bcd60e51b81526004016102099061115e565b61040182610660565b1561041e5760405162461bcd60e51b81526004016102099061114e565b60038054600181018083556000929092528351610462917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b019060208601906109dd565b5050336004836040516104759190611086565b908152602001604051809103902060006101000a8154816001600160a01b0302191690836001600160a01b03160217905550806005836040516104b89190611086565b908152604051908190036020018120919091556104d6908390611086565b604051908190038120907fbbced0d0b08bca32f3efec0aa77ad0d7d183168e8f976b456ad70129694c41a690600090a25050565b6001546000906001600160a01b0316331415610528575060016105ab565b60005460405163e04610ed60e01b81526001600160a01b039091169063e04610ed9061055890859060040161109e565b60206040518083038186803b15801561057057600080fd5b505afa158015610584573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506105a89190810190610c5e565b90505b919050565b600090815260026020819052604090912001546001600160a01b031690565b6001546001600160a01b031633146105f95760405162461bcd60e51b81526004016102099061112e565b6000818152600260205260408120906106128282610ab4565b610620600183016000610afb565b5060020180546001600160a01b031916905560405181907f072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe90600090a250565b60035460009081805b828110156106ec57846040516020016106829190611086565b60405160208183030381529060405280519060200120600382815481106106a557fe5b906000526020600020016040516020016106bf9190611092565b6040516020818303038152906040528051906020012014156106e457600191506106ec565b600101610669565b509392505050565b6001546000906001600160a01b031633148061072a576107138361076e565b6001600160a01b0316336001600160a01b03161490505b806107475760405162461bcd60e51b81526004016102099061113e565b6005836040516107579190611086565b908152602001604051809103902054915050919050565b60006004826040516107809190611086565b908152604051908190036020019020546001600160a01b03169050919050565b606060026000838152602001908152602001600020600101805480602002602001604051908101604052809291908181526020016000905b828210156108835760008481526020908190208301805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801561086f5780601f106108445761010080835404028352916020019161086f565b820191906000526020600020905b81548152906001019060200180831161085257829003601f168201915b5050505050815260200190600101906107d8565b505050509050919050565b6001546001600160a01b031633146108b85760405162461bcd60e51b81526004016102099061113e565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b60606003805480602002602001604051908101604052809291908181526020016000905b828210156109a95760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156109955780601f1061096a57610100808354040283529160200191610995565b820191906000526020600020905b81548152906001019060200180831161097857829003601f168201915b5050505050815260200190600101906108fe565b5050505090505b90565b6040518060600160405280606081526020016060815260200160006001600160a01b031681525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a1e57805160ff1916838001178555610a4b565b82800160010185558215610a4b579182015b82811115610a4b578251825591602001919060010190610a30565b50610a57929150610b19565b5090565b828054828255906000526020600020908101928215610aa8579160200282015b82811115610aa85782518051610a989184916020909101906109dd565b5091602001919060010190610a7b565b50610a57929150610b33565b50805460018160011615610100020316600290046000825580601f10610ada5750610af8565b601f016020900490600052602060002090810190610af89190610b19565b50565b5080546000825590600052602060002090810190610af89190610b33565b6109b091905b80821115610a575760008155600101610b1f565b6109b091905b80821115610a57576000610b4d8282610ab4565b50600101610b39565b8035610b618161125f565b92915050565b600082601f830112610b7857600080fd5b8135610b8b610b8682611195565b61116e565b81815260209384019390925082018360005b83811015610bc95781358601610bb38882610be9565b8452506020928301929190910190600101610b9d565b5050505092915050565b8051610b6181611273565b8035610b618161127c565b600082601f830112610bfa57600080fd5b8135610c08610b86826111b6565b91508082526020830160208301858383011115610c2457600080fd5b610c2f838284611219565b50505092915050565b600060208284031215610c4a57600080fd5b6000610c568484610b56565b949350505050565b600060208284031215610c7057600080fd5b6000610c568484610bd3565b600060208284031215610c8e57600080fd5b6000610c568484610bde565b600080600060608486031215610caf57600080fd5b6000610cbb8686610bde565b935050602084013567ffffffffffffffff811115610cd857600080fd5b610ce486828701610be9565b925050604084013567ffffffffffffffff811115610d0157600080fd5b610d0d86828701610b67565b9150509250925092565b600060208284031215610d2957600080fd5b813567ffffffffffffffff811115610d4057600080fd5b610c5684828501610be9565b60008060408385031215610d5f57600080fd5b823567ffffffffffffffff811115610d7657600080fd5b610d8285828601610be9565b9250506020610d9385828601610bde565b9150509250929050565b6000610da98383610e3f565b9392505050565b610db9816111fd565b82525050565b6000610dca826111f0565b610dd481856111f4565b935083602082028501610de6856111de565b8060005b85811015610e205784840389528151610e038582610d9d565b9450610e0e836111de565b60209a909a0199925050600101610dea565b5091979650505050505050565b610db981611208565b610db9816109b0565b6000610e4a826111f0565b610e5481856111f4565b9350610e64818560208601611225565b610e6d81611255565b9093019392505050565b6000610e82826111f0565b610e8c81856105ab565b9350610e9c818560208601611225565b9290920192915050565b600081546001811660008114610ec35760018114610ee657610f25565b607f6002830416610ed481876105ab565b60ff1984168152955085019250610f25565b60028204610ef481876105ab565b9550610eff856111e4565b60005b82811015610f1e57815488820152600190910190602001610f02565b5050850192505b505092915050565b6000610f3a6018836111f4565b7f54686520746f70696320646f6573206e6f742065786973740000000000000000815260200192915050565b6000610f736033836111f4565b7f596f7520646f206e6f74206861766520656e6f7567682070726976696c65676581527239903a37903237903a3434b99030b1ba34b7b760691b602082015260400192915050565b6000610fc8602c836111f4565b7f596f7520646f206e6f7420686176652070726976696c6567657320746f20646f81526b103a3434b99030b1ba34b7b760a11b602082015260400192915050565b60006110166014836111f4565b73546f70696320616c72656164792065786973747360601b815260200192915050565b6000611046602b836111f4565b7f546865204944207468617420796f7520617265207573696e67206973206e6f7481526a081c9959da5cdd195c995960aa1b602082015260400192915050565b6000610da98284610e77565b6000610da98284610ea6565b60208101610b618284610db0565b60208082528101610da98184610dbf565b60208101610b618284610e2d565b60208101610b618284610e36565b604080825281016110ea8185610e3f565b90508181036020830152610c568184610dbf565b6040808252810161110f8185610e3f565b9050610da96020830184610db0565b602080825281016105a881610f2d565b602080825281016105a881610f66565b602080825281016105a881610fbb565b602080825281016105a881611009565b602080825281016105a881611039565b60405181810167ffffffffffffffff8111828210171561118d57600080fd5b604052919050565b600067ffffffffffffffff8211156111ac57600080fd5b5060209081020190565b600067ffffffffffffffff8211156111cd57600080fd5b506020601f91909101601f19160190565b60200190565b60009081526020902090565b5190565b90815260200190565b60006105a88261120d565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015611240578181015183820152602001611228565b8381111561124f576000848401525b50505050565b601f01601f191690565b611268816111fd565b8114610af857600080fd5b61126881611208565b611268816109b056fea365627a7a72315820febdcdb1d3dbe6aa1ddb212034578e36a2dbc392f87bef5ec9c1af5b2a50d3a76c6578706572696d656e74616cf564736f6c63430005100040"

// DeployDataLedgerContract deploys a new Ethereum contract, binding an instance of DataLedgerContract to it.
func DeployDataLedgerContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataLedgerContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataLedgerContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataLedgerContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataLedgerContract{DataLedgerContractCaller: DataLedgerContractCaller{contract: contract}, DataLedgerContractTransactor: DataLedgerContractTransactor{contract: contract}, DataLedgerContractFilterer: DataLedgerContractFilterer{contract: contract}}, nil
}

// DataLedgerContract is an auto generated Go binding around an Ethereum contract.
type DataLedgerContract struct {
	DataLedgerContractCaller     // Read-only binding to the contract
	DataLedgerContractTransactor // Write-only binding to the contract
	DataLedgerContractFilterer   // Log filterer for contract events
}

// DataLedgerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataLedgerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLedgerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataLedgerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLedgerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataLedgerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLedgerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataLedgerContractSession struct {
	Contract     *DataLedgerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DataLedgerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataLedgerContractCallerSession struct {
	Contract *DataLedgerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DataLedgerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataLedgerContractTransactorSession struct {
	Contract     *DataLedgerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DataLedgerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataLedgerContractRaw struct {
	Contract *DataLedgerContract // Generic contract binding to access the raw methods on
}

// DataLedgerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataLedgerContractCallerRaw struct {
	Contract *DataLedgerContractCaller // Generic read-only contract binding to access the raw methods on
}

// DataLedgerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataLedgerContractTransactorRaw struct {
	Contract *DataLedgerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataLedgerContract creates a new instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContract(address common.Address, backend bind.ContractBackend) (*DataLedgerContract, error) {
	contract, err := bindDataLedgerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContract{DataLedgerContractCaller: DataLedgerContractCaller{contract: contract}, DataLedgerContractTransactor: DataLedgerContractTransactor{contract: contract}, DataLedgerContractFilterer: DataLedgerContractFilterer{contract: contract}}, nil
}

// NewDataLedgerContractCaller creates a new read-only instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContractCaller(address common.Address, caller bind.ContractCaller) (*DataLedgerContractCaller, error) {
	contract, err := bindDataLedgerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractCaller{contract: contract}, nil
}

// NewDataLedgerContractTransactor creates a new write-only instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DataLedgerContractTransactor, error) {
	contract, err := bindDataLedgerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractTransactor{contract: contract}, nil
}

// NewDataLedgerContractFilterer creates a new log filterer instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DataLedgerContractFilterer, error) {
	contract, err := bindDataLedgerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractFilterer{contract: contract}, nil
}

// bindDataLedgerContract binds a generic wrapper to an already deployed contract.
func bindDataLedgerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataLedgerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataLedgerContract *DataLedgerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataLedgerContract.Contract.DataLedgerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataLedgerContract *DataLedgerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DataLedgerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataLedgerContract *DataLedgerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DataLedgerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataLedgerContract *DataLedgerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataLedgerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataLedgerContract *DataLedgerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataLedgerContract *DataLedgerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.contract.Transact(opts, method, params...)
}

// CheckAccess is a free data retrieval call binding the contract method 0x466a0146.
//
// Solidity: function checkAccess(address producer) view returns(bool)
func (_DataLedgerContract *DataLedgerContractCaller) CheckAccess(opts *bind.CallOpts, producer common.Address) (bool, error) {
	var out []interface{}
	err := _DataLedgerContract.contract.Call(opts, &out, "checkAccess", producer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckAccess is a free data retrieval call binding the contract method 0x466a0146.
//
// Solidity: function checkAccess(address producer) view returns(bool)
func (_DataLedgerContract *DataLedgerContractSession) CheckAccess(producer common.Address) (bool, error) {
	return _DataLedgerContract.Contract.CheckAccess(&_DataLedgerContract.CallOpts, producer)
}

// CheckAccess is a free data retrieval call binding the contract method 0x466a0146.
//
// Solidity: function checkAccess(address producer) view returns(bool)
func (_DataLedgerContract *DataLedgerContractCallerSession) CheckAccess(producer common.Address) (bool, error) {
	return _DataLedgerContract.Contract.CheckAccess(&_DataLedgerContract.CallOpts, producer)
}

// GetIoTAddress is a free data retrieval call binding the contract method 0x6ade0219.
//
// Solidity: function getIoTAddress(bytes32 hash) view returns(address)
func (_DataLedgerContract *DataLedgerContractCaller) GetIoTAddress(opts *bind.CallOpts, hash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _DataLedgerContract.contract.Call(opts, &out, "getIoTAddress", hash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIoTAddress is a free data retrieval call binding the contract method 0x6ade0219.
//
// Solidity: function getIoTAddress(bytes32 hash) view returns(address)
func (_DataLedgerContract *DataLedgerContractSession) GetIoTAddress(hash [32]byte) (common.Address, error) {
	return _DataLedgerContract.Contract.GetIoTAddress(&_DataLedgerContract.CallOpts, hash)
}

// GetIoTAddress is a free data retrieval call binding the contract method 0x6ade0219.
//
// Solidity: function getIoTAddress(bytes32 hash) view returns(address)
func (_DataLedgerContract *DataLedgerContractCallerSession) GetIoTAddress(hash [32]byte) (common.Address, error) {
	return _DataLedgerContract.Contract.GetIoTAddress(&_DataLedgerContract.CallOpts, hash)
}

// GetIoTAddress0 is a free data retrieval call binding the contract method 0xab680e02.
//
// Solidity: function getIoTAddress(string topic) view returns(address)
func (_DataLedgerContract *DataLedgerContractCaller) GetIoTAddress0(opts *bind.CallOpts, topic string) (common.Address, error) {
	var out []interface{}
	err := _DataLedgerContract.contract.Call(opts, &out, "getIoTAddress0", topic)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIoTAddress0 is a free data retrieval call binding the contract method 0xab680e02.
//
// Solidity: function getIoTAddress(string topic) view returns(address)
func (_DataLedgerContract *DataLedgerContractSession) GetIoTAddress0(topic string) (common.Address, error) {
	return _DataLedgerContract.Contract.GetIoTAddress0(&_DataLedgerContract.CallOpts, topic)
}

// GetIoTAddress0 is a free data retrieval call binding the contract method 0xab680e02.
//
// Solidity: function getIoTAddress(string topic) view returns(address)
func (_DataLedgerContract *DataLedgerContractCallerSession) GetIoTAddress0(topic string) (common.Address, error) {
	return _DataLedgerContract.Contract.GetIoTAddress0(&_DataLedgerContract.CallOpts, topic)
}

// GetMeasurementTopics is a free data retrieval call binding the contract method 0xdda4997c.
//
// Solidity: function getMeasurementTopics(bytes32 hash) view returns(string[])
func (_DataLedgerContract *DataLedgerContractCaller) GetMeasurementTopics(opts *bind.CallOpts, hash [32]byte) ([]string, error) {
	var out []interface{}
	err := _DataLedgerContract.contract.Call(opts, &out, "getMeasurementTopics", hash)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetMeasurementTopics is a free data retrieval call binding the contract method 0xdda4997c.
//
// Solidity: function getMeasurementTopics(bytes32 hash) view returns(string[])
func (_DataLedgerContract *DataLedgerContractSession) GetMeasurementTopics(hash [32]byte) ([]string, error) {
	return _DataLedgerContract.Contract.GetMeasurementTopics(&_DataLedgerContract.CallOpts, hash)
}

// GetMeasurementTopics is a free data retrieval call binding the contract method 0xdda4997c.
//
// Solidity: function getMeasurementTopics(bytes32 hash) view returns(string[])
func (_DataLedgerContract *DataLedgerContractCallerSession) GetMeasurementTopics(hash [32]byte) ([]string, error) {
	return _DataLedgerContract.Contract.GetMeasurementTopics(&_DataLedgerContract.CallOpts, hash)
}

// GetTopicKey is a free data retrieval call binding the contract method 0x8a746949.
//
// Solidity: function getTopicKey(string topic) view returns(bytes32)
func (_DataLedgerContract *DataLedgerContractCaller) GetTopicKey(opts *bind.CallOpts, topic string) ([32]byte, error) {
	var out []interface{}
	err := _DataLedgerContract.contract.Call(opts, &out, "getTopicKey", topic)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTopicKey is a free data retrieval call binding the contract method 0x8a746949.
//
// Solidity: function getTopicKey(string topic) view returns(bytes32)
func (_DataLedgerContract *DataLedgerContractSession) GetTopicKey(topic string) ([32]byte, error) {
	return _DataLedgerContract.Contract.GetTopicKey(&_DataLedgerContract.CallOpts, topic)
}

// GetTopicKey is a free data retrieval call binding the contract method 0x8a746949.
//
// Solidity: function getTopicKey(string topic) view returns(bytes32)
func (_DataLedgerContract *DataLedgerContractCallerSession) GetTopicKey(topic string) ([32]byte, error) {
	return _DataLedgerContract.Contract.GetTopicKey(&_DataLedgerContract.CallOpts, topic)
}

// GetTopics is a free data retrieval call binding the contract method 0xfee01177.
//
// Solidity: function getTopics() view returns(string[])
func (_DataLedgerContract *DataLedgerContractCaller) GetTopics(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _DataLedgerContract.contract.Call(opts, &out, "getTopics")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetTopics is a free data retrieval call binding the contract method 0xfee01177.
//
// Solidity: function getTopics() view returns(string[])
func (_DataLedgerContract *DataLedgerContractSession) GetTopics() ([]string, error) {
	return _DataLedgerContract.Contract.GetTopics(&_DataLedgerContract.CallOpts)
}

// GetTopics is a free data retrieval call binding the contract method 0xfee01177.
//
// Solidity: function getTopics() view returns(string[])
func (_DataLedgerContract *DataLedgerContractCallerSession) GetTopics() ([]string, error) {
	return _DataLedgerContract.Contract.GetTopics(&_DataLedgerContract.CallOpts)
}

// Ledger is a free data retrieval call binding the contract method 0x15977d45.
//
// Solidity: function ledger(bytes32 ) view returns(string uri, address addr)
func (_DataLedgerContract *DataLedgerContractCaller) Ledger(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Uri  string
	Addr common.Address
}, error) {
	var out []interface{}
	err := _DataLedgerContract.contract.Call(opts, &out, "ledger", arg0)

	outstruct := new(struct {
		Uri  string
		Addr common.Address
	})

	outstruct.Uri = out[0].(string)
	outstruct.Addr = out[1].(common.Address)

	return *outstruct, err

}

// Ledger is a free data retrieval call binding the contract method 0x15977d45.
//
// Solidity: function ledger(bytes32 ) view returns(string uri, address addr)
func (_DataLedgerContract *DataLedgerContractSession) Ledger(arg0 [32]byte) (struct {
	Uri  string
	Addr common.Address
}, error) {
	return _DataLedgerContract.Contract.Ledger(&_DataLedgerContract.CallOpts, arg0)
}

// Ledger is a free data retrieval call binding the contract method 0x15977d45.
//
// Solidity: function ledger(bytes32 ) view returns(string uri, address addr)
func (_DataLedgerContract *DataLedgerContractCallerSession) Ledger(arg0 [32]byte) (struct {
	Uri  string
	Addr common.Address
}, error) {
	return _DataLedgerContract.Contract.Ledger(&_DataLedgerContract.CallOpts, arg0)
}

// TopicExists is a free data retrieval call binding the contract method 0x78b5749d.
//
// Solidity: function topicExists(string topic) view returns(bool)
func (_DataLedgerContract *DataLedgerContractCaller) TopicExists(opts *bind.CallOpts, topic string) (bool, error) {
	var out []interface{}
	err := _DataLedgerContract.contract.Call(opts, &out, "topicExists", topic)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TopicExists is a free data retrieval call binding the contract method 0x78b5749d.
//
// Solidity: function topicExists(string topic) view returns(bool)
func (_DataLedgerContract *DataLedgerContractSession) TopicExists(topic string) (bool, error) {
	return _DataLedgerContract.Contract.TopicExists(&_DataLedgerContract.CallOpts, topic)
}

// TopicExists is a free data retrieval call binding the contract method 0x78b5749d.
//
// Solidity: function topicExists(string topic) view returns(bool)
func (_DataLedgerContract *DataLedgerContractCallerSession) TopicExists(topic string) (bool, error) {
	return _DataLedgerContract.Contract.TopicExists(&_DataLedgerContract.CallOpts, topic)
}

// AddTopic is a paid mutator transaction binding the contract method 0x3eeb446e.
//
// Solidity: function addTopic(string topic, bytes32 key) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) AddTopic(opts *bind.TransactOpts, topic string, key [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "addTopic", topic, key)
}

// AddTopic is a paid mutator transaction binding the contract method 0x3eeb446e.
//
// Solidity: function addTopic(string topic, bytes32 key) returns()
func (_DataLedgerContract *DataLedgerContractSession) AddTopic(topic string, key [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.AddTopic(&_DataLedgerContract.TransactOpts, topic, key)
}

// AddTopic is a paid mutator transaction binding the contract method 0x3eeb446e.
//
// Solidity: function addTopic(string topic, bytes32 key) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) AddTopic(topic string, key [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.AddTopic(&_DataLedgerContract.TransactOpts, topic, key)
}

// DeleteMeasurement is a paid mutator transaction binding the contract method 0x77ad95ca.
//
// Solidity: function deleteMeasurement(bytes32 hash) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) DeleteMeasurement(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "deleteMeasurement", hash)
}

// DeleteMeasurement is a paid mutator transaction binding the contract method 0x77ad95ca.
//
// Solidity: function deleteMeasurement(bytes32 hash) returns()
func (_DataLedgerContract *DataLedgerContractSession) DeleteMeasurement(hash [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DeleteMeasurement(&_DataLedgerContract.TransactOpts, hash)
}

// DeleteMeasurement is a paid mutator transaction binding the contract method 0x77ad95ca.
//
// Solidity: function deleteMeasurement(bytes32 hash) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) DeleteMeasurement(hash [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DeleteMeasurement(&_DataLedgerContract.TransactOpts, hash)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) SetAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "setAddress", _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_DataLedgerContract *DataLedgerContractSession) SetAddress(_address common.Address) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.SetAddress(&_DataLedgerContract.TransactOpts, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) SetAddress(_address common.Address) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.SetAddress(&_DataLedgerContract.TransactOpts, _address)
}

// StoreInfo is a paid mutator transaction binding the contract method 0x0335e024.
//
// Solidity: function storeInfo(bytes32 hash, string uri, string[] topic) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) StoreInfo(opts *bind.TransactOpts, hash [32]byte, uri string, topic []string) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "storeInfo", hash, uri, topic)
}

// StoreInfo is a paid mutator transaction binding the contract method 0x0335e024.
//
// Solidity: function storeInfo(bytes32 hash, string uri, string[] topic) returns()
func (_DataLedgerContract *DataLedgerContractSession) StoreInfo(hash [32]byte, uri string, topic []string) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.StoreInfo(&_DataLedgerContract.TransactOpts, hash, uri, topic)
}

// StoreInfo is a paid mutator transaction binding the contract method 0x0335e024.
//
// Solidity: function storeInfo(bytes32 hash, string uri, string[] topic) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) StoreInfo(hash [32]byte, uri string, topic []string) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.StoreInfo(&_DataLedgerContract.TransactOpts, hash, uri, topic)
}

// DataLedgerContractDeleteInfoIterator is returned from FilterDeleteInfo and is used to iterate over the raw logs and unpacked data for DeleteInfo events raised by the DataLedgerContract contract.
type DataLedgerContractDeleteInfoIterator struct {
	Event *DataLedgerContractDeleteInfo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataLedgerContractDeleteInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLedgerContractDeleteInfo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataLedgerContractDeleteInfo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataLedgerContractDeleteInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLedgerContractDeleteInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLedgerContractDeleteInfo represents a DeleteInfo event raised by the DataLedgerContract contract.
type DataLedgerContractDeleteInfo struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeleteInfo is a free log retrieval operation binding the contract event 0x072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe.
//
// Solidity: event deleteInfo(bytes32 indexed _hash)
func (_DataLedgerContract *DataLedgerContractFilterer) FilterDeleteInfo(opts *bind.FilterOpts, _hash [][32]byte) (*DataLedgerContractDeleteInfoIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.FilterLogs(opts, "deleteInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractDeleteInfoIterator{contract: _DataLedgerContract.contract, event: "deleteInfo", logs: logs, sub: sub}, nil
}

// WatchDeleteInfo is a free log subscription operation binding the contract event 0x072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe.
//
// Solidity: event deleteInfo(bytes32 indexed _hash)
func (_DataLedgerContract *DataLedgerContractFilterer) WatchDeleteInfo(opts *bind.WatchOpts, sink chan<- *DataLedgerContractDeleteInfo, _hash [][32]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.WatchLogs(opts, "deleteInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLedgerContractDeleteInfo)
				if err := _DataLedgerContract.contract.UnpackLog(event, "deleteInfo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeleteInfo is a log parse operation binding the contract event 0x072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe.
//
// Solidity: event deleteInfo(bytes32 indexed _hash)
func (_DataLedgerContract *DataLedgerContractFilterer) ParseDeleteInfo(log types.Log) (*DataLedgerContractDeleteInfo, error) {
	event := new(DataLedgerContractDeleteInfo)
	if err := _DataLedgerContract.contract.UnpackLog(event, "deleteInfo", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DataLedgerContractEvtStoreInfoIterator is returned from FilterEvtStoreInfo and is used to iterate over the raw logs and unpacked data for EvtStoreInfo events raised by the DataLedgerContract contract.
type DataLedgerContractEvtStoreInfoIterator struct {
	Event *DataLedgerContractEvtStoreInfo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataLedgerContractEvtStoreInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLedgerContractEvtStoreInfo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataLedgerContractEvtStoreInfo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataLedgerContractEvtStoreInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLedgerContractEvtStoreInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLedgerContractEvtStoreInfo represents a EvtStoreInfo event raised by the DataLedgerContract contract.
type DataLedgerContractEvtStoreInfo struct {
	Hash   [32]byte
	Uri    string
	Topics []string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEvtStoreInfo is a free log retrieval operation binding the contract event 0x717b8d30e5c9a0e4112f8c62e289d73e44dc76c0e1254192334f3735ea4a8655.
//
// Solidity: event evtStoreInfo(bytes32 indexed _hash, string _uri, string[] _topics)
func (_DataLedgerContract *DataLedgerContractFilterer) FilterEvtStoreInfo(opts *bind.FilterOpts, _hash [][32]byte) (*DataLedgerContractEvtStoreInfoIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.FilterLogs(opts, "evtStoreInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractEvtStoreInfoIterator{contract: _DataLedgerContract.contract, event: "evtStoreInfo", logs: logs, sub: sub}, nil
}

// WatchEvtStoreInfo is a free log subscription operation binding the contract event 0x717b8d30e5c9a0e4112f8c62e289d73e44dc76c0e1254192334f3735ea4a8655.
//
// Solidity: event evtStoreInfo(bytes32 indexed _hash, string _uri, string[] _topics)
func (_DataLedgerContract *DataLedgerContractFilterer) WatchEvtStoreInfo(opts *bind.WatchOpts, sink chan<- *DataLedgerContractEvtStoreInfo, _hash [][32]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.WatchLogs(opts, "evtStoreInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLedgerContractEvtStoreInfo)
				if err := _DataLedgerContract.contract.UnpackLog(event, "evtStoreInfo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEvtStoreInfo is a log parse operation binding the contract event 0x717b8d30e5c9a0e4112f8c62e289d73e44dc76c0e1254192334f3735ea4a8655.
//
// Solidity: event evtStoreInfo(bytes32 indexed _hash, string _uri, string[] _topics)
func (_DataLedgerContract *DataLedgerContractFilterer) ParseEvtStoreInfo(log types.Log) (*DataLedgerContractEvtStoreInfo, error) {
	event := new(DataLedgerContractEvtStoreInfo)
	if err := _DataLedgerContract.contract.UnpackLog(event, "evtStoreInfo", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DataLedgerContractNewTopicIterator is returned from FilterNewTopic and is used to iterate over the raw logs and unpacked data for NewTopic events raised by the DataLedgerContract contract.
type DataLedgerContractNewTopicIterator struct {
	Event *DataLedgerContractNewTopic // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataLedgerContractNewTopicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLedgerContractNewTopic)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataLedgerContractNewTopic)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataLedgerContractNewTopicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLedgerContractNewTopicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLedgerContractNewTopic represents a NewTopic event raised by the DataLedgerContract contract.
type DataLedgerContractNewTopic struct {
	Topic common.Hash
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewTopic is a free log retrieval operation binding the contract event 0xbbced0d0b08bca32f3efec0aa77ad0d7d183168e8f976b456ad70129694c41a6.
//
// Solidity: event newTopic(string indexed _topic)
func (_DataLedgerContract *DataLedgerContractFilterer) FilterNewTopic(opts *bind.FilterOpts, _topic []string) (*DataLedgerContractNewTopicIterator, error) {

	var _topicRule []interface{}
	for _, _topicItem := range _topic {
		_topicRule = append(_topicRule, _topicItem)
	}

	logs, sub, err := _DataLedgerContract.contract.FilterLogs(opts, "newTopic", _topicRule)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractNewTopicIterator{contract: _DataLedgerContract.contract, event: "newTopic", logs: logs, sub: sub}, nil
}

// WatchNewTopic is a free log subscription operation binding the contract event 0xbbced0d0b08bca32f3efec0aa77ad0d7d183168e8f976b456ad70129694c41a6.
//
// Solidity: event newTopic(string indexed _topic)
func (_DataLedgerContract *DataLedgerContractFilterer) WatchNewTopic(opts *bind.WatchOpts, sink chan<- *DataLedgerContractNewTopic, _topic []string) (event.Subscription, error) {

	var _topicRule []interface{}
	for _, _topicItem := range _topic {
		_topicRule = append(_topicRule, _topicItem)
	}

	logs, sub, err := _DataLedgerContract.contract.WatchLogs(opts, "newTopic", _topicRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLedgerContractNewTopic)
				if err := _DataLedgerContract.contract.UnpackLog(event, "newTopic", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTopic is a log parse operation binding the contract event 0xbbced0d0b08bca32f3efec0aa77ad0d7d183168e8f976b456ad70129694c41a6.
//
// Solidity: event newTopic(string indexed _topic)
func (_DataLedgerContract *DataLedgerContractFilterer) ParseNewTopic(log types.Log) (*DataLedgerContractNewTopic, error) {
	event := new(DataLedgerContractNewTopic)
	if err := _DataLedgerContract.contract.UnpackLog(event, "newTopic", log); err != nil {
		return nil, err
	}
	return event, nil
}
