// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

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
	_ = abi.ConvertType
)

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"data\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506107098061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80630178fe3f1461004357806318562dae14610073578063f0ba84401461008f575b5f5ffd5b61005d60048036038101906100589190610259565b6100bf565b60405161006a91906102f4565b60405180910390f35b61008d60048036038101906100889190610375565b61015f565b005b6100a960048036038101906100a49190610259565b610184565b6040516100b691906102f4565b60405180910390f35b60605f5f8381526020019081526020015f2080546100dc906103ff565b80601f0160208091040260200160405190810160405280929190818152602001828054610108906103ff565b80156101535780601f1061012a57610100808354040283529160200191610153565b820191905f5260205f20905b81548152906001019060200180831161013657829003601f168201915b50505050509050919050565b81815f5f8681526020019081526020015f20918261017e929190610606565b50505050565b5f602052805f5260405f205f91509050805461019f906103ff565b80601f01602080910402602001604051908101604052809291908181526020018280546101cb906103ff565b80156102165780601f106101ed57610100808354040283529160200191610216565b820191905f5260205f20905b8154815290600101906020018083116101f957829003601f168201915b505050505081565b5f5ffd5b5f5ffd5b5f819050919050565b61023881610226565b8114610242575f5ffd5b50565b5f813590506102538161022f565b92915050565b5f6020828403121561026e5761026d61021e565b5b5f61027b84828501610245565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6102c682610284565b6102d0818561028e565b93506102e081856020860161029e565b6102e9816102ac565b840191505092915050565b5f6020820190508181035f83015261030c81846102bc565b905092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261033557610334610314565b5b8235905067ffffffffffffffff81111561035257610351610318565b5b60208301915083600182028301111561036e5761036d61031c565b5b9250929050565b5f5f5f6040848603121561038c5761038b61021e565b5b5f61039986828701610245565b935050602084013567ffffffffffffffff8111156103ba576103b9610222565b5b6103c686828701610320565b92509250509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061041657607f821691505b602082108103610429576104286103d2565b5b50919050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104c27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610487565b6104cc8683610487565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6105076105026104fd84610226565b6104e4565b610226565b9050919050565b5f819050919050565b610520836104ed565b61053461052c8261050e565b848454610493565b825550505050565b5f5f905090565b61054b61053c565b610556818484610517565b505050565b5b818110156105795761056e5f82610543565b60018101905061055c565b5050565b601f8211156105be5761058f81610466565b61059884610478565b810160208510156105a7578190505b6105bb6105b385610478565b83018261055b565b50505b505050565b5f82821c905092915050565b5f6105de5f19846008026105c3565b1980831691505092915050565b5f6105f683836105cf565b9150826002028217905092915050565b610610838361042f565b67ffffffffffffffff81111561062957610628610439565b5b61063382546103ff565b61063e82828561057d565b5f601f83116001811461066b575f8415610659578287013590505b61066385826105eb565b8655506106ca565b601f19841661067986610466565b5f5b828110156106a05784890135825560018201915060208501945060208101905061067b565b868310156106bd57848901356106b9601f8916826105cf565b8355505b6001600288020188555050505b5050505050505056fea26469706673582212200dfb091f0fc0d466fc5b32c6146ecea3e46c0f9461df0b7182313615151615d164736f6c634300081e0033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// Data is a free data retrieval call binding the contract method 0xf0ba8440.
//
// Solidity: function data(uint256 ) view returns(string)
func (_Storage *StorageCaller) Data(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "data", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Data is a free data retrieval call binding the contract method 0xf0ba8440.
//
// Solidity: function data(uint256 ) view returns(string)
func (_Storage *StorageSession) Data(arg0 *big.Int) (string, error) {
	return _Storage.Contract.Data(&_Storage.CallOpts, arg0)
}

// Data is a free data retrieval call binding the contract method 0xf0ba8440.
//
// Solidity: function data(uint256 ) view returns(string)
func (_Storage *StorageCallerSession) Data(arg0 *big.Int) (string, error) {
	return _Storage.Contract.Data(&_Storage.CallOpts, arg0)
}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 key) view returns(string)
func (_Storage *StorageCaller) GetData(opts *bind.CallOpts, key *big.Int) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getData", key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 key) view returns(string)
func (_Storage *StorageSession) GetData(key *big.Int) (string, error) {
	return _Storage.Contract.GetData(&_Storage.CallOpts, key)
}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 key) view returns(string)
func (_Storage *StorageCallerSession) GetData(key *big.Int) (string, error) {
	return _Storage.Contract.GetData(&_Storage.CallOpts, key)
}

// SetData is a paid mutator transaction binding the contract method 0x18562dae.
//
// Solidity: function setData(uint256 key, string value) returns()
func (_Storage *StorageTransactor) SetData(opts *bind.TransactOpts, key *big.Int, value string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setData", key, value)
}

// SetData is a paid mutator transaction binding the contract method 0x18562dae.
//
// Solidity: function setData(uint256 key, string value) returns()
func (_Storage *StorageSession) SetData(key *big.Int, value string) (*types.Transaction, error) {
	return _Storage.Contract.SetData(&_Storage.TransactOpts, key, value)
}

// SetData is a paid mutator transaction binding the contract method 0x18562dae.
//
// Solidity: function setData(uint256 key, string value) returns()
func (_Storage *StorageTransactorSession) SetData(key *big.Int, value string) (*types.Transaction, error) {
	return _Storage.Contract.SetData(&_Storage.TransactOpts, key, value)
}
