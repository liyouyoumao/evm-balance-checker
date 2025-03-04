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

// BalanceCheckerTokenBalance is an auto generated low-level Go binding around an user-defined struct.
type BalanceCheckerTokenBalance struct {
	UserToken BalanceCheckerUserToken
	Balance   *big.Int
	Decimals  uint8
	Invalid   bool
}

// BalanceCheckerUserToken is an auto generated low-level Go binding around an user-defined struct.
type BalanceCheckerUserToken struct {
	User  common.Address
	Token common.Address
}

// BalancecheckerMetaData contains all meta data concerning the Balancechecker contract.
var BalancecheckerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"IsERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBalanceChecker.UserToken[]\",\"name\":\"userTokens\",\"type\":\"tuple[]\"}],\"name\":\"balances\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBalanceChecker.UserToken\",\"name\":\"userToken\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"internalType\":\"structBalanceChecker.TokenBalance[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610e548061001d5f395ff3fe608060405260043610610037575f3560e01c80631049334f1461007b5780639bbb41c7146100b9578063ccb94eb8146100f557610077565b36610077576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161006e906107c3565b60405180910390fd5b5f80fd5b348015610086575f80fd5b506100a1600480360381019061009c919061084c565b610131565b6040516100b0939291906108d7565b60405180910390f35b3480156100c4575f80fd5b506100df60048036038101906100da919061090c565b610244565b6040516100ec9190610937565b60405180910390f35b348015610100575f80fd5b5061011b600480360381019061011691906109b1565b610461565b6040516101289190610b60565b60405180910390f35b5f805f61013d84610244565b61014f575f805f92509250925061023d565b8373ffffffffffffffffffffffffffffffffffffffff166370a08231866040518263ffffffff1660e01b81526004016101889190610b8f565b602060405180830381865afa1580156101a3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101c79190610bd2565b8473ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610210573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102349190610c27565b60019250925092505b9250925092565b5f808273ffffffffffffffffffffffffffffffffffffffff163b1161026b575f905061045c565b8173ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156102d357506040513d601f19601f820116820180604052508101906102d09190610bd2565b60015b6102df575f905061045c565b508173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016103199190610b8f565b602060405180830381865afa92505050801561035357506040513d601f19601f820116820180604052508101906103509190610bd2565b60015b61035f575f905061045c565b508173ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156103c857506040513d601f19601f820116820180604052508101906103c59190610c27565b60015b6103d4575f905061045c565b508173ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e305f6040518363ffffffff1660e01b8152600401610410929190610c52565b602060405180830381865afa92505050801561044a57506040513d601f19601f820116820180604052508101906104479190610bd2565b60015b610456575f905061045c565b50600190505b919050565b60605f8383905067ffffffffffffffff81111561048157610480610c79565b5b6040519080825280602002602001820160405280156104ba57816020015b6104a76106cf565b81526020019060019003908161049f5790505b5090505f5b848490508110156106c4575f73ffffffffffffffffffffffffffffffffffffffff168585838181106104f4576104f3610ca6565b5b905060400201602001602081019061050c919061090c565b73ffffffffffffffffffffffffffffffffffffffff1614610602575f805f6105878888868181106105405761053f610ca6565b5b9050604002015f016020810190610557919061090c565b89898781811061056a57610569610ca6565b5b9050604002016020016020810190610582919061090c565b610131565b92509250925060405180608001604052808989878181106105ab576105aa610ca6565b5b9050604002018036038101906105c19190610d7f565b81526020018481526020018360ff168152602001821515158152508585815181106105ef576105ee610ca6565b5b60200260200101819052505050506106b1565b60405180608001604052808686848181106106205761061f610ca6565b5b9050604002018036038101906106369190610d7f565b815260200186868481811061064e5761064d610ca6565b5b9050604002015f016020810190610665919061090c565b73ffffffffffffffffffffffffffffffffffffffff16318152602001601260ff168152602001600115158152508282815181106106a5576106a4610ca6565b5b60200260200101819052505b80806106bc90610dd7565b9150506104bf565b508091505092915050565b60405180608001604052806106e26106ff565b81526020015f81526020015f60ff1681526020015f151581525090565b60405180604001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b5f82825260208201905092915050565b7f42616c616e6365436865636b657220646f6573206e6f742061636365707420705f8201527f61796d656e747300000000000000000000000000000000000000000000000000602082015250565b5f6107ad602783610743565b91506107b882610753565b604082019050919050565b5f6020820190508181035f8301526107da816107a1565b9050919050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61081b826107f2565b9050919050565b61082b81610811565b8114610835575f80fd5b50565b5f8135905061084681610822565b92915050565b5f8060408385031215610862576108616107ea565b5b5f61086f85828601610838565b925050602061088085828601610838565b9150509250929050565b5f819050919050565b61089c8161088a565b82525050565b5f60ff82169050919050565b6108b7816108a2565b82525050565b5f8115159050919050565b6108d1816108bd565b82525050565b5f6060820190506108ea5f830186610893565b6108f760208301856108ae565b61090460408301846108c8565b949350505050565b5f60208284031215610921576109206107ea565b5b5f61092e84828501610838565b91505092915050565b5f60208201905061094a5f8301846108c8565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261097157610970610950565b5b8235905067ffffffffffffffff81111561098e5761098d610954565b5b6020830191508360408202830111156109aa576109a9610958565b5b9250929050565b5f80602083850312156109c7576109c66107ea565b5b5f83013567ffffffffffffffff8111156109e4576109e36107ee565b5b6109f08582860161095c565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610a2e81610811565b82525050565b604082015f820151610a485f850182610a25565b506020820151610a5b6020850182610a25565b50505050565b610a6a8161088a565b82525050565b610a79816108a2565b82525050565b610a88816108bd565b82525050565b60a082015f820151610aa25f850182610a34565b506020820151610ab56040850182610a61565b506040820151610ac86060850182610a70565b506060820151610adb6080850182610a7f565b50505050565b5f610aec8383610a8e565b60a08301905092915050565b5f602082019050919050565b5f610b0e826109fc565b610b188185610a06565b9350610b2383610a16565b805f5b83811015610b53578151610b3a8882610ae1565b9750610b4583610af8565b925050600181019050610b26565b5085935050505092915050565b5f6020820190508181035f830152610b788184610b04565b905092915050565b610b8981610811565b82525050565b5f602082019050610ba25f830184610b80565b92915050565b610bb18161088a565b8114610bbb575f80fd5b50565b5f81519050610bcc81610ba8565b92915050565b5f60208284031215610be757610be66107ea565b5b5f610bf484828501610bbe565b91505092915050565b610c06816108a2565b8114610c10575f80fd5b50565b5f81519050610c2181610bfd565b92915050565b5f60208284031215610c3c57610c3b6107ea565b5b5f610c4984828501610c13565b91505092915050565b5f604082019050610c655f830185610b80565b610c726020830184610b80565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f80fd5b5f601f19601f8301169050919050565b610cf082610cd7565b810181811067ffffffffffffffff82111715610d0f57610d0e610c79565b5b80604052505050565b5f610d216107e1565b9050610d2d8282610ce7565b919050565b5f60408284031215610d4757610d46610cd3565b5b610d516040610d18565b90505f610d6084828501610838565b5f830152506020610d7384828501610838565b60208301525092915050565b5f60408284031215610d9457610d936107ea565b5b5f610da184828501610d32565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610de18261088a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e1357610e12610daa565b5b60018201905091905056fea264697066735822122079db8108f31d6459c5e792c984e6d4bd9cbf1396d0eec53d91f806b662a5e9a164736f6c63430008150033",
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

// IsERC20 is a free data retrieval call binding the contract method 0x9bbb41c7.
//
// Solidity: function IsERC20(address token) view returns(bool)
func (_Balancechecker *BalancecheckerCaller) IsERC20(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Balancechecker.contract.Call(opts, &out, "IsERC20", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsERC20 is a free data retrieval call binding the contract method 0x9bbb41c7.
//
// Solidity: function IsERC20(address token) view returns(bool)
func (_Balancechecker *BalancecheckerSession) IsERC20(token common.Address) (bool, error) {
	return _Balancechecker.Contract.IsERC20(&_Balancechecker.CallOpts, token)
}

// IsERC20 is a free data retrieval call binding the contract method 0x9bbb41c7.
//
// Solidity: function IsERC20(address token) view returns(bool)
func (_Balancechecker *BalancecheckerCallerSession) IsERC20(token common.Address) (bool, error) {
	return _Balancechecker.Contract.IsERC20(&_Balancechecker.CallOpts, token)
}

// Balances is a free data retrieval call binding the contract method 0xccb94eb8.
//
// Solidity: function balances((address,address)[] userTokens) view returns(((address,address),uint256,uint8,bool)[])
func (_Balancechecker *BalancecheckerCaller) Balances(opts *bind.CallOpts, userTokens []BalanceCheckerUserToken) ([]BalanceCheckerTokenBalance, error) {
	var out []interface{}
	err := _Balancechecker.contract.Call(opts, &out, "balances", userTokens)

	if err != nil {
		return *new([]BalanceCheckerTokenBalance), err
	}

	out0 := *abi.ConvertType(out[0], new([]BalanceCheckerTokenBalance)).(*[]BalanceCheckerTokenBalance)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0xccb94eb8.
//
// Solidity: function balances((address,address)[] userTokens) view returns(((address,address),uint256,uint8,bool)[])
func (_Balancechecker *BalancecheckerSession) Balances(userTokens []BalanceCheckerUserToken) ([]BalanceCheckerTokenBalance, error) {
	return _Balancechecker.Contract.Balances(&_Balancechecker.CallOpts, userTokens)
}

// Balances is a free data retrieval call binding the contract method 0xccb94eb8.
//
// Solidity: function balances((address,address)[] userTokens) view returns(((address,address),uint256,uint8,bool)[])
func (_Balancechecker *BalancecheckerCallerSession) Balances(userTokens []BalanceCheckerUserToken) ([]BalanceCheckerTokenBalance, error) {
	return _Balancechecker.Contract.Balances(&_Balancechecker.CallOpts, userTokens)
}

// TokenBalance is a free data retrieval call binding the contract method 0x1049334f.
//
// Solidity: function tokenBalance(address user, address token) view returns(uint256, uint8, bool)
func (_Balancechecker *BalancecheckerCaller) TokenBalance(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, uint8, bool, error) {
	var out []interface{}
	err := _Balancechecker.contract.Call(opts, &out, "tokenBalance", user, token)

	if err != nil {
		return *new(*big.Int), *new(uint8), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// TokenBalance is a free data retrieval call binding the contract method 0x1049334f.
//
// Solidity: function tokenBalance(address user, address token) view returns(uint256, uint8, bool)
func (_Balancechecker *BalancecheckerSession) TokenBalance(user common.Address, token common.Address) (*big.Int, uint8, bool, error) {
	return _Balancechecker.Contract.TokenBalance(&_Balancechecker.CallOpts, user, token)
}

// TokenBalance is a free data retrieval call binding the contract method 0x1049334f.
//
// Solidity: function tokenBalance(address user, address token) view returns(uint256, uint8, bool)
func (_Balancechecker *BalancecheckerCallerSession) TokenBalance(user common.Address, token common.Address) (*big.Int, uint8, bool, error) {
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
