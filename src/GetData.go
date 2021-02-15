// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// GetDataABI is the input ABI used to generate the binding from.
const GetDataABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"dataMapping\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_index\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_calldata\",\"type\":\"string\"}],\"name\":\"send\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GetData is an auto generated Go binding around an Ethereum contract.
type GetData struct {
	GetDataCaller     // Read-only binding to the contract
	GetDataTransactor // Write-only binding to the contract
	GetDataFilterer   // Log filterer for contract events
}

// GetDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type GetDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GetDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GetDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GetDataSession struct {
	Contract     *GetData          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GetDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GetDataCallerSession struct {
	Contract *GetDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GetDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GetDataTransactorSession struct {
	Contract     *GetDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GetDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type GetDataRaw struct {
	Contract *GetData // Generic contract binding to access the raw methods on
}

// GetDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GetDataCallerRaw struct {
	Contract *GetDataCaller // Generic read-only contract binding to access the raw methods on
}

// GetDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GetDataTransactorRaw struct {
	Contract *GetDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGetData creates a new instance of GetData, bound to a specific deployed contract.
func NewGetData(address common.Address, backend bind.ContractBackend) (*GetData, error) {
	contract, err := bindGetData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GetData{GetDataCaller: GetDataCaller{contract: contract}, GetDataTransactor: GetDataTransactor{contract: contract}, GetDataFilterer: GetDataFilterer{contract: contract}}, nil
}

// NewGetDataCaller creates a new read-only instance of GetData, bound to a specific deployed contract.
func NewGetDataCaller(address common.Address, caller bind.ContractCaller) (*GetDataCaller, error) {
	contract, err := bindGetData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GetDataCaller{contract: contract}, nil
}

// NewGetDataTransactor creates a new write-only instance of GetData, bound to a specific deployed contract.
func NewGetDataTransactor(address common.Address, transactor bind.ContractTransactor) (*GetDataTransactor, error) {
	contract, err := bindGetData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GetDataTransactor{contract: contract}, nil
}

// NewGetDataFilterer creates a new log filterer instance of GetData, bound to a specific deployed contract.
func NewGetDataFilterer(address common.Address, filterer bind.ContractFilterer) (*GetDataFilterer, error) {
	contract, err := bindGetData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GetDataFilterer{contract: contract}, nil
}

// bindGetData binds a generic wrapper to an already deployed contract.
func bindGetData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GetDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetData *GetDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GetData.Contract.GetDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetData *GetDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetData.Contract.GetDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetData *GetDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetData.Contract.GetDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetData *GetDataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GetData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetData *GetDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetData *GetDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetData.Contract.contract.Transact(opts, method, params...)
}

// DataMapping is a free data retrieval call binding the contract method 0x158ea8f5.
//
// Solidity: function dataMapping(string ) view returns(string)
func (_GetData *GetDataCaller) DataMapping(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _GetData.contract.Call(opts, &out, "dataMapping", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DataMapping is a free data retrieval call binding the contract method 0x158ea8f5.
//
// Solidity: function dataMapping(string ) view returns(string)
func (_GetData *GetDataSession) DataMapping(arg0 string) (string, error) {
	return _GetData.Contract.DataMapping(&_GetData.CallOpts, arg0)
}

// DataMapping is a free data retrieval call binding the contract method 0x158ea8f5.
//
// Solidity: function dataMapping(string ) view returns(string)
func (_GetData *GetDataCallerSession) DataMapping(arg0 string) (string, error) {
	return _GetData.Contract.DataMapping(&_GetData.CallOpts, arg0)
}

// Send is a paid mutator transaction binding the contract method 0xbd6de11c.
//
// Solidity: function send(string _index, string _calldata) returns()
func (_GetData *GetDataTransactor) Send(opts *bind.TransactOpts, _index string, _calldata string) (*types.Transaction, error) {
	return _GetData.contract.Transact(opts, "send", _index, _calldata)
}

// Send is a paid mutator transaction binding the contract method 0xbd6de11c.
//
// Solidity: function send(string _index, string _calldata) returns()
func (_GetData *GetDataSession) Send(_index string, _calldata string) (*types.Transaction, error) {
	return _GetData.Contract.Send(&_GetData.TransactOpts, _index, _calldata)
}

// Send is a paid mutator transaction binding the contract method 0xbd6de11c.
//
// Solidity: function send(string _index, string _calldata) returns()
func (_GetData *GetDataTransactorSession) Send(_index string, _calldata string) (*types.Transaction, error) {
	return _GetData.Contract.Send(&_GetData.TransactOpts, _index, _calldata)
}
