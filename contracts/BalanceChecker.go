// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancechecker

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BalanceCheckerTokenBalanceInfo is an auto generated low-level Go binding around an user-defined struct.
type BalanceCheckerTokenBalanceInfo struct {
	Token    common.Address
	Balance  *big.Int
	Decimals uint8
}

// BalancecheckerMetaData contains all meta data concerning the Balancechecker contract.
var BalancecheckerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"balances\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structBalanceChecker.TokenBalanceInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610a0d8061001d5f395ff3fe60806040526004361061002c575f3560e01c80631049334f14610070578063d3e5ca87146100ad5761006c565b3661006c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161006390610501565b60405180910390fd5b5f80fd5b34801561007b575f80fd5b5061009660048036038101906100919190610581565b6100e9565b6040516100a49291906105f2565b60405180910390f35b3480156100b8575f80fd5b506100d360048036038101906100ce919061067a565b61021f565b6040516100e091906107ec565b60405180910390f35b5f805f8373ffffffffffffffffffffffffffffffffffffffff163b1115610211578273ffffffffffffffffffffffffffffffffffffffff166370a08231856040518263ffffffff1660e01b8152600401610143919061081b565b602060405180830381865afa92505050801561017d57506040513d601f19601f8201168201806040525081019061017a919061085e565b60015b61018c575f8091509150610218565b8373ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156101f457506040513d601f19601f820116820180604052508101906101f191906108b3565b60015b610204575f809250925050610218565b8181935093505050610218565b5f80915091505b9250929050565b60605f8383905067ffffffffffffffff81111561023f5761023e6108de565b5b60405190808252806020026020018201604052801561027857816020015b61026561044a565b81526020019060019003908161025d5790505b5090505f5b8484905081101561043e575f73ffffffffffffffffffffffffffffffffffffffff168585838181106102b2576102b161090b565b5b90506020020160208101906102c79190610938565b73ffffffffffffffffffffffffffffffffffffffff1614610399575f80610315888888868181106102fb576102fa61090b565b5b90506020020160208101906103109190610938565b6100e9565b9150915060405180606001604052808888868181106103375761033661090b565b5b905060200201602081019061034c9190610938565b73ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018260ff168152508484815181106103875761038661090b565b5b6020026020010181905250505061042b565b60405180606001604052808686848181106103b7576103b661090b565b5b90506020020160208101906103cc9190610938565b73ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff16318152602001601260ff1681525082828151811061041f5761041e61090b565b5b60200260200101819052505b808061043690610990565b91505061027d565b50809150509392505050565b60405180606001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f60ff1681525090565b5f82825260208201905092915050565b7f42616c616e6365436865636b657220646f6573206e6f742061636365707420705f8201527f61796d656e747300000000000000000000000000000000000000000000000000602082015250565b5f6104eb602783610481565b91506104f682610491565b604082019050919050565b5f6020820190508181035f830152610518816104df565b9050919050565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61055082610527565b9050919050565b61056081610546565b811461056a575f80fd5b50565b5f8135905061057b81610557565b92915050565b5f80604083850312156105975761059661051f565b5b5f6105a48582860161056d565b92505060206105b58582860161056d565b9150509250929050565b5f819050919050565b6105d1816105bf565b82525050565b5f60ff82169050919050565b6105ec816105d7565b82525050565b5f6040820190506106055f8301856105c8565b61061260208301846105e3565b9392505050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261063a57610639610619565b5b8235905067ffffffffffffffff8111156106575761065661061d565b5b60208301915083602082028301111561067357610672610621565b5b9250929050565b5f805f604084860312156106915761069061051f565b5b5f61069e8682870161056d565b935050602084013567ffffffffffffffff8111156106bf576106be610523565b5b6106cb86828701610625565b92509250509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61070981610546565b82525050565b610718816105bf565b82525050565b610727816105d7565b82525050565b606082015f8201516107415f850182610700565b506020820151610754602085018261070f565b506040820151610767604085018261071e565b50505050565b5f610778838361072d565b60608301905092915050565b5f602082019050919050565b5f61079a826106d7565b6107a481856106e1565b93506107af836106f1565b805f5b838110156107df5781516107c6888261076d565b97506107d183610784565b9250506001810190506107b2565b5085935050505092915050565b5f6020820190508181035f8301526108048184610790565b905092915050565b61081581610546565b82525050565b5f60208201905061082e5f83018461080c565b92915050565b61083d816105bf565b8114610847575f80fd5b50565b5f8151905061085881610834565b92915050565b5f602082840312156108735761087261051f565b5b5f6108808482850161084a565b91505092915050565b610892816105d7565b811461089c575f80fd5b50565b5f815190506108ad81610889565b92915050565b5f602082840312156108c8576108c761051f565b5b5f6108d58482850161089f565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6020828403121561094d5761094c61051f565b5b5f61095a8482850161056d565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61099a826105bf565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036109cc576109cb610963565b5b60018201905091905056fea26469706673582212209e1eeae88f56a7501eb5a226022593a414cbc25ff2a095e701ba6dc1824ad80764736f6c63430008150033",
}

// BalancecheckerABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancecheckerMetaData.ABI instead.
var BalancecheckerABI = BalancecheckerMetaData.ABI

// BalancecheckerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BalancecheckerMetaData.Bin instead.
var BalancecheckerBin = BalancecheckerMetaData.Bin

// DeployBalancechecker deploys a new Ethereum contract, binding an instance of Balancechecker to it.
func DeployBalancechecker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Balancechecker, error) {
	parsed, err := BalancecheckerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BalancecheckerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Balancechecker{BalancecheckerCaller: BalancecheckerCaller{contract: contract}, BalancecheckerTransactor: BalancecheckerTransactor{contract: contract}, BalancecheckerFilterer: BalancecheckerFilterer{contract: contract}}, nil
}

// Balancechecker is an auto generated Go binding around an Ethereum contract.
type Balancechecker struct {
	BalancecheckerCaller     // Read-only binding to the contract
	BalancecheckerTransactor // Write-only binding to the contract
	BalancecheckerFilterer   // Log filterer for contract events
}

// BalancecheckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancecheckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancecheckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancecheckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancecheckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancecheckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancecheckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancecheckerSession struct {
	Contract     *Balancechecker   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancecheckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancecheckerCallerSession struct {
	Contract *BalancecheckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BalancecheckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancecheckerTransactorSession struct {
	Contract     *BalancecheckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BalancecheckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancecheckerRaw struct {
	Contract *Balancechecker // Generic contract binding to access the raw methods on
}

// BalancecheckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancecheckerCallerRaw struct {
	Contract *BalancecheckerCaller // Generic read-only contract binding to access the raw methods on
}

// BalancecheckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancecheckerTransactorRaw struct {
	Contract *BalancecheckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancechecker creates a new instance of Balancechecker, bound to a specific deployed contract.
func NewBalancechecker(address common.Address, backend bind.ContractBackend) (*Balancechecker, error) {
	contract, err := bindBalancechecker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balancechecker{BalancecheckerCaller: BalancecheckerCaller{contract: contract}, BalancecheckerTransactor: BalancecheckerTransactor{contract: contract}, BalancecheckerFilterer: BalancecheckerFilterer{contract: contract}}, nil
}

// NewBalancecheckerCaller creates a new read-only instance of Balancechecker, bound to a specific deployed contract.
func NewBalancecheckerCaller(address common.Address, caller bind.ContractCaller) (*BalancecheckerCaller, error) {
	contract, err := bindBalancechecker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancecheckerCaller{contract: contract}, nil
}

// NewBalancecheckerTransactor creates a new write-only instance of Balancechecker, bound to a specific deployed contract.
func NewBalancecheckerTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancecheckerTransactor, error) {
	contract, err := bindBalancechecker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancecheckerTransactor{contract: contract}, nil
}

// NewBalancecheckerFilterer creates a new log filterer instance of Balancechecker, bound to a specific deployed contract.
func NewBalancecheckerFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancecheckerFilterer, error) {
	contract, err := bindBalancechecker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancecheckerFilterer{contract: contract}, nil
}

// bindBalancechecker binds a generic wrapper to an already deployed contract.
func bindBalancechecker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancecheckerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancechecker *BalancecheckerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancechecker.Contract.BalancecheckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancechecker *BalancecheckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancechecker.Contract.BalancecheckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancechecker *BalancecheckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancechecker.Contract.BalancecheckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancechecker *BalancecheckerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancechecker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancechecker *BalancecheckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancechecker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancechecker *BalancecheckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancechecker.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0xd3e5ca87.
//
// Solidity: function balances(address user, address[] tokens) view returns((address,uint256,uint8)[])
func (_Balancechecker *BalancecheckerCaller) Balances(opts *bind.CallOpts, user common.Address, tokens []common.Address) ([]BalanceCheckerTokenBalanceInfo, error) {
	var out []interface{}
	err := _Balancechecker.contract.Call(opts, &out, "balances", user, tokens)

	if err != nil {
		return *new([]BalanceCheckerTokenBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]BalanceCheckerTokenBalanceInfo)).(*[]BalanceCheckerTokenBalanceInfo)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0xd3e5ca87.
//
// Solidity: function balances(address user, address[] tokens) view returns((address,uint256,uint8)[])
func (_Balancechecker *BalancecheckerSession) Balances(user common.Address, tokens []common.Address) ([]BalanceCheckerTokenBalanceInfo, error) {
	return _Balancechecker.Contract.Balances(&_Balancechecker.CallOpts, user, tokens)
}

// Balances is a free data retrieval call binding the contract method 0xd3e5ca87.
//
// Solidity: function balances(address user, address[] tokens) view returns((address,uint256,uint8)[])
func (_Balancechecker *BalancecheckerCallerSession) Balances(user common.Address, tokens []common.Address) ([]BalanceCheckerTokenBalanceInfo, error) {
	return _Balancechecker.Contract.Balances(&_Balancechecker.CallOpts, user, tokens)
}

// TokenBalance is a free data retrieval call binding the contract method 0x1049334f.
//
// Solidity: function tokenBalance(address user, address token) view returns(uint256, uint8)
func (_Balancechecker *BalancecheckerCaller) TokenBalance(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, uint8, error) {
	var out []interface{}
	err := _Balancechecker.contract.Call(opts, &out, "tokenBalance", user, token)

	if err != nil {
		return *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// TokenBalance is a free data retrieval call binding the contract method 0x1049334f.
//
// Solidity: function tokenBalance(address user, address token) view returns(uint256, uint8)
func (_Balancechecker *BalancecheckerSession) TokenBalance(user common.Address, token common.Address) (*big.Int, uint8, error) {
	return _Balancechecker.Contract.TokenBalance(&_Balancechecker.CallOpts, user, token)
}

// TokenBalance is a free data retrieval call binding the contract method 0x1049334f.
//
// Solidity: function tokenBalance(address user, address token) view returns(uint256, uint8)
func (_Balancechecker *BalancecheckerCallerSession) TokenBalance(user common.Address, token common.Address) (*big.Int, uint8, error) {
	return _Balancechecker.Contract.TokenBalance(&_Balancechecker.CallOpts, user, token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Balancechecker *BalancecheckerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancechecker.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Balancechecker *BalancecheckerSession) Receive() (*types.Transaction, error) {
	return _Balancechecker.Contract.Receive(&_Balancechecker.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Balancechecker *BalancecheckerTransactorSession) Receive() (*types.Transaction, error) {
	return _Balancechecker.Contract.Receive(&_Balancechecker.TransactOpts)
}
