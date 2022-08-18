// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balanceContract

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

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dummyAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"setTotalSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20BasicFuncSigs maps the 4-byte function signature to its string representation.
var ERC20BasicFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"0e2a4417": "dummyAccount()",
	"06fdde03": "name()",
	"f7ea7a3d": "setTotalSupply(uint256)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
var ERC20BasicBin = "0x6080604052600380546001600160a01b0319908116737fff1f93d22e9b2aa4c6a536531950ed386bd95e1790915560048054909116732ce7526814b4c895b98e4605651e5ac8ed1667d617905534801561005857600080fd5b506106c4806100686000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063313ce56711610071578063313ce567146101df57806370a08231146101fd57806395d89b4114610223578063a9059cbb1461022b578063dd62ed3e14610257578063f7ea7a3d14610285576100a9565b806306fdde03146100ae578063095ea7b31461012b5780630e2a44171461016b57806318160ddd1461018f57806323b872dd146101a9575b600080fd5b6100b66102a4565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100f05781810151838201526020016100d8565b50505050905090810190601f16801561011d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101576004803603604081101561014157600080fd5b506001600160a01b0381351690602001356102ca565b604080519115158252519081900360200190f35b610173610330565b604080516001600160a01b039092168252519081900360200190f35b61019761033f565b60408051918252519081900360200190f35b610157600480360360608110156101bf57600080fd5b506001600160a01b03813581169160208101359091169060400135610345565b6101e76104a0565b6040805160ff9092168252519081900360200190f35b6101976004803603602081101561021357600080fd5b50356001600160a01b03166104a5565b6100b66104c0565b6101576004803603604081101561024157600080fd5b506001600160a01b0381351690602001356104df565b6101976004803603604081101561026d57600080fd5b506001600160a01b03813581169160200135166105a9565b6102a26004803603602081101561029b57600080fd5b50356105d4565b005b6040518060400160405280600a8152602001694552433230426173696360b01b81525081565b3360008181526001602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b6004546001600160a01b031681565b60025490565b6001600160a01b03831660009081526020819052604081205482111561036a57600080fd5b6001600160a01b038416600090815260016020908152604080832033845290915290205482111561039a57600080fd5b6001600160a01b0384166000908152602081905260409020546103c3908363ffffffff61063416565b6001600160a01b0385166000908152602081815260408083209390935560018152828220338352905220546103fe908363ffffffff61063416565b6001600160a01b0380861660009081526001602090815260408083203384528252808320949094559186168152908190522054610441908363ffffffff61064616565b6001600160a01b038085166000818152602081815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35060019392505050565b601281565b6001600160a01b031660009081526020819052604090205490565b6040518060400160405280600381526020016242534360e81b81525081565b336000908152602081905260408120548211156104fb57600080fd5b3360009081526020819052604090205461051b908363ffffffff61063416565b33600090815260208190526040808220929092556001600160a01b0385168152205461054d908363ffffffff61064616565b6001600160a01b038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6003546001600160a01b0316331461061d5760405162461bcd60e51b815260040180806020018281038252603381526020018061065d6033913960400191505060405180910390fd5b600281905533600090815260208190526040902055565b60008282111561064057fe5b50900390565b60008282018381101561065557fe5b939250505056fe596f7520646f206e6f74206861766520656e6f7567682070726976696c6567657320746f20646f207468697320616374696f6ea265627a7a7231582098a6f64e670ac5dbdd9841657959d529506f6973a81bbec17546440c68d8334e64736f6c63430005100032"

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) Allowance(opts *bind.CallOpts, owner common.Address, delegate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Basic.contract.Call(opts, &out, "allowance", owner, delegate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_ERC20Basic *ERC20BasicSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.Allowance(&_ERC20Basic.CallOpts, owner, delegate)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.Allowance(&_ERC20Basic.CallOpts, owner, delegate)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Basic.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Basic *ERC20BasicCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Basic.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Basic *ERC20BasicSession) Decimals() (uint8, error) {
	return _ERC20Basic.Contract.Decimals(&_ERC20Basic.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Basic *ERC20BasicCallerSession) Decimals() (uint8, error) {
	return _ERC20Basic.Contract.Decimals(&_ERC20Basic.CallOpts)
}

// DummyAccount is a free data retrieval call binding the contract method 0x0e2a4417.
//
// Solidity: function dummyAccount() view returns(address)
func (_ERC20Basic *ERC20BasicCaller) DummyAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Basic.contract.Call(opts, &out, "dummyAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DummyAccount is a free data retrieval call binding the contract method 0x0e2a4417.
//
// Solidity: function dummyAccount() view returns(address)
func (_ERC20Basic *ERC20BasicSession) DummyAccount() (common.Address, error) {
	return _ERC20Basic.Contract.DummyAccount(&_ERC20Basic.CallOpts)
}

// DummyAccount is a free data retrieval call binding the contract method 0x0e2a4417.
//
// Solidity: function dummyAccount() view returns(address)
func (_ERC20Basic *ERC20BasicCallerSession) DummyAccount() (common.Address, error) {
	return _ERC20Basic.Contract.DummyAccount(&_ERC20Basic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Basic *ERC20BasicCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Basic.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Basic *ERC20BasicSession) Name() (string, error) {
	return _ERC20Basic.Contract.Name(&_ERC20Basic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Basic *ERC20BasicCallerSession) Name() (string, error) {
	return _ERC20Basic.Contract.Name(&_ERC20Basic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Basic *ERC20BasicCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Basic.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Basic *ERC20BasicSession) Symbol() (string, error) {
	return _ERC20Basic.Contract.Symbol(&_ERC20Basic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Basic *ERC20BasicCallerSession) Symbol() (string, error) {
	return _ERC20Basic.Contract.Symbol(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Basic.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Approve(opts *bind.TransactOpts, delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "approve", delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Approve(delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Approve(&_ERC20Basic.TransactOpts, delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Approve(delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Approve(&_ERC20Basic.TransactOpts, delegate, numTokens)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(uint256 total) returns()
func (_ERC20Basic *ERC20BasicTransactor) SetTotalSupply(opts *bind.TransactOpts, total *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "setTotalSupply", total)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(uint256 total) returns()
func (_ERC20Basic *ERC20BasicSession) SetTotalSupply(total *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.SetTotalSupply(&_ERC20Basic.TransactOpts, total)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(uint256 total) returns()
func (_ERC20Basic *ERC20BasicTransactorSession) SetTotalSupply(total *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.SetTotalSupply(&_ERC20Basic.TransactOpts, total)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, receiver, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) TransferFrom(opts *bind.TransactOpts, owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transferFrom", owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_ERC20Basic *ERC20BasicSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.TransferFrom(&_ERC20Basic.TransactOpts, owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.TransferFrom(&_ERC20Basic.TransactOpts, owner, buyer, numTokens)
}

// ERC20BasicApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Basic contract.
type ERC20BasicApprovalIterator struct {
	Event *ERC20BasicApproval // Event containing the contract specifics and raw log

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
func (it *ERC20BasicApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicApproval)
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
		it.Event = new(ERC20BasicApproval)
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
func (it *ERC20BasicApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicApproval represents a Approval event raised by the ERC20Basic contract.
type ERC20BasicApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_ERC20Basic *ERC20BasicFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*ERC20BasicApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicApprovalIterator{contract: _ERC20Basic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_ERC20Basic *ERC20BasicFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20BasicApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicApproval)
				if err := _ERC20Basic.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_ERC20Basic *ERC20BasicFilterer) ParseApproval(log types.Log) (*ERC20BasicApproval, error) {
	event := new(ERC20BasicApproval)
	if err := _ERC20Basic.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_ERC20Basic *ERC20BasicFilterer) ParseTransfer(log types.Log) (*ERC20BasicTransfer, error) {
	event := new(ERC20BasicTransfer)
	if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820eef8a741f3d77210640552a441615adc647a4691afaca42fe19846c192d985dd64736f6c63430005100032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// BalanceContractABI is the input ABI used to generate the binding from.
const BalanceContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"AdquireTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"CompletePurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"PriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"PurchaseRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"RequestPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_topic\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"Subscribe\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"completePurchase\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dummyAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getPriceMeasurement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"getPriceSubscription\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"isSubscribed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"purchaseMeasurement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"purchaseSubscription\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"revokeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"sendTokenToClient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPriceToMeasurement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"setTotalSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BalanceContractFuncSigs maps the 4-byte function signature to its string representation.
var BalanceContractFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"dd5bab86": "completePurchase(bytes32,address,bytes32)",
	"313ce567": "decimals()",
	"0e2a4417": "dummyAccount()",
	"325655b2": "getPriceMeasurement(bytes32)",
	"e7ca2b9e": "getPriceSubscription(string)",
	"7a1a14f5": "isSubscribed(address,string)",
	"06fdde03": "name()",
	"96b4ae3c": "purchaseMeasurement(bytes32)",
	"1d9b30c7": "purchaseSubscription(string)",
	"7b2cf65c": "revokeTransaction(bytes32,address)",
	"7d7cb5cc": "sendTokenToClient(address,uint256)",
	"e30081a0": "setAddress(address)",
	"c8a7d28f": "setPriceToMeasurement(bytes32,uint256)",
	"f7ea7a3d": "setTotalSupply(uint256)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// BalanceContractBin is the compiled bytecode used for deploying new contracts.
var BalanceContractBin = "0x6080604052600380546001600160a01b0319908116737fff1f93d22e9b2aa4c6a536531950ed386bd95e1790915560048054909116732ce7526814b4c895b98e4605651e5ac8ed1667d617905534801561005857600080fd5b50611688806100686000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c80637b2cf65c116100b8578063c8a7d28f1161007c578063c8a7d28f146104d5578063dd5bab86146104f8578063dd62ed3e1461052a578063e30081a014610558578063e7ca2b9e1461057e578063f7ea7a3d1461062457610137565b80637b2cf65c1461042c5780637d7cb5cc1461045857806395d89b411461048457806396b4ae3c1461048c578063a9059cbb146104a957610137565b806323b872dd116100ff57806323b872dd146102df578063313ce56714610315578063325655b21461033357806370a08231146103505780637a1a14f51461037657610137565b806306fdde031461013c578063095ea7b3146101b95780630e2a4417146101f957806318160ddd1461021d5780631d9b30c714610237575b600080fd5b610144610641565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561017e578181015183820152602001610166565b50505050905090810190601f1680156101ab5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101e5600480360360408110156101cf57600080fd5b506001600160a01b038135169060200135610667565b604080519115158252519081900360200190f35b6102016106cd565b604080516001600160a01b039092168252519081900360200190f35b6102256106dc565b60408051918252519081900360200190f35b6102dd6004803603602081101561024d57600080fd5b81019060208101813564010000000081111561026857600080fd5b82018360208201111561027a57600080fd5b8035906020019184600183028401116401000000008311171561029c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106e2945050505050565b005b6101e5600480360360608110156102f557600080fd5b506001600160a01b03813581169160208101359091169060400135610985565b61031d610ae0565b6040805160ff9092168252519081900360200190f35b6102256004803603602081101561034957600080fd5b5035610ae5565b6102256004803603602081101561036657600080fd5b50356001600160a01b0316610afa565b6101e56004803603604081101561038c57600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156103b757600080fd5b8201836020820111156103c957600080fd5b803590602001918460018302840111640100000000831117156103eb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b15945050505050565b6102dd6004803603604081101561044257600080fd5b50803590602001356001600160a01b0316610c23565b6102dd6004803603604081101561046e57600080fd5b506001600160a01b038135169060200135610d19565b610144610e19565b6102dd600480360360208110156104a257600080fd5b5035610e38565b6101e5600480360360408110156104bf57600080fd5b506001600160a01b038135169060200135610fe6565b6102dd600480360360408110156104eb57600080fd5b50803590602001356110b0565b6102dd6004803603606081101561050e57600080fd5b508035906001600160a01b0360208201351690604001356111bd565b6102256004803603604081101561054057600080fd5b506001600160a01b0381358116916020013516611321565b6102dd6004803603602081101561056e57600080fd5b50356001600160a01b031661134c565b6102256004803603602081101561059457600080fd5b8101906020810181356401000000008111156105af57600080fd5b8201836020820111156105c157600080fd5b803590602001918460018302840111640100000000831117156105e357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506113b7945050505050565b6102dd6004803603602081101561063a57600080fd5b50356114a0565b6040518060400160405280600a8152602001694552433230426173696360b01b81525081565b3360008181526001602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b6004546001600160a01b031681565b60025490565b60006106ed826113b7565b905060008111610744576040805162461bcd60e51b815260206004820152601860248201527f54686520746f70696320646f6573206e6f742065786973740000000000000000604482015290519081900360640190fd5b8061074e33610afa565b101561078b5760405162461bcd60e51b81526004018080602001828103825260368152602001806115296036913960400191505060405180910390fd5b6005546040516355b4070160e11b81526020600482018181528551602484015285516000946001600160a01b03169363ab680e02938893928392604401918501908083838b5b838110156107e95781810151838201526020016107d1565b50505050905090810190601f1680156108165780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b15801561083357600080fd5b505afa158015610847573d6000803e3d6000fd5b505050506040513d602081101561085d57600080fd5b5051905061086b8183610fe6565b506008836040518082805190602001908083835b6020831061089e5780518252601f19909201916020918201910161087f565b51815160209384036101000a600019018019909216911617905292019485525060405193849003810184208054600181018255600091825290829020018054336001600160a01b0319909116811790915587516001600160a01b0387169591945088935090918291908401908083835b6020831061092d5780518252601f19909201916020918201910161090e565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507fc36d692ea82425e7776ac80f320816b8cbd7b02289e0d83d5ef5c489af7355ca92506000919050a4505050565b6001600160a01b0383166000908152602081905260408120548211156109aa57600080fd5b6001600160a01b03841660009081526001602090815260408083203384529091529020548211156109da57600080fd5b6001600160a01b038416600090815260208190526040902054610a03908363ffffffff61150016565b6001600160a01b038516600090815260208181526040808320939093556001815282822033835290522054610a3e908363ffffffff61150016565b6001600160a01b0380861660009081526001602090815260408083203384528252808320949094559186168152908190522054610a81908363ffffffff61151216565b6001600160a01b038085166000818152602081815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35060019392505050565b601281565b6000818152600660205260409020545b919050565b6001600160a01b031660009081526020819052604090205490565b600060606008836040518082805190602001908083835b60208310610b4b5780518252601f199092019160209182019101610b2c565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201852080548084028701840190925281865293509150830182828015610bc357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610ba5575b5050835193945060009250829150505b82811015610c1957866001600160a01b0316848281518110610bf157fe5b60200260200101516001600160a01b03161415610c115760019150610c19565b600101610bd3565b5095945050505050565b6003546001600160a01b03163314610c6c5760405162461bcd60e51b81526004018080602001828103825260338152602001806115976033913960400191505060405180910390fd5b60008281526007602090815260408083206001600160a01b038581168552925290912060018101549054600454919290811691610cab91168484610985565b5060008481526007602090815260408083206001600160a01b038088168086529190935281842080546001600160a01b0319168155600101849055905191841692909187917fe4f046578a8ec0e49030155ac1ec29d863bb3e723e3a429811634562ef141cc091a450505050565b6003546001600160a01b03163314610d625760405162461bcd60e51b81526004018080602001828103825260218152602001806115f66021913960400191505060405180910390fd5b600354610d77906001600160a01b0316610afa565b811115610dcb576040805162461bcd60e51b815260206004820152601b60248201527f546865726520617265206e6f7420656e6f75676820746f6b656e730000000000604482015290519081900360640190fd5b610dd58282610fe6565b506040805182815290516001600160a01b038416917f22a6cb4fc5d61567d5b2a20f88a7049394dcac13b3834916f402f66a6600346c919081900360200190a25050565b6040518060400160405280600381526020016242534360e81b81525081565b6000610e4382610ae5565b60055460408051636ade021960e01b81526004810186905290519293506000926001600160a01b0390921691636ade021991602480820192602092909190829003018186803b158015610e9557600080fd5b505afa158015610ea9573d6000803e3d6000fd5b505050506040513d6020811015610ebf57600080fd5b5051905081610f15576040805162461bcd60e51b815260206004820152601760248201527f546865206461746120646f6573206e6f74206578697374000000000000000000604482015290519081900360640190fd5b81610f1f33610afa565b1015610f5c5760405162461bcd60e51b81526004018080602001828103825260368152602001806115296036913960400191505060405180910390fd5b6000838152600760209081526040808320338452909152902080546001600160a01b0319166001600160a01b038381169190911782556001909101839055600454610fa8911683610fe6565b506040516001600160a01b03821690339085907fd457a19a91892104bd37ed9c2de0d3398da0d5f42bca745436a5c9a17ab70e4190600090a4505050565b3360009081526020819052604081205482111561100257600080fd5b33600090815260208190526040902054611022908363ffffffff61150016565b33600090815260208190526040808220929092556001600160a01b03851681522054611054908363ffffffff61151216565b6001600160a01b038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b60055460408051636ade021960e01b81526004810185905290516000926001600160a01b031691636ade0219916024808301926020929190829003018186803b1580156110fc57600080fd5b505afa158015611110573d6000803e3d6000fd5b505050506040513d602081101561112657600080fd5b50519050336001600160a01b038216146111715760405162461bcd60e51b815260040180806020018281038252603881526020018061155f6038913960400191505060405180910390fd5b6000838152600660209081526040918290208490558151848152915185927fee47534b2400bc7be3fbdc39f0283b8643fc472a30ffd7324e6161b4b2b91f2492908290030190a2505050565b6003546001600160a01b031633146112065760405162461bcd60e51b81526004018080602001828103825260338152602001806115976033913960400191505060405180910390fd5b60008381526007602090815260408083206001600160a01b03861684529091529020600101546112675760405162461bcd60e51b815260040180806020018281038252603d815260200180611617603d913960400191505060405180910390fd5b60008381526007602090815260408083206001600160a01b03868116855292529091208054600454600190920154908316926112a69216908390610985565b5060008481526007602090815260408083206001600160a01b0380881680865291845282852080546001600160a01b0319168155600101949094558151868152915193851693909288927fbedfb3b4f53acc8876b63deec056b5a4252ea296380b4d55473f7d6f2f76889c929081900390910190a450505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6003546001600160a01b031633146113955760405162461bcd60e51b815260040180806020018281038252602c8152602001806115ca602c913960400191505060405180910390fd5b600580546001600160a01b0319166001600160a01b0392909216919091179055565b6005546040516378b5749d60e01b81526020600482018181528451602484015284516000946001600160a01b0316936378b5749d938793928392604401918501908083838b5b838110156114155781810151838201526020016113fd565b50505050905090810190601f1680156114425780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b15801561145f57600080fd5b505afa158015611473573d6000803e3d6000fd5b505050506040513d602081101561148957600080fd5b5051156114985750600a610af5565b506000610af5565b6003546001600160a01b031633146114e95760405162461bcd60e51b81526004018080602001828103825260338152602001806115976033913960400191505060405180910390fd5b600281905533600090815260208190526040902055565b60008282111561150c57fe5b50900390565b60008282018381101561152157fe5b939250505056fe596f7520646f206e6f74206861766520656e6f75676820746f6b656e7320746f20636f6d706c657465207468652070757263686173654f6e6c7920746865206163636f756e742077686f20696e7365727465642074686520646174612063616e2073657420697473207072696365596f7520646f206e6f74206861766520656e6f7567682070726976696c6567657320746f20646f207468697320616374696f6e596f7520646f206e6f7420686176652070726976696c6567657320746f20646f207468697320616374696f6e596f7520646f206e6f74206861766520656e6f7567682070726976696c656765735468657265206973206e6f74207472616e73616374696f6e206173736f636961746520746f207468697320686173682062792074686973206275796572a265627a7a72315820adc0c4cff1f8707ef4540ab38aa2daa8701c5ef73b4bd762310edd8136c71d3564736f6c63430005100032"

// DeployBalanceContract deploys a new Ethereum contract, binding an instance of BalanceContract to it.
func DeployBalanceContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BalanceContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BalanceContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BalanceContract{BalanceContractCaller: BalanceContractCaller{contract: contract}, BalanceContractTransactor: BalanceContractTransactor{contract: contract}, BalanceContractFilterer: BalanceContractFilterer{contract: contract}}, nil
}

// BalanceContract is an auto generated Go binding around an Ethereum contract.
type BalanceContract struct {
	BalanceContractCaller     // Read-only binding to the contract
	BalanceContractTransactor // Write-only binding to the contract
	BalanceContractFilterer   // Log filterer for contract events
}

// BalanceContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceContractSession struct {
	Contract     *BalanceContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceContractCallerSession struct {
	Contract *BalanceContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BalanceContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceContractTransactorSession struct {
	Contract     *BalanceContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BalanceContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceContractRaw struct {
	Contract *BalanceContract // Generic contract binding to access the raw methods on
}

// BalanceContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceContractCallerRaw struct {
	Contract *BalanceContractCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceContractTransactorRaw struct {
	Contract *BalanceContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalanceContract creates a new instance of BalanceContract, bound to a specific deployed contract.
func NewBalanceContract(address common.Address, backend bind.ContractBackend) (*BalanceContract, error) {
	contract, err := bindBalanceContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceContract{BalanceContractCaller: BalanceContractCaller{contract: contract}, BalanceContractTransactor: BalanceContractTransactor{contract: contract}, BalanceContractFilterer: BalanceContractFilterer{contract: contract}}, nil
}

// NewBalanceContractCaller creates a new read-only instance of BalanceContract, bound to a specific deployed contract.
func NewBalanceContractCaller(address common.Address, caller bind.ContractCaller) (*BalanceContractCaller, error) {
	contract, err := bindBalanceContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceContractCaller{contract: contract}, nil
}

// NewBalanceContractTransactor creates a new write-only instance of BalanceContract, bound to a specific deployed contract.
func NewBalanceContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceContractTransactor, error) {
	contract, err := bindBalanceContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceContractTransactor{contract: contract}, nil
}

// NewBalanceContractFilterer creates a new log filterer instance of BalanceContract, bound to a specific deployed contract.
func NewBalanceContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceContractFilterer, error) {
	contract, err := bindBalanceContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceContractFilterer{contract: contract}, nil
}

// bindBalanceContract binds a generic wrapper to an already deployed contract.
func bindBalanceContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceContract *BalanceContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceContract.Contract.BalanceContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceContract *BalanceContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceContract.Contract.BalanceContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceContract *BalanceContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceContract.Contract.BalanceContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceContract *BalanceContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceContract *BalanceContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceContract *BalanceContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceContract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_BalanceContract *BalanceContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, delegate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalanceContract.contract.Call(opts, &out, "allowance", owner, delegate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_BalanceContract *BalanceContractSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _BalanceContract.Contract.Allowance(&_BalanceContract.CallOpts, owner, delegate)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_BalanceContract *BalanceContractCallerSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _BalanceContract.Contract.Allowance(&_BalanceContract.CallOpts, owner, delegate)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_BalanceContract *BalanceContractCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalanceContract.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_BalanceContract *BalanceContractSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _BalanceContract.Contract.BalanceOf(&_BalanceContract.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_BalanceContract *BalanceContractCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _BalanceContract.Contract.BalanceOf(&_BalanceContract.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalanceContract *BalanceContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BalanceContract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalanceContract *BalanceContractSession) Decimals() (uint8, error) {
	return _BalanceContract.Contract.Decimals(&_BalanceContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalanceContract *BalanceContractCallerSession) Decimals() (uint8, error) {
	return _BalanceContract.Contract.Decimals(&_BalanceContract.CallOpts)
}

// DummyAccount is a free data retrieval call binding the contract method 0x0e2a4417.
//
// Solidity: function dummyAccount() view returns(address)
func (_BalanceContract *BalanceContractCaller) DummyAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalanceContract.contract.Call(opts, &out, "dummyAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DummyAccount is a free data retrieval call binding the contract method 0x0e2a4417.
//
// Solidity: function dummyAccount() view returns(address)
func (_BalanceContract *BalanceContractSession) DummyAccount() (common.Address, error) {
	return _BalanceContract.Contract.DummyAccount(&_BalanceContract.CallOpts)
}

// DummyAccount is a free data retrieval call binding the contract method 0x0e2a4417.
//
// Solidity: function dummyAccount() view returns(address)
func (_BalanceContract *BalanceContractCallerSession) DummyAccount() (common.Address, error) {
	return _BalanceContract.Contract.DummyAccount(&_BalanceContract.CallOpts)
}

// GetPriceMeasurement is a free data retrieval call binding the contract method 0x325655b2.
//
// Solidity: function getPriceMeasurement(bytes32 hash) view returns(uint256)
func (_BalanceContract *BalanceContractCaller) GetPriceMeasurement(opts *bind.CallOpts, hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BalanceContract.contract.Call(opts, &out, "getPriceMeasurement", hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceMeasurement is a free data retrieval call binding the contract method 0x325655b2.
//
// Solidity: function getPriceMeasurement(bytes32 hash) view returns(uint256)
func (_BalanceContract *BalanceContractSession) GetPriceMeasurement(hash [32]byte) (*big.Int, error) {
	return _BalanceContract.Contract.GetPriceMeasurement(&_BalanceContract.CallOpts, hash)
}

// GetPriceMeasurement is a free data retrieval call binding the contract method 0x325655b2.
//
// Solidity: function getPriceMeasurement(bytes32 hash) view returns(uint256)
func (_BalanceContract *BalanceContractCallerSession) GetPriceMeasurement(hash [32]byte) (*big.Int, error) {
	return _BalanceContract.Contract.GetPriceMeasurement(&_BalanceContract.CallOpts, hash)
}

// GetPriceSubscription is a free data retrieval call binding the contract method 0xe7ca2b9e.
//
// Solidity: function getPriceSubscription(string topic) view returns(uint256)
func (_BalanceContract *BalanceContractCaller) GetPriceSubscription(opts *bind.CallOpts, topic string) (*big.Int, error) {
	var out []interface{}
	err := _BalanceContract.contract.Call(opts, &out, "getPriceSubscription", topic)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceSubscription is a free data retrieval call binding the contract method 0xe7ca2b9e.
//
// Solidity: function getPriceSubscription(string topic) view returns(uint256)
func (_BalanceContract *BalanceContractSession) GetPriceSubscription(topic string) (*big.Int, error) {
	return _BalanceContract.Contract.GetPriceSubscription(&_BalanceContract.CallOpts, topic)
}

// GetPriceSubscription is a free data retrieval call binding the contract method 0xe7ca2b9e.
//
// Solidity: function getPriceSubscription(string topic) view returns(uint256)
func (_BalanceContract *BalanceContractCallerSession) GetPriceSubscription(topic string) (*big.Int, error) {
	return _BalanceContract.Contract.GetPriceSubscription(&_BalanceContract.CallOpts, topic)
}

// IsSubscribed is a free data retrieval call binding the contract method 0x7a1a14f5.
//
// Solidity: function isSubscribed(address client, string topic) view returns(bool)
func (_BalanceContract *BalanceContractCaller) IsSubscribed(opts *bind.CallOpts, client common.Address, topic string) (bool, error) {
	var out []interface{}
	err := _BalanceContract.contract.Call(opts, &out, "isSubscribed", client, topic)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSubscribed is a free data retrieval call binding the contract method 0x7a1a14f5.
//
// Solidity: function isSubscribed(address client, string topic) view returns(bool)
func (_BalanceContract *BalanceContractSession) IsSubscribed(client common.Address, topic string) (bool, error) {
	return _BalanceContract.Contract.IsSubscribed(&_BalanceContract.CallOpts, client, topic)
}

// IsSubscribed is a free data retrieval call binding the contract method 0x7a1a14f5.
//
// Solidity: function isSubscribed(address client, string topic) view returns(bool)
func (_BalanceContract *BalanceContractCallerSession) IsSubscribed(client common.Address, topic string) (bool, error) {
	return _BalanceContract.Contract.IsSubscribed(&_BalanceContract.CallOpts, client, topic)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalanceContract *BalanceContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalanceContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalanceContract *BalanceContractSession) Name() (string, error) {
	return _BalanceContract.Contract.Name(&_BalanceContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalanceContract *BalanceContractCallerSession) Name() (string, error) {
	return _BalanceContract.Contract.Name(&_BalanceContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalanceContract *BalanceContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalanceContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalanceContract *BalanceContractSession) Symbol() (string, error) {
	return _BalanceContract.Contract.Symbol(&_BalanceContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalanceContract *BalanceContractCallerSession) Symbol() (string, error) {
	return _BalanceContract.Contract.Symbol(&_BalanceContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalanceContract *BalanceContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalanceContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalanceContract *BalanceContractSession) TotalSupply() (*big.Int, error) {
	return _BalanceContract.Contract.TotalSupply(&_BalanceContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalanceContract *BalanceContractCallerSession) TotalSupply() (*big.Int, error) {
	return _BalanceContract.Contract.TotalSupply(&_BalanceContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_BalanceContract *BalanceContractTransactor) Approve(opts *bind.TransactOpts, delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "approve", delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_BalanceContract *BalanceContractSession) Approve(delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.Approve(&_BalanceContract.TransactOpts, delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_BalanceContract *BalanceContractTransactorSession) Approve(delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.Approve(&_BalanceContract.TransactOpts, delegate, numTokens)
}

// CompletePurchase is a paid mutator transaction binding the contract method 0xdd5bab86.
//
// Solidity: function completePurchase(bytes32 hash, address buyer, bytes32 txHash) returns()
func (_BalanceContract *BalanceContractTransactor) CompletePurchase(opts *bind.TransactOpts, hash [32]byte, buyer common.Address, txHash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "completePurchase", hash, buyer, txHash)
}

// CompletePurchase is a paid mutator transaction binding the contract method 0xdd5bab86.
//
// Solidity: function completePurchase(bytes32 hash, address buyer, bytes32 txHash) returns()
func (_BalanceContract *BalanceContractSession) CompletePurchase(hash [32]byte, buyer common.Address, txHash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.CompletePurchase(&_BalanceContract.TransactOpts, hash, buyer, txHash)
}

// CompletePurchase is a paid mutator transaction binding the contract method 0xdd5bab86.
//
// Solidity: function completePurchase(bytes32 hash, address buyer, bytes32 txHash) returns()
func (_BalanceContract *BalanceContractTransactorSession) CompletePurchase(hash [32]byte, buyer common.Address, txHash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.CompletePurchase(&_BalanceContract.TransactOpts, hash, buyer, txHash)
}

// PurchaseMeasurement is a paid mutator transaction binding the contract method 0x96b4ae3c.
//
// Solidity: function purchaseMeasurement(bytes32 hash) returns()
func (_BalanceContract *BalanceContractTransactor) PurchaseMeasurement(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "purchaseMeasurement", hash)
}

// PurchaseMeasurement is a paid mutator transaction binding the contract method 0x96b4ae3c.
//
// Solidity: function purchaseMeasurement(bytes32 hash) returns()
func (_BalanceContract *BalanceContractSession) PurchaseMeasurement(hash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.PurchaseMeasurement(&_BalanceContract.TransactOpts, hash)
}

// PurchaseMeasurement is a paid mutator transaction binding the contract method 0x96b4ae3c.
//
// Solidity: function purchaseMeasurement(bytes32 hash) returns()
func (_BalanceContract *BalanceContractTransactorSession) PurchaseMeasurement(hash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.PurchaseMeasurement(&_BalanceContract.TransactOpts, hash)
}

// PurchaseSubscription is a paid mutator transaction binding the contract method 0x1d9b30c7.
//
// Solidity: function purchaseSubscription(string topic) returns()
func (_BalanceContract *BalanceContractTransactor) PurchaseSubscription(opts *bind.TransactOpts, topic string) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "purchaseSubscription", topic)
}

// PurchaseSubscription is a paid mutator transaction binding the contract method 0x1d9b30c7.
//
// Solidity: function purchaseSubscription(string topic) returns()
func (_BalanceContract *BalanceContractSession) PurchaseSubscription(topic string) (*types.Transaction, error) {
	return _BalanceContract.Contract.PurchaseSubscription(&_BalanceContract.TransactOpts, topic)
}

// PurchaseSubscription is a paid mutator transaction binding the contract method 0x1d9b30c7.
//
// Solidity: function purchaseSubscription(string topic) returns()
func (_BalanceContract *BalanceContractTransactorSession) PurchaseSubscription(topic string) (*types.Transaction, error) {
	return _BalanceContract.Contract.PurchaseSubscription(&_BalanceContract.TransactOpts, topic)
}

// RevokeTransaction is a paid mutator transaction binding the contract method 0x7b2cf65c.
//
// Solidity: function revokeTransaction(bytes32 hash, address buyer) returns()
func (_BalanceContract *BalanceContractTransactor) RevokeTransaction(opts *bind.TransactOpts, hash [32]byte, buyer common.Address) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "revokeTransaction", hash, buyer)
}

// RevokeTransaction is a paid mutator transaction binding the contract method 0x7b2cf65c.
//
// Solidity: function revokeTransaction(bytes32 hash, address buyer) returns()
func (_BalanceContract *BalanceContractSession) RevokeTransaction(hash [32]byte, buyer common.Address) (*types.Transaction, error) {
	return _BalanceContract.Contract.RevokeTransaction(&_BalanceContract.TransactOpts, hash, buyer)
}

// RevokeTransaction is a paid mutator transaction binding the contract method 0x7b2cf65c.
//
// Solidity: function revokeTransaction(bytes32 hash, address buyer) returns()
func (_BalanceContract *BalanceContractTransactorSession) RevokeTransaction(hash [32]byte, buyer common.Address) (*types.Transaction, error) {
	return _BalanceContract.Contract.RevokeTransaction(&_BalanceContract.TransactOpts, hash, buyer)
}

// SendTokenToClient is a paid mutator transaction binding the contract method 0x7d7cb5cc.
//
// Solidity: function sendTokenToClient(address to, uint256 total) returns()
func (_BalanceContract *BalanceContractTransactor) SendTokenToClient(opts *bind.TransactOpts, to common.Address, total *big.Int) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "sendTokenToClient", to, total)
}

// SendTokenToClient is a paid mutator transaction binding the contract method 0x7d7cb5cc.
//
// Solidity: function sendTokenToClient(address to, uint256 total) returns()
func (_BalanceContract *BalanceContractSession) SendTokenToClient(to common.Address, total *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.SendTokenToClient(&_BalanceContract.TransactOpts, to, total)
}

// SendTokenToClient is a paid mutator transaction binding the contract method 0x7d7cb5cc.
//
// Solidity: function sendTokenToClient(address to, uint256 total) returns()
func (_BalanceContract *BalanceContractTransactorSession) SendTokenToClient(to common.Address, total *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.SendTokenToClient(&_BalanceContract.TransactOpts, to, total)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_BalanceContract *BalanceContractTransactor) SetAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "setAddress", _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_BalanceContract *BalanceContractSession) SetAddress(_address common.Address) (*types.Transaction, error) {
	return _BalanceContract.Contract.SetAddress(&_BalanceContract.TransactOpts, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_BalanceContract *BalanceContractTransactorSession) SetAddress(_address common.Address) (*types.Transaction, error) {
	return _BalanceContract.Contract.SetAddress(&_BalanceContract.TransactOpts, _address)
}

// SetPriceToMeasurement is a paid mutator transaction binding the contract method 0xc8a7d28f.
//
// Solidity: function setPriceToMeasurement(bytes32 hash, uint256 price) returns()
func (_BalanceContract *BalanceContractTransactor) SetPriceToMeasurement(opts *bind.TransactOpts, hash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "setPriceToMeasurement", hash, price)
}

// SetPriceToMeasurement is a paid mutator transaction binding the contract method 0xc8a7d28f.
//
// Solidity: function setPriceToMeasurement(bytes32 hash, uint256 price) returns()
func (_BalanceContract *BalanceContractSession) SetPriceToMeasurement(hash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.SetPriceToMeasurement(&_BalanceContract.TransactOpts, hash, price)
}

// SetPriceToMeasurement is a paid mutator transaction binding the contract method 0xc8a7d28f.
//
// Solidity: function setPriceToMeasurement(bytes32 hash, uint256 price) returns()
func (_BalanceContract *BalanceContractTransactorSession) SetPriceToMeasurement(hash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.SetPriceToMeasurement(&_BalanceContract.TransactOpts, hash, price)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(uint256 total) returns()
func (_BalanceContract *BalanceContractTransactor) SetTotalSupply(opts *bind.TransactOpts, total *big.Int) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "setTotalSupply", total)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(uint256 total) returns()
func (_BalanceContract *BalanceContractSession) SetTotalSupply(total *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.SetTotalSupply(&_BalanceContract.TransactOpts, total)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(uint256 total) returns()
func (_BalanceContract *BalanceContractTransactorSession) SetTotalSupply(total *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.SetTotalSupply(&_BalanceContract.TransactOpts, total)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_BalanceContract *BalanceContractTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "transfer", receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_BalanceContract *BalanceContractSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.Transfer(&_BalanceContract.TransactOpts, receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_BalanceContract *BalanceContractTransactorSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.Transfer(&_BalanceContract.TransactOpts, receiver, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_BalanceContract *BalanceContractTransactor) TransferFrom(opts *bind.TransactOpts, owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "transferFrom", owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_BalanceContract *BalanceContractSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.TransferFrom(&_BalanceContract.TransactOpts, owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_BalanceContract *BalanceContractTransactorSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.TransferFrom(&_BalanceContract.TransactOpts, owner, buyer, numTokens)
}

// BalanceContractAdquireTokensIterator is returned from FilterAdquireTokens and is used to iterate over the raw logs and unpacked data for AdquireTokens events raised by the BalanceContract contract.
type BalanceContractAdquireTokensIterator struct {
	Event *BalanceContractAdquireTokens // Event containing the contract specifics and raw log

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
func (it *BalanceContractAdquireTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractAdquireTokens)
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
		it.Event = new(BalanceContractAdquireTokens)
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
func (it *BalanceContractAdquireTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractAdquireTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractAdquireTokens represents a AdquireTokens event raised by the BalanceContract contract.
type BalanceContractAdquireTokens struct {
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAdquireTokens is a free log retrieval operation binding the contract event 0x22a6cb4fc5d61567d5b2a20f88a7049394dcac13b3834916f402f66a6600346c.
//
// Solidity: event AdquireTokens(address indexed _to, uint256 tokens)
func (_BalanceContract *BalanceContractFilterer) FilterAdquireTokens(opts *bind.FilterOpts, _to []common.Address) (*BalanceContractAdquireTokensIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "AdquireTokens", _toRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractAdquireTokensIterator{contract: _BalanceContract.contract, event: "AdquireTokens", logs: logs, sub: sub}, nil
}

// WatchAdquireTokens is a free log subscription operation binding the contract event 0x22a6cb4fc5d61567d5b2a20f88a7049394dcac13b3834916f402f66a6600346c.
//
// Solidity: event AdquireTokens(address indexed _to, uint256 tokens)
func (_BalanceContract *BalanceContractFilterer) WatchAdquireTokens(opts *bind.WatchOpts, sink chan<- *BalanceContractAdquireTokens, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "AdquireTokens", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractAdquireTokens)
				if err := _BalanceContract.contract.UnpackLog(event, "AdquireTokens", log); err != nil {
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

// ParseAdquireTokens is a log parse operation binding the contract event 0x22a6cb4fc5d61567d5b2a20f88a7049394dcac13b3834916f402f66a6600346c.
//
// Solidity: event AdquireTokens(address indexed _to, uint256 tokens)
func (_BalanceContract *BalanceContractFilterer) ParseAdquireTokens(log types.Log) (*BalanceContractAdquireTokens, error) {
	event := new(BalanceContractAdquireTokens)
	if err := _BalanceContract.contract.UnpackLog(event, "AdquireTokens", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalanceContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BalanceContract contract.
type BalanceContractApprovalIterator struct {
	Event *BalanceContractApproval // Event containing the contract specifics and raw log

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
func (it *BalanceContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractApproval)
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
		it.Event = new(BalanceContractApproval)
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
func (it *BalanceContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractApproval represents a Approval event raised by the BalanceContract contract.
type BalanceContractApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_BalanceContract *BalanceContractFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*BalanceContractApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractApprovalIterator{contract: _BalanceContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_BalanceContract *BalanceContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BalanceContractApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractApproval)
				if err := _BalanceContract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_BalanceContract *BalanceContractFilterer) ParseApproval(log types.Log) (*BalanceContractApproval, error) {
	event := new(BalanceContractApproval)
	if err := _BalanceContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalanceContractCompletePurchaseIterator is returned from FilterCompletePurchase and is used to iterate over the raw logs and unpacked data for CompletePurchase events raised by the BalanceContract contract.
type BalanceContractCompletePurchaseIterator struct {
	Event *BalanceContractCompletePurchase // Event containing the contract specifics and raw log

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
func (it *BalanceContractCompletePurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractCompletePurchase)
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
		it.Event = new(BalanceContractCompletePurchase)
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
func (it *BalanceContractCompletePurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractCompletePurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractCompletePurchase represents a CompletePurchase event raised by the BalanceContract contract.
type BalanceContractCompletePurchase struct {
	Hash   [32]byte
	From   common.Address
	To     common.Address
	TxHash [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCompletePurchase is a free log retrieval operation binding the contract event 0xbedfb3b4f53acc8876b63deec056b5a4252ea296380b4d55473f7d6f2f76889c.
//
// Solidity: event CompletePurchase(bytes32 indexed _hash, address indexed _from, address indexed _to, bytes32 _txHash)
func (_BalanceContract *BalanceContractFilterer) FilterCompletePurchase(opts *bind.FilterOpts, _hash [][32]byte, _from []common.Address, _to []common.Address) (*BalanceContractCompletePurchaseIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "CompletePurchase", _hashRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractCompletePurchaseIterator{contract: _BalanceContract.contract, event: "CompletePurchase", logs: logs, sub: sub}, nil
}

// WatchCompletePurchase is a free log subscription operation binding the contract event 0xbedfb3b4f53acc8876b63deec056b5a4252ea296380b4d55473f7d6f2f76889c.
//
// Solidity: event CompletePurchase(bytes32 indexed _hash, address indexed _from, address indexed _to, bytes32 _txHash)
func (_BalanceContract *BalanceContractFilterer) WatchCompletePurchase(opts *bind.WatchOpts, sink chan<- *BalanceContractCompletePurchase, _hash [][32]byte, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "CompletePurchase", _hashRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractCompletePurchase)
				if err := _BalanceContract.contract.UnpackLog(event, "CompletePurchase", log); err != nil {
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

// ParseCompletePurchase is a log parse operation binding the contract event 0xbedfb3b4f53acc8876b63deec056b5a4252ea296380b4d55473f7d6f2f76889c.
//
// Solidity: event CompletePurchase(bytes32 indexed _hash, address indexed _from, address indexed _to, bytes32 _txHash)
func (_BalanceContract *BalanceContractFilterer) ParseCompletePurchase(log types.Log) (*BalanceContractCompletePurchase, error) {
	event := new(BalanceContractCompletePurchase)
	if err := _BalanceContract.contract.UnpackLog(event, "CompletePurchase", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalanceContractPriceSetIterator is returned from FilterPriceSet and is used to iterate over the raw logs and unpacked data for PriceSet events raised by the BalanceContract contract.
type BalanceContractPriceSetIterator struct {
	Event *BalanceContractPriceSet // Event containing the contract specifics and raw log

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
func (it *BalanceContractPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractPriceSet)
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
		it.Event = new(BalanceContractPriceSet)
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
func (it *BalanceContractPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractPriceSet represents a PriceSet event raised by the BalanceContract contract.
type BalanceContractPriceSet struct {
	Hash  [32]byte
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceSet is a free log retrieval operation binding the contract event 0xee47534b2400bc7be3fbdc39f0283b8643fc472a30ffd7324e6161b4b2b91f24.
//
// Solidity: event PriceSet(bytes32 indexed _hash, uint256 _price)
func (_BalanceContract *BalanceContractFilterer) FilterPriceSet(opts *bind.FilterOpts, _hash [][32]byte) (*BalanceContractPriceSetIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "PriceSet", _hashRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractPriceSetIterator{contract: _BalanceContract.contract, event: "PriceSet", logs: logs, sub: sub}, nil
}

// WatchPriceSet is a free log subscription operation binding the contract event 0xee47534b2400bc7be3fbdc39f0283b8643fc472a30ffd7324e6161b4b2b91f24.
//
// Solidity: event PriceSet(bytes32 indexed _hash, uint256 _price)
func (_BalanceContract *BalanceContractFilterer) WatchPriceSet(opts *bind.WatchOpts, sink chan<- *BalanceContractPriceSet, _hash [][32]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "PriceSet", _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractPriceSet)
				if err := _BalanceContract.contract.UnpackLog(event, "PriceSet", log); err != nil {
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

// ParsePriceSet is a log parse operation binding the contract event 0xee47534b2400bc7be3fbdc39f0283b8643fc472a30ffd7324e6161b4b2b91f24.
//
// Solidity: event PriceSet(bytes32 indexed _hash, uint256 _price)
func (_BalanceContract *BalanceContractFilterer) ParsePriceSet(log types.Log) (*BalanceContractPriceSet, error) {
	event := new(BalanceContractPriceSet)
	if err := _BalanceContract.contract.UnpackLog(event, "PriceSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalanceContractPurchaseRevokedIterator is returned from FilterPurchaseRevoked and is used to iterate over the raw logs and unpacked data for PurchaseRevoked events raised by the BalanceContract contract.
type BalanceContractPurchaseRevokedIterator struct {
	Event *BalanceContractPurchaseRevoked // Event containing the contract specifics and raw log

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
func (it *BalanceContractPurchaseRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractPurchaseRevoked)
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
		it.Event = new(BalanceContractPurchaseRevoked)
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
func (it *BalanceContractPurchaseRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractPurchaseRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractPurchaseRevoked represents a PurchaseRevoked event raised by the BalanceContract contract.
type BalanceContractPurchaseRevoked struct {
	Hash [32]byte
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPurchaseRevoked is a free log retrieval operation binding the contract event 0xe4f046578a8ec0e49030155ac1ec29d863bb3e723e3a429811634562ef141cc0.
//
// Solidity: event PurchaseRevoked(bytes32 indexed _hash, address indexed _from, address indexed _to)
func (_BalanceContract *BalanceContractFilterer) FilterPurchaseRevoked(opts *bind.FilterOpts, _hash [][32]byte, _from []common.Address, _to []common.Address) (*BalanceContractPurchaseRevokedIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "PurchaseRevoked", _hashRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractPurchaseRevokedIterator{contract: _BalanceContract.contract, event: "PurchaseRevoked", logs: logs, sub: sub}, nil
}

// WatchPurchaseRevoked is a free log subscription operation binding the contract event 0xe4f046578a8ec0e49030155ac1ec29d863bb3e723e3a429811634562ef141cc0.
//
// Solidity: event PurchaseRevoked(bytes32 indexed _hash, address indexed _from, address indexed _to)
func (_BalanceContract *BalanceContractFilterer) WatchPurchaseRevoked(opts *bind.WatchOpts, sink chan<- *BalanceContractPurchaseRevoked, _hash [][32]byte, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "PurchaseRevoked", _hashRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractPurchaseRevoked)
				if err := _BalanceContract.contract.UnpackLog(event, "PurchaseRevoked", log); err != nil {
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

// ParsePurchaseRevoked is a log parse operation binding the contract event 0xe4f046578a8ec0e49030155ac1ec29d863bb3e723e3a429811634562ef141cc0.
//
// Solidity: event PurchaseRevoked(bytes32 indexed _hash, address indexed _from, address indexed _to)
func (_BalanceContract *BalanceContractFilterer) ParsePurchaseRevoked(log types.Log) (*BalanceContractPurchaseRevoked, error) {
	event := new(BalanceContractPurchaseRevoked)
	if err := _BalanceContract.contract.UnpackLog(event, "PurchaseRevoked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalanceContractRequestPurchaseIterator is returned from FilterRequestPurchase and is used to iterate over the raw logs and unpacked data for RequestPurchase events raised by the BalanceContract contract.
type BalanceContractRequestPurchaseIterator struct {
	Event *BalanceContractRequestPurchase // Event containing the contract specifics and raw log

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
func (it *BalanceContractRequestPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractRequestPurchase)
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
		it.Event = new(BalanceContractRequestPurchase)
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
func (it *BalanceContractRequestPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractRequestPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractRequestPurchase represents a RequestPurchase event raised by the BalanceContract contract.
type BalanceContractRequestPurchase struct {
	Hash [32]byte
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRequestPurchase is a free log retrieval operation binding the contract event 0xd457a19a91892104bd37ed9c2de0d3398da0d5f42bca745436a5c9a17ab70e41.
//
// Solidity: event RequestPurchase(bytes32 indexed _hash, address indexed _from, address indexed _to)
func (_BalanceContract *BalanceContractFilterer) FilterRequestPurchase(opts *bind.FilterOpts, _hash [][32]byte, _from []common.Address, _to []common.Address) (*BalanceContractRequestPurchaseIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "RequestPurchase", _hashRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractRequestPurchaseIterator{contract: _BalanceContract.contract, event: "RequestPurchase", logs: logs, sub: sub}, nil
}

// WatchRequestPurchase is a free log subscription operation binding the contract event 0xd457a19a91892104bd37ed9c2de0d3398da0d5f42bca745436a5c9a17ab70e41.
//
// Solidity: event RequestPurchase(bytes32 indexed _hash, address indexed _from, address indexed _to)
func (_BalanceContract *BalanceContractFilterer) WatchRequestPurchase(opts *bind.WatchOpts, sink chan<- *BalanceContractRequestPurchase, _hash [][32]byte, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "RequestPurchase", _hashRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractRequestPurchase)
				if err := _BalanceContract.contract.UnpackLog(event, "RequestPurchase", log); err != nil {
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

// ParseRequestPurchase is a log parse operation binding the contract event 0xd457a19a91892104bd37ed9c2de0d3398da0d5f42bca745436a5c9a17ab70e41.
//
// Solidity: event RequestPurchase(bytes32 indexed _hash, address indexed _from, address indexed _to)
func (_BalanceContract *BalanceContractFilterer) ParseRequestPurchase(log types.Log) (*BalanceContractRequestPurchase, error) {
	event := new(BalanceContractRequestPurchase)
	if err := _BalanceContract.contract.UnpackLog(event, "RequestPurchase", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalanceContractSubscribeIterator is returned from FilterSubscribe and is used to iterate over the raw logs and unpacked data for Subscribe events raised by the BalanceContract contract.
type BalanceContractSubscribeIterator struct {
	Event *BalanceContractSubscribe // Event containing the contract specifics and raw log

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
func (it *BalanceContractSubscribeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractSubscribe)
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
		it.Event = new(BalanceContractSubscribe)
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
func (it *BalanceContractSubscribeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractSubscribeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractSubscribe represents a Subscribe event raised by the BalanceContract contract.
type BalanceContractSubscribe struct {
	Topic common.Hash
	From  common.Address
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSubscribe is a free log retrieval operation binding the contract event 0xc36d692ea82425e7776ac80f320816b8cbd7b02289e0d83d5ef5c489af7355ca.
//
// Solidity: event Subscribe(string indexed _topic, address indexed _from, address indexed _to)
func (_BalanceContract *BalanceContractFilterer) FilterSubscribe(opts *bind.FilterOpts, _topic []string, _from []common.Address, _to []common.Address) (*BalanceContractSubscribeIterator, error) {

	var _topicRule []interface{}
	for _, _topicItem := range _topic {
		_topicRule = append(_topicRule, _topicItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "Subscribe", _topicRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractSubscribeIterator{contract: _BalanceContract.contract, event: "Subscribe", logs: logs, sub: sub}, nil
}

// WatchSubscribe is a free log subscription operation binding the contract event 0xc36d692ea82425e7776ac80f320816b8cbd7b02289e0d83d5ef5c489af7355ca.
//
// Solidity: event Subscribe(string indexed _topic, address indexed _from, address indexed _to)
func (_BalanceContract *BalanceContractFilterer) WatchSubscribe(opts *bind.WatchOpts, sink chan<- *BalanceContractSubscribe, _topic []string, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _topicRule []interface{}
	for _, _topicItem := range _topic {
		_topicRule = append(_topicRule, _topicItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "Subscribe", _topicRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractSubscribe)
				if err := _BalanceContract.contract.UnpackLog(event, "Subscribe", log); err != nil {
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

// ParseSubscribe is a log parse operation binding the contract event 0xc36d692ea82425e7776ac80f320816b8cbd7b02289e0d83d5ef5c489af7355ca.
//
// Solidity: event Subscribe(string indexed _topic, address indexed _from, address indexed _to)
func (_BalanceContract *BalanceContractFilterer) ParseSubscribe(log types.Log) (*BalanceContractSubscribe, error) {
	event := new(BalanceContractSubscribe)
	if err := _BalanceContract.contract.UnpackLog(event, "Subscribe", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalanceContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BalanceContract contract.
type BalanceContractTransferIterator struct {
	Event *BalanceContractTransfer // Event containing the contract specifics and raw log

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
func (it *BalanceContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractTransfer)
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
		it.Event = new(BalanceContractTransfer)
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
func (it *BalanceContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractTransfer represents a Transfer event raised by the BalanceContract contract.
type BalanceContractTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_BalanceContract *BalanceContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BalanceContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractTransferIterator{contract: _BalanceContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_BalanceContract *BalanceContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BalanceContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractTransfer)
				if err := _BalanceContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_BalanceContract *BalanceContractFilterer) ParseTransfer(log types.Log) (*BalanceContractTransfer, error) {
	event := new(BalanceContractTransfer)
	if err := _BalanceContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DataLedgerContractABI is the input ABI used to generate the binding from.
const DataLedgerContractABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getIoTAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"getIoTAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"topic\",\"type\":\"string\"}],\"name\":\"topicExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataLedgerContractFuncSigs maps the 4-byte function signature to its string representation.
var DataLedgerContractFuncSigs = map[string]string{
	"6ade0219": "getIoTAddress(bytes32)",
	"ab680e02": "getIoTAddress(string)",
	"78b5749d": "topicExists(string)",
}

// DataLedgerContractBin is the compiled bytecode used for deploying new contracts.
var DataLedgerContractBin = "0x608060405234801561001057600080fd5b5061018a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80636ade02191461004657806378b5749d1461007f578063ab680e0214610139575b600080fd5b6100636004803603602081101561005c57600080fd5b503561014f565b604080516001600160a01b039092168252519081900360200190f35b6101256004803603602081101561009557600080fd5b8101906020810181356401000000008111156100b057600080fd5b8201836020820111156100c257600080fd5b803590602001918460018302840111640100000000831117156100e457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061014f945050505050565b604080519115158252519081900360200190f35b6100636004803603602081101561009557600080fd5b5060009056fea265627a7a72315820757d2c4a914351b1179f354fcb40819faae008af6c11b176b6cabfd32f28fdeb64736f6c63430005100032"

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
