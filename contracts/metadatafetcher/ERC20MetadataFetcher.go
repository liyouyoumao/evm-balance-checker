// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package metadatafetcher

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

// ERC20MetadataFetcherERC20Metadata is an auto generated low-level Go binding around an user-defined struct.
type ERC20MetadataFetcherERC20Metadata struct {
	Symbol      string
	Decimals    uint8
	TotalSupply *big.Int
	Valid       bool
}

// MetadatafetcherMetaData contains all meta data concerning the Metadatafetcher contract.
var MetadatafetcherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"internalType\":\"structERC20MetadataFetcher.ERC20Metadata\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506106ac8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c8063c00f14ab1461002d575b5f80fd5b610047600480360381019061004291906102e5565b61005d565b6040516100549190610447565b60405180910390f35b610065610250565b61006d610250565b5f8160600190151590811515815250505f8373ffffffffffffffffffffffffffffffffffffffff163b116100a4578091505061024b565b5f8390508073ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa92505050801561011357506040513d5f823e3d601f19601f820116820180604052508101906101109190610585565b60015b61012157819250505061024b565b80835f0181905250508073ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561019257506040513d601f19601f8201168201806040525081019061018f91906105f6565b60015b6101a057819250505061024b565b80836020019060ff16908160ff1681525050508073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561021b57506040513d601f19601f82011682018060405250810190610218919061064b565b60015b61022957819250505061024b565b8083604001818152505050600182606001901515908115158152505081925050505b919050565b6040518060800160405280606081526020015f60ff1681526020015f81526020015f151581525090565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102b48261028b565b9050919050565b6102c4816102aa565b81146102ce575f80fd5b50565b5f813590506102df816102bb565b92915050565b5f602082840312156102fa576102f9610283565b5b5f610307848285016102d1565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561034757808201518184015260208101905061032c565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61036c82610310565b610376818561031a565b935061038681856020860161032a565b61038f81610352565b840191505092915050565b5f60ff82169050919050565b6103af8161039a565b82525050565b5f819050919050565b6103c7816103b5565b82525050565b5f8115159050919050565b6103e1816103cd565b82525050565b5f608083015f8301518482035f8601526104018282610362565b915050602083015161041660208601826103a6565b50604083015161042960408601826103be565b50606083015161043c60608601826103d8565b508091505092915050565b5f6020820190508181035f83015261045f81846103e7565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6104a582610352565b810181811067ffffffffffffffff821117156104c4576104c361046f565b5b80604052505050565b5f6104d661027a565b90506104e2828261049c565b919050565b5f67ffffffffffffffff8211156105015761050061046f565b5b61050a82610352565b9050602081019050919050565b5f610529610524846104e7565b6104cd565b9050828152602081018484840111156105455761054461046b565b5b61055084828561032a565b509392505050565b5f82601f83011261056c5761056b610467565b5b815161057c848260208601610517565b91505092915050565b5f6020828403121561059a57610599610283565b5b5f82015167ffffffffffffffff8111156105b7576105b6610287565b5b6105c384828501610558565b91505092915050565b6105d58161039a565b81146105df575f80fd5b50565b5f815190506105f0816105cc565b92915050565b5f6020828403121561060b5761060a610283565b5b5f610618848285016105e2565b91505092915050565b61062a816103b5565b8114610634575f80fd5b50565b5f8151905061064581610621565b92915050565b5f602082840312156106605761065f610283565b5b5f61066d84828501610637565b9150509291505056fea264697066735822122062ebda2a782d3689139ae71e37beffa46435f22a0851283dc248520cc44e9b2a64736f6c63430008150033",
}

// MetadatafetcherABI is the input ABI used to generate the binding from.
// Deprecated: Use MetadatafetcherMetaData.ABI instead.
var MetadatafetcherABI = MetadatafetcherMetaData.ABI

// MetadatafetcherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MetadatafetcherMetaData.Bin instead.
var MetadatafetcherBin = MetadatafetcherMetaData.Bin

// DeployMetadatafetcher deploys a new Ethereum contract, binding an instance of Metadatafetcher to it.
func DeployMetadatafetcher(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Metadatafetcher, error) {
	parsed, err := MetadatafetcherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MetadatafetcherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Metadatafetcher{MetadatafetcherCaller: MetadatafetcherCaller{contract: contract}, MetadatafetcherTransactor: MetadatafetcherTransactor{contract: contract}, MetadatafetcherFilterer: MetadatafetcherFilterer{contract: contract}}, nil
}

// Metadatafetcher is an auto generated Go binding around an Ethereum contract.
type Metadatafetcher struct {
	MetadatafetcherCaller     // Read-only binding to the contract
	MetadatafetcherTransactor // Write-only binding to the contract
	MetadatafetcherFilterer   // Log filterer for contract events
}

// MetadatafetcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetadatafetcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetadatafetcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetadatafetcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetadatafetcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetadatafetcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetadatafetcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetadatafetcherSession struct {
	Contract     *Metadatafetcher  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetadatafetcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetadatafetcherCallerSession struct {
	Contract *MetadatafetcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MetadatafetcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetadatafetcherTransactorSession struct {
	Contract     *MetadatafetcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MetadatafetcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetadatafetcherRaw struct {
	Contract *Metadatafetcher // Generic contract binding to access the raw methods on
}

// MetadatafetcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetadatafetcherCallerRaw struct {
	Contract *MetadatafetcherCaller // Generic read-only contract binding to access the raw methods on
}

// MetadatafetcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetadatafetcherTransactorRaw struct {
	Contract *MetadatafetcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetadatafetcher creates a new instance of Metadatafetcher, bound to a specific deployed contract.
func NewMetadatafetcher(address common.Address, backend bind.ContractBackend) (*Metadatafetcher, error) {
	contract, err := bindMetadatafetcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Metadatafetcher{MetadatafetcherCaller: MetadatafetcherCaller{contract: contract}, MetadatafetcherTransactor: MetadatafetcherTransactor{contract: contract}, MetadatafetcherFilterer: MetadatafetcherFilterer{contract: contract}}, nil
}

// NewMetadatafetcherCaller creates a new read-only instance of Metadatafetcher, bound to a specific deployed contract.
func NewMetadatafetcherCaller(address common.Address, caller bind.ContractCaller) (*MetadatafetcherCaller, error) {
	contract, err := bindMetadatafetcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetadatafetcherCaller{contract: contract}, nil
}

// NewMetadatafetcherTransactor creates a new write-only instance of Metadatafetcher, bound to a specific deployed contract.
func NewMetadatafetcherTransactor(address common.Address, transactor bind.ContractTransactor) (*MetadatafetcherTransactor, error) {
	contract, err := bindMetadatafetcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetadatafetcherTransactor{contract: contract}, nil
}

// NewMetadatafetcherFilterer creates a new log filterer instance of Metadatafetcher, bound to a specific deployed contract.
func NewMetadatafetcherFilterer(address common.Address, filterer bind.ContractFilterer) (*MetadatafetcherFilterer, error) {
	contract, err := bindMetadatafetcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetadatafetcherFilterer{contract: contract}, nil
}

// bindMetadatafetcher binds a generic wrapper to an already deployed contract.
func bindMetadatafetcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetadatafetcherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Metadatafetcher *MetadatafetcherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Metadatafetcher.Contract.MetadatafetcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Metadatafetcher *MetadatafetcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Metadatafetcher.Contract.MetadatafetcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Metadatafetcher *MetadatafetcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Metadatafetcher.Contract.MetadatafetcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Metadatafetcher *MetadatafetcherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Metadatafetcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Metadatafetcher *MetadatafetcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Metadatafetcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Metadatafetcher *MetadatafetcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Metadatafetcher.Contract.contract.Transact(opts, method, params...)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address tokenAddress) view returns((string,uint8,uint256,bool))
func (_Metadatafetcher *MetadatafetcherCaller) GetTokenMetadata(opts *bind.CallOpts, tokenAddress common.Address) (ERC20MetadataFetcherERC20Metadata, error) {
	var out []interface{}
	err := _Metadatafetcher.contract.Call(opts, &out, "getTokenMetadata", tokenAddress)

	if err != nil {
		return *new(ERC20MetadataFetcherERC20Metadata), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC20MetadataFetcherERC20Metadata)).(*ERC20MetadataFetcherERC20Metadata)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address tokenAddress) view returns((string,uint8,uint256,bool))
func (_Metadatafetcher *MetadatafetcherSession) GetTokenMetadata(tokenAddress common.Address) (ERC20MetadataFetcherERC20Metadata, error) {
	return _Metadatafetcher.Contract.GetTokenMetadata(&_Metadatafetcher.CallOpts, tokenAddress)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address tokenAddress) view returns((string,uint8,uint256,bool))
func (_Metadatafetcher *MetadatafetcherCallerSession) GetTokenMetadata(tokenAddress common.Address) (ERC20MetadataFetcherERC20Metadata, error) {
	return _Metadatafetcher.Contract.GetTokenMetadata(&_Metadatafetcher.CallOpts, tokenAddress)
}
