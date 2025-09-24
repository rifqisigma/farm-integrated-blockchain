// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package messagestore

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

// MessagestoreMetaData contains all meta data concerning the Messagestore contract.
var MessagestoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"MessageCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_text\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"createMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610ddf806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630d80fefd1461005157806386f79edb14610086578063a87d942c146100bb578063b264d171146100d9575b600080fd5b61006b600480360381019061006691906106ce565b6100f5565b60405161007d969594939291906107db565b60405180910390f35b6100a0600480360381019061009b91906106ce565b610291565b6040516100b2969594939291906107db565b60405180910390f35b6100c36104e3565b6040516100d0919061084a565b60405180910390f35b6100f360048036038101906100ee91906109c6565b6104ef565b005b6000818154811061010557600080fd5b906000526020600020906006020160009150905080600001549080600101805461012e90610a94565b80601f016020809104026020016040519081016040528092919081815260200182805461015a90610a94565b80156101a75780601f1061017c576101008083540402835291602001916101a7565b820191906000526020600020905b81548152906001019060200180831161018a57829003601f168201915b5050505050908060020180546101bc90610a94565b80601f01602080910402602001604051908101604052809291908181526020018280546101e890610a94565b80156102355780601f1061020a57610100808354040283529160200191610235565b820191906000526020600020905b81548152906001019060200180831161021857829003601f168201915b5050505050908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905086565b6000606080600080600080600088815481106102b0576102af610ac5565b5b90600052602060002090600602016040518060c0016040529081600082015481526020016001820180546102e390610a94565b80601f016020809104026020016040519081016040528092919081815260200182805461030f90610a94565b801561035c5780601f106103315761010080835404028352916020019161035c565b820191906000526020600020905b81548152906001019060200180831161033f57829003601f168201915b5050505050815260200160028201805461037590610a94565b80601f01602080910402602001604051908101604052809291908181526020018280546103a190610a94565b80156103ee5780601f106103c3576101008083540402835291602001916103ee565b820191906000526020600020905b8154815290600101906020018083116103d157829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815250509050806000015181602001518260400151836060015184608001518560a001519650965096509650965096505091939550919395565b60008080549050905090565b60006040518060c001604052808681526020018581526020018481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020014281525090806001815401808255809150506001900390600052602060002090600602016000909190919091506000820151816000015560208201518160010190816105929190610ca0565b5060408201518160020190816105a89190610ca0565b5060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816005015550507f4f72a37d9f9e1bf477dadee4a4e8da66b6cb01c05674ec446e688fe31b467dc384338360405161067693929190610d72565b60405180910390a150505050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6106ab81610698565b81146106b657600080fd5b50565b6000813590506106c8816106a2565b92915050565b6000602082840312156106e4576106e361068e565b5b60006106f2848285016106b9565b91505092915050565b61070481610698565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610744578082015181840152602081019050610729565b60008484015250505050565b6000601f19601f8301169050919050565b600061076c8261070a565b6107768185610715565b9350610786818560208601610726565b61078f81610750565b840191505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006107c58261079a565b9050919050565b6107d5816107ba565b82525050565b600060c0820190506107f060008301896106fb565b81810360208301526108028188610761565b905081810360408301526108168187610761565b905061082560608301866107cc565b61083260808301856107cc565b61083f60a08301846106fb565b979650505050505050565b600060208201905061085f60008301846106fb565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6108a782610750565b810181811067ffffffffffffffff821117156108c6576108c561086f565b5b80604052505050565b60006108d9610684565b90506108e5828261089e565b919050565b600067ffffffffffffffff8211156109055761090461086f565b5b61090e82610750565b9050602081019050919050565b82818337600083830152505050565b600061093d610938846108ea565b6108cf565b9050828152602081018484840111156109595761095861086a565b5b61096484828561091b565b509392505050565b600082601f83011261098157610980610865565b5b813561099184826020860161092a565b91505092915050565b6109a3816107ba565b81146109ae57600080fd5b50565b6000813590506109c08161099a565b92915050565b600080600080608085870312156109e0576109df61068e565b5b60006109ee878288016106b9565b945050602085013567ffffffffffffffff811115610a0f57610a0e610693565b5b610a1b8782880161096c565b935050604085013567ffffffffffffffff811115610a3c57610a3b610693565b5b610a488782880161096c565b9250506060610a59878288016109b1565b91505092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610aac57607f821691505b602082108103610abf57610abe610a65565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610b567fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610b19565b610b608683610b19565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610b9d610b98610b9384610698565b610b78565b610698565b9050919050565b6000819050919050565b610bb783610b82565b610bcb610bc382610ba4565b848454610b26565b825550505050565b600090565b610be0610bd3565b610beb818484610bae565b505050565b5b81811015610c0f57610c04600082610bd8565b600181019050610bf1565b5050565b601f821115610c5457610c2581610af4565b610c2e84610b09565b81016020851015610c3d578190505b610c51610c4985610b09565b830182610bf0565b50505b505050565b600082821c905092915050565b6000610c7760001984600802610c59565b1980831691505092915050565b6000610c908383610c66565b9150826002028217905092915050565b610ca98261070a565b67ffffffffffffffff811115610cc257610cc161086f565b5b610ccc8254610a94565b610cd7828285610c13565b600060209050601f831160018114610d0a5760008415610cf8578287015190505b610d028582610c84565b865550610d6a565b601f198416610d1886610af4565b60005b82811015610d4057848901518255600182019150602085019450602081019050610d1b565b86831015610d5d5784890151610d59601f891682610c66565b8355505b6001600288020188555050505b505050505050565b6000606082019050610d8760008301866106fb565b610d9460208301856107cc565b610da160408301846107cc565b94935050505056fea264697066735822122017cc721ff6b046b23c16161ffae1c24cf0c1d53b4a5b821e8cbec5f156c8783a64736f6c63430008130033",
}

// MessagestoreABI is the input ABI used to generate the binding from.
// Deprecated: Use MessagestoreMetaData.ABI instead.
var MessagestoreABI = MessagestoreMetaData.ABI

// MessagestoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessagestoreMetaData.Bin instead.
var MessagestoreBin = MessagestoreMetaData.Bin

// DeployMessagestore deploys a new Ethereum contract, binding an instance of Messagestore to it.
func DeployMessagestore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Messagestore, error) {
	parsed, err := MessagestoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessagestoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Messagestore{MessagestoreCaller: MessagestoreCaller{contract: contract}, MessagestoreTransactor: MessagestoreTransactor{contract: contract}, MessagestoreFilterer: MessagestoreFilterer{contract: contract}}, nil
}

// Messagestore is an auto generated Go binding around an Ethereum contract.
type Messagestore struct {
	MessagestoreCaller     // Read-only binding to the contract
	MessagestoreTransactor // Write-only binding to the contract
	MessagestoreFilterer   // Log filterer for contract events
}

// MessagestoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessagestoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagestoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessagestoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagestoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessagestoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagestoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessagestoreSession struct {
	Contract     *Messagestore     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessagestoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessagestoreCallerSession struct {
	Contract *MessagestoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MessagestoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessagestoreTransactorSession struct {
	Contract     *MessagestoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MessagestoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessagestoreRaw struct {
	Contract *Messagestore // Generic contract binding to access the raw methods on
}

// MessagestoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessagestoreCallerRaw struct {
	Contract *MessagestoreCaller // Generic read-only contract binding to access the raw methods on
}

// MessagestoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessagestoreTransactorRaw struct {
	Contract *MessagestoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessagestore creates a new instance of Messagestore, bound to a specific deployed contract.
func NewMessagestore(address common.Address, backend bind.ContractBackend) (*Messagestore, error) {
	contract, err := bindMessagestore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Messagestore{MessagestoreCaller: MessagestoreCaller{contract: contract}, MessagestoreTransactor: MessagestoreTransactor{contract: contract}, MessagestoreFilterer: MessagestoreFilterer{contract: contract}}, nil
}

// NewMessagestoreCaller creates a new read-only instance of Messagestore, bound to a specific deployed contract.
func NewMessagestoreCaller(address common.Address, caller bind.ContractCaller) (*MessagestoreCaller, error) {
	contract, err := bindMessagestore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessagestoreCaller{contract: contract}, nil
}

// NewMessagestoreTransactor creates a new write-only instance of Messagestore, bound to a specific deployed contract.
func NewMessagestoreTransactor(address common.Address, transactor bind.ContractTransactor) (*MessagestoreTransactor, error) {
	contract, err := bindMessagestore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessagestoreTransactor{contract: contract}, nil
}

// NewMessagestoreFilterer creates a new log filterer instance of Messagestore, bound to a specific deployed contract.
func NewMessagestoreFilterer(address common.Address, filterer bind.ContractFilterer) (*MessagestoreFilterer, error) {
	contract, err := bindMessagestore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessagestoreFilterer{contract: contract}, nil
}

// bindMessagestore binds a generic wrapper to an already deployed contract.
func bindMessagestore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessagestoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messagestore *MessagestoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Messagestore.Contract.MessagestoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messagestore *MessagestoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messagestore.Contract.MessagestoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messagestore *MessagestoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messagestore.Contract.MessagestoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messagestore *MessagestoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Messagestore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messagestore *MessagestoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messagestore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messagestore *MessagestoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messagestore.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Messagestore *MessagestoreCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Messagestore.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Messagestore *MessagestoreSession) GetCount() (*big.Int, error) {
	return _Messagestore.Contract.GetCount(&_Messagestore.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Messagestore *MessagestoreCallerSession) GetCount() (*big.Int, error) {
	return _Messagestore.Contract.GetCount(&_Messagestore.CallOpts)
}

// GetMessage is a free data retrieval call binding the contract method 0x86f79edb.
//
// Solidity: function getMessage(uint256 index) view returns(uint256, string, string, address, address, uint256)
func (_Messagestore *MessagestoreCaller) GetMessage(opts *bind.CallOpts, index *big.Int) (*big.Int, string, string, common.Address, common.Address, *big.Int, error) {
	var out []interface{}
	err := _Messagestore.contract.Call(opts, &out, "getMessage", index)

	if err != nil {
		return *new(*big.Int), *new(string), *new(string), *new(common.Address), *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetMessage is a free data retrieval call binding the contract method 0x86f79edb.
//
// Solidity: function getMessage(uint256 index) view returns(uint256, string, string, address, address, uint256)
func (_Messagestore *MessagestoreSession) GetMessage(index *big.Int) (*big.Int, string, string, common.Address, common.Address, *big.Int, error) {
	return _Messagestore.Contract.GetMessage(&_Messagestore.CallOpts, index)
}

// GetMessage is a free data retrieval call binding the contract method 0x86f79edb.
//
// Solidity: function getMessage(uint256 index) view returns(uint256, string, string, address, address, uint256)
func (_Messagestore *MessagestoreCallerSession) GetMessage(index *big.Int) (*big.Int, string, string, common.Address, common.Address, *big.Int, error) {
	return _Messagestore.Contract.GetMessage(&_Messagestore.CallOpts, index)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(uint256 id, string status, string text, address from, address to, uint256 createdAt)
func (_Messagestore *MessagestoreCaller) Messages(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Status    string
	Text      string
	From      common.Address
	To        common.Address
	CreatedAt *big.Int
}, error) {
	var out []interface{}
	err := _Messagestore.contract.Call(opts, &out, "messages", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Status    string
		Text      string
		From      common.Address
		To        common.Address
		CreatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Text = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.From = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.To = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(uint256 id, string status, string text, address from, address to, uint256 createdAt)
func (_Messagestore *MessagestoreSession) Messages(arg0 *big.Int) (struct {
	Id        *big.Int
	Status    string
	Text      string
	From      common.Address
	To        common.Address
	CreatedAt *big.Int
}, error) {
	return _Messagestore.Contract.Messages(&_Messagestore.CallOpts, arg0)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(uint256 id, string status, string text, address from, address to, uint256 createdAt)
func (_Messagestore *MessagestoreCallerSession) Messages(arg0 *big.Int) (struct {
	Id        *big.Int
	Status    string
	Text      string
	From      common.Address
	To        common.Address
	CreatedAt *big.Int
}, error) {
	return _Messagestore.Contract.Messages(&_Messagestore.CallOpts, arg0)
}

// CreateMessage is a paid mutator transaction binding the contract method 0xb264d171.
//
// Solidity: function createMessage(uint256 _id, string _status, string _text, address _to) returns()
func (_Messagestore *MessagestoreTransactor) CreateMessage(opts *bind.TransactOpts, _id *big.Int, _status string, _text string, _to common.Address) (*types.Transaction, error) {
	return _Messagestore.contract.Transact(opts, "createMessage", _id, _status, _text, _to)
}

// CreateMessage is a paid mutator transaction binding the contract method 0xb264d171.
//
// Solidity: function createMessage(uint256 _id, string _status, string _text, address _to) returns()
func (_Messagestore *MessagestoreSession) CreateMessage(_id *big.Int, _status string, _text string, _to common.Address) (*types.Transaction, error) {
	return _Messagestore.Contract.CreateMessage(&_Messagestore.TransactOpts, _id, _status, _text, _to)
}

// CreateMessage is a paid mutator transaction binding the contract method 0xb264d171.
//
// Solidity: function createMessage(uint256 _id, string _status, string _text, address _to) returns()
func (_Messagestore *MessagestoreTransactorSession) CreateMessage(_id *big.Int, _status string, _text string, _to common.Address) (*types.Transaction, error) {
	return _Messagestore.Contract.CreateMessage(&_Messagestore.TransactOpts, _id, _status, _text, _to)
}

// MessagestoreMessageCreatedIterator is returned from FilterMessageCreated and is used to iterate over the raw logs and unpacked data for MessageCreated events raised by the Messagestore contract.
type MessagestoreMessageCreatedIterator struct {
	Event *MessagestoreMessageCreated // Event containing the contract specifics and raw log

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
func (it *MessagestoreMessageCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessagestoreMessageCreated)
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
		it.Event = new(MessagestoreMessageCreated)
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
func (it *MessagestoreMessageCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessagestoreMessageCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessagestoreMessageCreated represents a MessageCreated event raised by the Messagestore contract.
type MessagestoreMessageCreated struct {
	Id   *big.Int
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMessageCreated is a free log retrieval operation binding the contract event 0x4f72a37d9f9e1bf477dadee4a4e8da66b6cb01c05674ec446e688fe31b467dc3.
//
// Solidity: event MessageCreated(uint256 id, address from, address to)
func (_Messagestore *MessagestoreFilterer) FilterMessageCreated(opts *bind.FilterOpts) (*MessagestoreMessageCreatedIterator, error) {

	logs, sub, err := _Messagestore.contract.FilterLogs(opts, "MessageCreated")
	if err != nil {
		return nil, err
	}
	return &MessagestoreMessageCreatedIterator{contract: _Messagestore.contract, event: "MessageCreated", logs: logs, sub: sub}, nil
}

// WatchMessageCreated is a free log subscription operation binding the contract event 0x4f72a37d9f9e1bf477dadee4a4e8da66b6cb01c05674ec446e688fe31b467dc3.
//
// Solidity: event MessageCreated(uint256 id, address from, address to)
func (_Messagestore *MessagestoreFilterer) WatchMessageCreated(opts *bind.WatchOpts, sink chan<- *MessagestoreMessageCreated) (event.Subscription, error) {

	logs, sub, err := _Messagestore.contract.WatchLogs(opts, "MessageCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessagestoreMessageCreated)
				if err := _Messagestore.contract.UnpackLog(event, "MessageCreated", log); err != nil {
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

// ParseMessageCreated is a log parse operation binding the contract event 0x4f72a37d9f9e1bf477dadee4a4e8da66b6cb01c05674ec446e688fe31b467dc3.
//
// Solidity: event MessageCreated(uint256 id, address from, address to)
func (_Messagestore *MessagestoreFilterer) ParseMessageCreated(log types.Log) (*MessagestoreMessageCreated, error) {
	event := new(MessagestoreMessageCreated)
	if err := _Messagestore.contract.UnpackLog(event, "MessageCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
