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
	Bin: "0x608060405234801561000f575f80fd5b50610e538061001d5f395ff3fe608060405260043610610037575f3560e01c80631049334f1461007b5780639bbb41c7146100b9578063ccb94eb8146100f557610077565b36610077576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161006e906107c2565b60405180910390fd5b5f80fd5b348015610086575f80fd5b506100a1600480360381019061009c919061084b565b610131565b6040516100b0939291906108d6565b60405180910390f35b3480156100c4575f80fd5b506100df60048036038101906100da919061090b565b610244565b6040516100ec9190610936565b60405180910390f35b348015610100575f80fd5b5061011b600480360381019061011691906109b0565b610461565b6040516101289190610b5f565b60405180910390f35b5f805f61013d84610244565b61014f575f805f92509250925061023d565b8373ffffffffffffffffffffffffffffffffffffffff166370a08231866040518263ffffffff1660e01b81526004016101889190610b8e565b602060405180830381865afa1580156101a3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101c79190610bd1565b8473ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610210573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102349190610c26565b60019250925092505b9250925092565b5f808273ffffffffffffffffffffffffffffffffffffffff163b1161026b575f905061045c565b8173ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156102d357506040513d601f19601f820116820180604052508101906102d09190610bd1565b60015b6102df575f905061045c565b508173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016103199190610b8e565b602060405180830381865afa92505050801561035357506040513d601f19601f820116820180604052508101906103509190610bd1565b60015b61035f575f905061045c565b508173ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156103c857506040513d601f19601f820116820180604052508101906103c59190610c26565b60015b6103d4575f905061045c565b508173ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e305f6040518363ffffffff1660e01b8152600401610410929190610c51565b602060405180830381865afa92505050801561044a57506040513d601f19601f820116820180604052508101906104479190610bd1565b60015b610456575f905061045c565b50600190505b919050565b60605f8383905067ffffffffffffffff81111561048157610480610c78565b5b6040519080825280602002602001820160405280156104ba57816020015b6104a76106ce565b81526020019060019003908161049f5790505b5090505f5b848490508110156106c3575f73ffffffffffffffffffffffffffffffffffffffff168585838181106104f4576104f3610ca5565b5b905060400201602001602081019061050c919061090b565b73ffffffffffffffffffffffffffffffffffffffff1614610601575f805f6105878888868181106105405761053f610ca5565b5b9050604002015f016020810190610557919061090b565b89898781811061056a57610569610ca5565b5b9050604002016020016020810190610582919061090b565b610131565b92509250925060405180608001604052808989878181106105ab576105aa610ca5565b5b9050604002018036038101906105c19190610d7e565b81526020018481526020018360ff1681526020018215158152508585815181106105ee576105ed610ca5565b5b60200260200101819052505050506106b0565b604051806080016040528086868481811061061f5761061e610ca5565b5b9050604002018036038101906106359190610d7e565b815260200186868481811061064d5761064c610ca5565b5b9050604002015f016020810190610664919061090b565b73ffffffffffffffffffffffffffffffffffffffff16318152602001601260ff168152602001600115158152508282815181106106a4576106a3610ca5565b5b60200260200101819052505b80806106bb90610dd6565b9150506104bf565b508091505092915050565b60405180608001604052806106e16106fe565b81526020015f81526020015f60ff1681526020015f151581525090565b60405180604001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b5f82825260208201905092915050565b7f42616c616e6365436865636b657220646f6573206e6f742061636365707420705f8201527f61796d656e747300000000000000000000000000000000000000000000000000602082015250565b5f6107ac602783610742565b91506107b782610752565b604082019050919050565b5f6020820190508181035f8301526107d9816107a0565b9050919050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61081a826107f1565b9050919050565b61082a81610810565b8114610834575f80fd5b50565b5f8135905061084581610821565b92915050565b5f8060408385031215610861576108606107e9565b5b5f61086e85828601610837565b925050602061087f85828601610837565b9150509250929050565b5f819050919050565b61089b81610889565b82525050565b5f60ff82169050919050565b6108b6816108a1565b82525050565b5f8115159050919050565b6108d0816108bc565b82525050565b5f6060820190506108e95f830186610892565b6108f660208301856108ad565b61090360408301846108c7565b949350505050565b5f602082840312156109205761091f6107e9565b5b5f61092d84828501610837565b91505092915050565b5f6020820190506109495f8301846108c7565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f8401126109705761096f61094f565b5b8235905067ffffffffffffffff81111561098d5761098c610953565b5b6020830191508360408202830111156109a9576109a8610957565b5b9250929050565b5f80602083850312156109c6576109c56107e9565b5b5f83013567ffffffffffffffff8111156109e3576109e26107ed565b5b6109ef8582860161095b565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610a2d81610810565b82525050565b604082015f820151610a475f850182610a24565b506020820151610a5a6020850182610a24565b50505050565b610a6981610889565b82525050565b610a78816108a1565b82525050565b610a87816108bc565b82525050565b60a082015f820151610aa15f850182610a33565b506020820151610ab46040850182610a60565b506040820151610ac76060850182610a6f565b506060820151610ada6080850182610a7e565b50505050565b5f610aeb8383610a8d565b60a08301905092915050565b5f602082019050919050565b5f610b0d826109fb565b610b178185610a05565b9350610b2283610a15565b805f5b83811015610b52578151610b398882610ae0565b9750610b4483610af7565b925050600181019050610b25565b5085935050505092915050565b5f6020820190508181035f830152610b778184610b03565b905092915050565b610b8881610810565b82525050565b5f602082019050610ba15f830184610b7f565b92915050565b610bb081610889565b8114610bba575f80fd5b50565b5f81519050610bcb81610ba7565b92915050565b5f60208284031215610be657610be56107e9565b5b5f610bf384828501610bbd565b91505092915050565b610c05816108a1565b8114610c0f575f80fd5b50565b5f81519050610c2081610bfc565b92915050565b5f60208284031215610c3b57610c3a6107e9565b5b5f610c4884828501610c12565b91505092915050565b5f604082019050610c645f830185610b7f565b610c716020830184610b7f565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f80fd5b5f601f19601f8301169050919050565b610cef82610cd6565b810181811067ffffffffffffffff82111715610d0e57610d0d610c78565b5b80604052505050565b5f610d206107e0565b9050610d2c8282610ce6565b919050565b5f60408284031215610d4657610d45610cd2565b5b610d506040610d17565b90505f610d5f84828501610837565b5f830152506020610d7284828501610837565b60208301525092915050565b5f60408284031215610d9357610d926107e9565b5b5f610da084828501610d31565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610de082610889565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e1257610e11610da9565b5b60018201905091905056fea2646970667358221220472889b3ee64cec0a0840ccf9b9075450bfba2e3fcd718b44eb94c53a12fa31b64736f6c63430008150033",
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
