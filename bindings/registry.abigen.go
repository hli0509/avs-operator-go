// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// Quorum is an auto generated low-level Go binding around an user-defined struct.
type Quorum struct {
	Strategies []StrategyParams
}

// StrategyParams is an auto generated low-level Go binding around an user-defined struct.
type StrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLastCheckpointOperatorWeight\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastCheckpointThresholdWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastCheckpointThresholdWeightAtBlock\",\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastCheckpointTotalWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastCheckpointTotalWeightAtBlock\",\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeight\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeightAtBlock\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_serviceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_thresholdWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_quorum\",\"type\":\"tuple\",\"internalType\":\"structQuorum\",\"components\":[{\"name\":\"strategies\",\"type\":\"tuple[]\",\"internalType\":\"structStrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isValidSignature\",\"inputs\":[{\"name\":\"_dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signatureData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorRegistered\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structQuorum\",\"components\":[{\"name\":\"strategies\",\"type\":\"tuple[]\",\"internalType\":\"structStrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorWithSignature\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMinimumWeight\",\"inputs\":[{\"name\":\"_newMinimumWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[{\"name\":\"_operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorsForQuorum\",\"inputs\":[{\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateQuorumConfig\",\"inputs\":[{\"name\":\"_quorum\",\"type\":\"tuple\",\"internalType\":\"structQuorum\",\"components\":[{\"name\":\"strategies\",\"type\":\"tuple[]\",\"internalType\":\"structStrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]},{\"name\":\"_operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateStakeThreshold\",\"inputs\":[{\"name\":\"_thresholdWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinimumWeightUpdated\",\"inputs\":[{\"name\":\"_old\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_new\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorWeightUpdated\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumUpdated\",\"inputs\":[{\"name\":\"_old\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structQuorum\",\"components\":[{\"name\":\"strategies\",\"type\":\"tuple[]\",\"internalType\":\"structStrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]},{\"name\":\"_new\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structQuorum\",\"components\":[{\"name\":\"strategies\",\"type\":\"tuple[]\",\"internalType\":\"structStrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdWeightUpdated\",\"inputs\":[{\"name\":\"_thresholdWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalWeightUpdated\",\"inputs\":[{\"name\":\"oldTotalWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newTotalWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateMinimumWeight\",\"inputs\":[{\"name\":\"oldMinimumWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newMinimumWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InsufficientSignedStake\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientWeight\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidQuorum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignedWeight\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustUpdateAllOperators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]}]",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// GetLastCheckpointOperatorWeight is a free data retrieval call binding the contract method 0x3b242e4a.
//
// Solidity: function getLastCheckpointOperatorWeight(address _operator) view returns(uint256)
func (_Registry *RegistryCaller) GetLastCheckpointOperatorWeight(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getLastCheckpointOperatorWeight", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointOperatorWeight is a free data retrieval call binding the contract method 0x3b242e4a.
//
// Solidity: function getLastCheckpointOperatorWeight(address _operator) view returns(uint256)
func (_Registry *RegistrySession) GetLastCheckpointOperatorWeight(_operator common.Address) (*big.Int, error) {
	return _Registry.Contract.GetLastCheckpointOperatorWeight(&_Registry.CallOpts, _operator)
}

// GetLastCheckpointOperatorWeight is a free data retrieval call binding the contract method 0x3b242e4a.
//
// Solidity: function getLastCheckpointOperatorWeight(address _operator) view returns(uint256)
func (_Registry *RegistryCallerSession) GetLastCheckpointOperatorWeight(_operator common.Address) (*big.Int, error) {
	return _Registry.Contract.GetLastCheckpointOperatorWeight(&_Registry.CallOpts, _operator)
}

// GetLastCheckpointThresholdWeight is a free data retrieval call binding the contract method 0xb933fa74.
//
// Solidity: function getLastCheckpointThresholdWeight() view returns(uint256)
func (_Registry *RegistryCaller) GetLastCheckpointThresholdWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getLastCheckpointThresholdWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointThresholdWeight is a free data retrieval call binding the contract method 0xb933fa74.
//
// Solidity: function getLastCheckpointThresholdWeight() view returns(uint256)
func (_Registry *RegistrySession) GetLastCheckpointThresholdWeight() (*big.Int, error) {
	return _Registry.Contract.GetLastCheckpointThresholdWeight(&_Registry.CallOpts)
}

// GetLastCheckpointThresholdWeight is a free data retrieval call binding the contract method 0xb933fa74.
//
// Solidity: function getLastCheckpointThresholdWeight() view returns(uint256)
func (_Registry *RegistryCallerSession) GetLastCheckpointThresholdWeight() (*big.Int, error) {
	return _Registry.Contract.GetLastCheckpointThresholdWeight(&_Registry.CallOpts)
}

// GetLastCheckpointThresholdWeightAtBlock is a free data retrieval call binding the contract method 0x1e4cd85e.
//
// Solidity: function getLastCheckpointThresholdWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_Registry *RegistryCaller) GetLastCheckpointThresholdWeightAtBlock(opts *bind.CallOpts, _blockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getLastCheckpointThresholdWeightAtBlock", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointThresholdWeightAtBlock is a free data retrieval call binding the contract method 0x1e4cd85e.
//
// Solidity: function getLastCheckpointThresholdWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_Registry *RegistrySession) GetLastCheckpointThresholdWeightAtBlock(_blockNumber uint32) (*big.Int, error) {
	return _Registry.Contract.GetLastCheckpointThresholdWeightAtBlock(&_Registry.CallOpts, _blockNumber)
}

// GetLastCheckpointThresholdWeightAtBlock is a free data retrieval call binding the contract method 0x1e4cd85e.
//
// Solidity: function getLastCheckpointThresholdWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_Registry *RegistryCallerSession) GetLastCheckpointThresholdWeightAtBlock(_blockNumber uint32) (*big.Int, error) {
	return _Registry.Contract.GetLastCheckpointThresholdWeightAtBlock(&_Registry.CallOpts, _blockNumber)
}

// GetLastCheckpointTotalWeight is a free data retrieval call binding the contract method 0x314f3a49.
//
// Solidity: function getLastCheckpointTotalWeight() view returns(uint256)
func (_Registry *RegistryCaller) GetLastCheckpointTotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getLastCheckpointTotalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointTotalWeight is a free data retrieval call binding the contract method 0x314f3a49.
//
// Solidity: function getLastCheckpointTotalWeight() view returns(uint256)
func (_Registry *RegistrySession) GetLastCheckpointTotalWeight() (*big.Int, error) {
	return _Registry.Contract.GetLastCheckpointTotalWeight(&_Registry.CallOpts)
}

// GetLastCheckpointTotalWeight is a free data retrieval call binding the contract method 0x314f3a49.
//
// Solidity: function getLastCheckpointTotalWeight() view returns(uint256)
func (_Registry *RegistryCallerSession) GetLastCheckpointTotalWeight() (*big.Int, error) {
	return _Registry.Contract.GetLastCheckpointTotalWeight(&_Registry.CallOpts)
}

// GetLastCheckpointTotalWeightAtBlock is a free data retrieval call binding the contract method 0x0dba3394.
//
// Solidity: function getLastCheckpointTotalWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_Registry *RegistryCaller) GetLastCheckpointTotalWeightAtBlock(opts *bind.CallOpts, _blockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getLastCheckpointTotalWeightAtBlock", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointTotalWeightAtBlock is a free data retrieval call binding the contract method 0x0dba3394.
//
// Solidity: function getLastCheckpointTotalWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_Registry *RegistrySession) GetLastCheckpointTotalWeightAtBlock(_blockNumber uint32) (*big.Int, error) {
	return _Registry.Contract.GetLastCheckpointTotalWeightAtBlock(&_Registry.CallOpts, _blockNumber)
}

// GetLastCheckpointTotalWeightAtBlock is a free data retrieval call binding the contract method 0x0dba3394.
//
// Solidity: function getLastCheckpointTotalWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_Registry *RegistryCallerSession) GetLastCheckpointTotalWeightAtBlock(_blockNumber uint32) (*big.Int, error) {
	return _Registry.Contract.GetLastCheckpointTotalWeightAtBlock(&_Registry.CallOpts, _blockNumber)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x98ec1ac9.
//
// Solidity: function getOperatorWeight(address _operator) view returns(uint256)
func (_Registry *RegistryCaller) GetOperatorWeight(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getOperatorWeight", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x98ec1ac9.
//
// Solidity: function getOperatorWeight(address _operator) view returns(uint256)
func (_Registry *RegistrySession) GetOperatorWeight(_operator common.Address) (*big.Int, error) {
	return _Registry.Contract.GetOperatorWeight(&_Registry.CallOpts, _operator)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x98ec1ac9.
//
// Solidity: function getOperatorWeight(address _operator) view returns(uint256)
func (_Registry *RegistryCallerSession) GetOperatorWeight(_operator common.Address) (*big.Int, error) {
	return _Registry.Contract.GetOperatorWeight(&_Registry.CallOpts, _operator)
}

// GetOperatorWeightAtBlock is a free data retrieval call binding the contract method 0x955f2d90.
//
// Solidity: function getOperatorWeightAtBlock(address _operator, uint32 _blockNumber) view returns(uint256)
func (_Registry *RegistryCaller) GetOperatorWeightAtBlock(opts *bind.CallOpts, _operator common.Address, _blockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getOperatorWeightAtBlock", _operator, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWeightAtBlock is a free data retrieval call binding the contract method 0x955f2d90.
//
// Solidity: function getOperatorWeightAtBlock(address _operator, uint32 _blockNumber) view returns(uint256)
func (_Registry *RegistrySession) GetOperatorWeightAtBlock(_operator common.Address, _blockNumber uint32) (*big.Int, error) {
	return _Registry.Contract.GetOperatorWeightAtBlock(&_Registry.CallOpts, _operator, _blockNumber)
}

// GetOperatorWeightAtBlock is a free data retrieval call binding the contract method 0x955f2d90.
//
// Solidity: function getOperatorWeightAtBlock(address _operator, uint32 _blockNumber) view returns(uint256)
func (_Registry *RegistryCallerSession) GetOperatorWeightAtBlock(_operator common.Address, _blockNumber uint32) (*big.Int, error) {
	return _Registry.Contract.GetOperatorWeightAtBlock(&_Registry.CallOpts, _operator, _blockNumber)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signatureData) view returns(bytes4)
func (_Registry *RegistryCaller) IsValidSignature(opts *bind.CallOpts, _dataHash [32]byte, _signatureData []byte) ([4]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isValidSignature", _dataHash, _signatureData)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signatureData) view returns(bytes4)
func (_Registry *RegistrySession) IsValidSignature(_dataHash [32]byte, _signatureData []byte) ([4]byte, error) {
	return _Registry.Contract.IsValidSignature(&_Registry.CallOpts, _dataHash, _signatureData)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signatureData) view returns(bytes4)
func (_Registry *RegistryCallerSession) IsValidSignature(_dataHash [32]byte, _signatureData []byte) ([4]byte, error) {
	return _Registry.Contract.IsValidSignature(&_Registry.CallOpts, _dataHash, _signatureData)
}

// MinimumWeight is a free data retrieval call binding the contract method 0x40bf2fb7.
//
// Solidity: function minimumWeight() view returns(uint256)
func (_Registry *RegistryCaller) MinimumWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "minimumWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumWeight is a free data retrieval call binding the contract method 0x40bf2fb7.
//
// Solidity: function minimumWeight() view returns(uint256)
func (_Registry *RegistrySession) MinimumWeight() (*big.Int, error) {
	return _Registry.Contract.MinimumWeight(&_Registry.CallOpts)
}

// MinimumWeight is a free data retrieval call binding the contract method 0x40bf2fb7.
//
// Solidity: function minimumWeight() view returns(uint256)
func (_Registry *RegistryCallerSession) MinimumWeight() (*big.Int, error) {
	return _Registry.Contract.MinimumWeight(&_Registry.CallOpts)
}

// OperatorRegistered is a free data retrieval call binding the contract method 0xec7fbb31.
//
// Solidity: function operatorRegistered(address _operator) view returns(bool)
func (_Registry *RegistryCaller) OperatorRegistered(opts *bind.CallOpts, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "operatorRegistered", _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorRegistered is a free data retrieval call binding the contract method 0xec7fbb31.
//
// Solidity: function operatorRegistered(address _operator) view returns(bool)
func (_Registry *RegistrySession) OperatorRegistered(_operator common.Address) (bool, error) {
	return _Registry.Contract.OperatorRegistered(&_Registry.CallOpts, _operator)
}

// OperatorRegistered is a free data retrieval call binding the contract method 0xec7fbb31.
//
// Solidity: function operatorRegistered(address _operator) view returns(bool)
func (_Registry *RegistryCallerSession) OperatorRegistered(_operator common.Address) (bool, error) {
	return _Registry.Contract.OperatorRegistered(&_Registry.CallOpts, _operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(((address,uint96)[]))
func (_Registry *RegistryCaller) Quorum(opts *bind.CallOpts) (Quorum, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(Quorum), err
	}

	out0 := *abi.ConvertType(out[0], new(Quorum)).(*Quorum)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(((address,uint96)[]))
func (_Registry *RegistrySession) Quorum() (Quorum, error) {
	return _Registry.Contract.Quorum(&_Registry.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(((address,uint96)[]))
func (_Registry *RegistryCallerSession) Quorum() (Quorum, error) {
	return _Registry.Contract.Quorum(&_Registry.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x857dc190.
//
// Solidity: function deregisterOperator() returns()
func (_Registry *RegistryTransactor) DeregisterOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "deregisterOperator")
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x857dc190.
//
// Solidity: function deregisterOperator() returns()
func (_Registry *RegistrySession) DeregisterOperator() (*types.Transaction, error) {
	return _Registry.Contract.DeregisterOperator(&_Registry.TransactOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x857dc190.
//
// Solidity: function deregisterOperator() returns()
func (_Registry *RegistryTransactorSession) DeregisterOperator() (*types.Transaction, error) {
	return _Registry.Contract.DeregisterOperator(&_Registry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xab118995.
//
// Solidity: function initialize(address _serviceManager, uint256 _thresholdWeight, ((address,uint96)[]) _quorum) returns()
func (_Registry *RegistryTransactor) Initialize(opts *bind.TransactOpts, _serviceManager common.Address, _thresholdWeight *big.Int, _quorum Quorum) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "initialize", _serviceManager, _thresholdWeight, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0xab118995.
//
// Solidity: function initialize(address _serviceManager, uint256 _thresholdWeight, ((address,uint96)[]) _quorum) returns()
func (_Registry *RegistrySession) Initialize(_serviceManager common.Address, _thresholdWeight *big.Int, _quorum Quorum) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, _serviceManager, _thresholdWeight, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0xab118995.
//
// Solidity: function initialize(address _serviceManager, uint256 _thresholdWeight, ((address,uint96)[]) _quorum) returns()
func (_Registry *RegistryTransactorSession) Initialize(_serviceManager common.Address, _thresholdWeight *big.Int, _quorum Quorum) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, _serviceManager, _thresholdWeight, _quorum)
}

// RegisterOperatorWithSignature is a paid mutator transaction binding the contract method 0x0a601a12.
//
// Solidity: function registerOperatorWithSignature(address _operator, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_Registry *RegistryTransactor) RegisterOperatorWithSignature(opts *bind.TransactOpts, _operator common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerOperatorWithSignature", _operator, _operatorSignature)
}

// RegisterOperatorWithSignature is a paid mutator transaction binding the contract method 0x0a601a12.
//
// Solidity: function registerOperatorWithSignature(address _operator, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_Registry *RegistrySession) RegisterOperatorWithSignature(_operator common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Registry.Contract.RegisterOperatorWithSignature(&_Registry.TransactOpts, _operator, _operatorSignature)
}

// RegisterOperatorWithSignature is a paid mutator transaction binding the contract method 0x0a601a12.
//
// Solidity: function registerOperatorWithSignature(address _operator, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_Registry *RegistryTransactorSession) RegisterOperatorWithSignature(_operator common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _Registry.Contract.RegisterOperatorWithSignature(&_Registry.TransactOpts, _operator, _operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// UpdateMinimumWeight is a paid mutator transaction binding the contract method 0x696255be.
//
// Solidity: function updateMinimumWeight(uint256 _newMinimumWeight, address[] _operators) returns()
func (_Registry *RegistryTransactor) UpdateMinimumWeight(opts *bind.TransactOpts, _newMinimumWeight *big.Int, _operators []common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateMinimumWeight", _newMinimumWeight, _operators)
}

// UpdateMinimumWeight is a paid mutator transaction binding the contract method 0x696255be.
//
// Solidity: function updateMinimumWeight(uint256 _newMinimumWeight, address[] _operators) returns()
func (_Registry *RegistrySession) UpdateMinimumWeight(_newMinimumWeight *big.Int, _operators []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpdateMinimumWeight(&_Registry.TransactOpts, _newMinimumWeight, _operators)
}

// UpdateMinimumWeight is a paid mutator transaction binding the contract method 0x696255be.
//
// Solidity: function updateMinimumWeight(uint256 _newMinimumWeight, address[] _operators) returns()
func (_Registry *RegistryTransactorSession) UpdateMinimumWeight(_newMinimumWeight *big.Int, _operators []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpdateMinimumWeight(&_Registry.TransactOpts, _newMinimumWeight, _operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] _operators) returns()
func (_Registry *RegistryTransactor) UpdateOperators(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateOperators", _operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] _operators) returns()
func (_Registry *RegistrySession) UpdateOperators(_operators []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpdateOperators(&_Registry.TransactOpts, _operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] _operators) returns()
func (_Registry *RegistryTransactorSession) UpdateOperators(_operators []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpdateOperators(&_Registry.TransactOpts, _operators)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes ) returns()
func (_Registry *RegistryTransactor) UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateOperatorsForQuorum", operatorsPerQuorum, arg1)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes ) returns()
func (_Registry *RegistrySession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateOperatorsForQuorum(&_Registry.TransactOpts, operatorsPerQuorum, arg1)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes ) returns()
func (_Registry *RegistryTransactorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateOperatorsForQuorum(&_Registry.TransactOpts, operatorsPerQuorum, arg1)
}

// UpdateQuorumConfig is a paid mutator transaction binding the contract method 0xdec5d1f6.
//
// Solidity: function updateQuorumConfig(((address,uint96)[]) _quorum, address[] _operators) returns()
func (_Registry *RegistryTransactor) UpdateQuorumConfig(opts *bind.TransactOpts, _quorum Quorum, _operators []common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateQuorumConfig", _quorum, _operators)
}

// UpdateQuorumConfig is a paid mutator transaction binding the contract method 0xdec5d1f6.
//
// Solidity: function updateQuorumConfig(((address,uint96)[]) _quorum, address[] _operators) returns()
func (_Registry *RegistrySession) UpdateQuorumConfig(_quorum Quorum, _operators []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpdateQuorumConfig(&_Registry.TransactOpts, _quorum, _operators)
}

// UpdateQuorumConfig is a paid mutator transaction binding the contract method 0xdec5d1f6.
//
// Solidity: function updateQuorumConfig(((address,uint96)[]) _quorum, address[] _operators) returns()
func (_Registry *RegistryTransactorSession) UpdateQuorumConfig(_quorum Quorum, _operators []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpdateQuorumConfig(&_Registry.TransactOpts, _quorum, _operators)
}

// UpdateStakeThreshold is a paid mutator transaction binding the contract method 0x5ef53329.
//
// Solidity: function updateStakeThreshold(uint256 _thresholdWeight) returns()
func (_Registry *RegistryTransactor) UpdateStakeThreshold(opts *bind.TransactOpts, _thresholdWeight *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateStakeThreshold", _thresholdWeight)
}

// UpdateStakeThreshold is a paid mutator transaction binding the contract method 0x5ef53329.
//
// Solidity: function updateStakeThreshold(uint256 _thresholdWeight) returns()
func (_Registry *RegistrySession) UpdateStakeThreshold(_thresholdWeight *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UpdateStakeThreshold(&_Registry.TransactOpts, _thresholdWeight)
}

// UpdateStakeThreshold is a paid mutator transaction binding the contract method 0x5ef53329.
//
// Solidity: function updateStakeThreshold(uint256 _thresholdWeight) returns()
func (_Registry *RegistryTransactorSession) UpdateStakeThreshold(_thresholdWeight *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UpdateStakeThreshold(&_Registry.TransactOpts, _thresholdWeight)
}

// RegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Registry contract.
type RegistryInitializedIterator struct {
	Event *RegistryInitialized // Event containing the contract specifics and raw log

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
func (it *RegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryInitialized)
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
		it.Event = new(RegistryInitialized)
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
func (it *RegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryInitialized represents a Initialized event raised by the Registry contract.
type RegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registry *RegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*RegistryInitializedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RegistryInitializedIterator{contract: _Registry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registry *RegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryInitialized)
				if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registry *RegistryFilterer) ParseInitialized(log types.Log) (*RegistryInitialized, error) {
	event := new(RegistryInitialized)
	if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryMinimumWeightUpdatedIterator is returned from FilterMinimumWeightUpdated and is used to iterate over the raw logs and unpacked data for MinimumWeightUpdated events raised by the Registry contract.
type RegistryMinimumWeightUpdatedIterator struct {
	Event *RegistryMinimumWeightUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryMinimumWeightUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryMinimumWeightUpdated)
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
		it.Event = new(RegistryMinimumWeightUpdated)
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
func (it *RegistryMinimumWeightUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryMinimumWeightUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryMinimumWeightUpdated represents a MinimumWeightUpdated event raised by the Registry contract.
type RegistryMinimumWeightUpdated struct {
	Old *big.Int
	New *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMinimumWeightUpdated is a free log retrieval operation binding the contract event 0x713ca53b88d6eb63f5b1854cb8cbdd736ec51eda225e46791aa9298b0160648f.
//
// Solidity: event MinimumWeightUpdated(uint256 _old, uint256 _new)
func (_Registry *RegistryFilterer) FilterMinimumWeightUpdated(opts *bind.FilterOpts) (*RegistryMinimumWeightUpdatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "MinimumWeightUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistryMinimumWeightUpdatedIterator{contract: _Registry.contract, event: "MinimumWeightUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumWeightUpdated is a free log subscription operation binding the contract event 0x713ca53b88d6eb63f5b1854cb8cbdd736ec51eda225e46791aa9298b0160648f.
//
// Solidity: event MinimumWeightUpdated(uint256 _old, uint256 _new)
func (_Registry *RegistryFilterer) WatchMinimumWeightUpdated(opts *bind.WatchOpts, sink chan<- *RegistryMinimumWeightUpdated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "MinimumWeightUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryMinimumWeightUpdated)
				if err := _Registry.contract.UnpackLog(event, "MinimumWeightUpdated", log); err != nil {
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

// ParseMinimumWeightUpdated is a log parse operation binding the contract event 0x713ca53b88d6eb63f5b1854cb8cbdd736ec51eda225e46791aa9298b0160648f.
//
// Solidity: event MinimumWeightUpdated(uint256 _old, uint256 _new)
func (_Registry *RegistryFilterer) ParseMinimumWeightUpdated(log types.Log) (*RegistryMinimumWeightUpdated, error) {
	event := new(RegistryMinimumWeightUpdated)
	if err := _Registry.contract.UnpackLog(event, "MinimumWeightUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the Registry contract.
type RegistryOperatorDeregisteredIterator struct {
	Event *RegistryOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *RegistryOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOperatorDeregistered)
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
		it.Event = new(RegistryOperatorDeregistered)
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
func (it *RegistryOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOperatorDeregistered represents a OperatorDeregistered event raised by the Registry contract.
type RegistryOperatorDeregistered struct {
	Operator common.Address
	Avs      common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x31e0adfec71bccee37b6e83a90c2fedb17d8f1693fee863c4771e7bfe2aed580.
//
// Solidity: event OperatorDeregistered(address indexed _operator, address indexed _avs)
func (_Registry *RegistryFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, _operator []common.Address, _avs []common.Address) (*RegistryOperatorDeregisteredIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _avsRule []interface{}
	for _, _avsItem := range _avs {
		_avsRule = append(_avsRule, _avsItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OperatorDeregistered", _operatorRule, _avsRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOperatorDeregisteredIterator{contract: _Registry.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x31e0adfec71bccee37b6e83a90c2fedb17d8f1693fee863c4771e7bfe2aed580.
//
// Solidity: event OperatorDeregistered(address indexed _operator, address indexed _avs)
func (_Registry *RegistryFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *RegistryOperatorDeregistered, _operator []common.Address, _avs []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _avsRule []interface{}
	for _, _avsItem := range _avs {
		_avsRule = append(_avsRule, _avsItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OperatorDeregistered", _operatorRule, _avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOperatorDeregistered)
				if err := _Registry.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x31e0adfec71bccee37b6e83a90c2fedb17d8f1693fee863c4771e7bfe2aed580.
//
// Solidity: event OperatorDeregistered(address indexed _operator, address indexed _avs)
func (_Registry *RegistryFilterer) ParseOperatorDeregistered(log types.Log) (*RegistryOperatorDeregistered, error) {
	event := new(RegistryOperatorDeregistered)
	if err := _Registry.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the Registry contract.
type RegistryOperatorRegisteredIterator struct {
	Event *RegistryOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *RegistryOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOperatorRegistered)
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
		it.Event = new(RegistryOperatorRegistered)
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
func (it *RegistryOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOperatorRegistered represents a OperatorRegistered event raised by the Registry contract.
type RegistryOperatorRegistered struct {
	Operator common.Address
	Avs      common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed _operator, address indexed _avs)
func (_Registry *RegistryFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, _operator []common.Address, _avs []common.Address) (*RegistryOperatorRegisteredIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _avsRule []interface{}
	for _, _avsItem := range _avs {
		_avsRule = append(_avsRule, _avsItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OperatorRegistered", _operatorRule, _avsRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOperatorRegisteredIterator{contract: _Registry.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed _operator, address indexed _avs)
func (_Registry *RegistryFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *RegistryOperatorRegistered, _operator []common.Address, _avs []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _avsRule []interface{}
	for _, _avsItem := range _avs {
		_avsRule = append(_avsRule, _avsItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OperatorRegistered", _operatorRule, _avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOperatorRegistered)
				if err := _Registry.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed _operator, address indexed _avs)
func (_Registry *RegistryFilterer) ParseOperatorRegistered(log types.Log) (*RegistryOperatorRegistered, error) {
	event := new(RegistryOperatorRegistered)
	if err := _Registry.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryOperatorWeightUpdatedIterator is returned from FilterOperatorWeightUpdated and is used to iterate over the raw logs and unpacked data for OperatorWeightUpdated events raised by the Registry contract.
type RegistryOperatorWeightUpdatedIterator struct {
	Event *RegistryOperatorWeightUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryOperatorWeightUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOperatorWeightUpdated)
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
		it.Event = new(RegistryOperatorWeightUpdated)
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
func (it *RegistryOperatorWeightUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOperatorWeightUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOperatorWeightUpdated represents a OperatorWeightUpdated event raised by the Registry contract.
type RegistryOperatorWeightUpdated struct {
	Operator  common.Address
	OldWeight *big.Int
	NewWeight *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperatorWeightUpdated is a free log retrieval operation binding the contract event 0x88770dc862e47a7ed586907857eb1b75e4c5ffc8b707c7ee10eb74d6885fe594.
//
// Solidity: event OperatorWeightUpdated(address indexed _operator, uint256 oldWeight, uint256 newWeight)
func (_Registry *RegistryFilterer) FilterOperatorWeightUpdated(opts *bind.FilterOpts, _operator []common.Address) (*RegistryOperatorWeightUpdatedIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OperatorWeightUpdated", _operatorRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOperatorWeightUpdatedIterator{contract: _Registry.contract, event: "OperatorWeightUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorWeightUpdated is a free log subscription operation binding the contract event 0x88770dc862e47a7ed586907857eb1b75e4c5ffc8b707c7ee10eb74d6885fe594.
//
// Solidity: event OperatorWeightUpdated(address indexed _operator, uint256 oldWeight, uint256 newWeight)
func (_Registry *RegistryFilterer) WatchOperatorWeightUpdated(opts *bind.WatchOpts, sink chan<- *RegistryOperatorWeightUpdated, _operator []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OperatorWeightUpdated", _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOperatorWeightUpdated)
				if err := _Registry.contract.UnpackLog(event, "OperatorWeightUpdated", log); err != nil {
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

// ParseOperatorWeightUpdated is a log parse operation binding the contract event 0x88770dc862e47a7ed586907857eb1b75e4c5ffc8b707c7ee10eb74d6885fe594.
//
// Solidity: event OperatorWeightUpdated(address indexed _operator, uint256 oldWeight, uint256 newWeight)
func (_Registry *RegistryFilterer) ParseOperatorWeightUpdated(log types.Log) (*RegistryOperatorWeightUpdated, error) {
	event := new(RegistryOperatorWeightUpdated)
	if err := _Registry.contract.UnpackLog(event, "OperatorWeightUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryOwnershipTransferred, error) {
	event := new(RegistryOwnershipTransferred)
	if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryQuorumUpdatedIterator is returned from FilterQuorumUpdated and is used to iterate over the raw logs and unpacked data for QuorumUpdated events raised by the Registry contract.
type RegistryQuorumUpdatedIterator struct {
	Event *RegistryQuorumUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryQuorumUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryQuorumUpdated)
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
		it.Event = new(RegistryQuorumUpdated)
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
func (it *RegistryQuorumUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryQuorumUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryQuorumUpdated represents a QuorumUpdated event raised by the Registry contract.
type RegistryQuorumUpdated struct {
	Old Quorum
	New Quorum
	Raw types.Log // Blockchain specific contextual infos
}

// FilterQuorumUpdated is a free log retrieval operation binding the contract event 0x23aad4e61744ece164130aa415c1616e80136b0f0770e56589438b90b269265e.
//
// Solidity: event QuorumUpdated(((address,uint96)[]) _old, ((address,uint96)[]) _new)
func (_Registry *RegistryFilterer) FilterQuorumUpdated(opts *bind.FilterOpts) (*RegistryQuorumUpdatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "QuorumUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistryQuorumUpdatedIterator{contract: _Registry.contract, event: "QuorumUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumUpdated is a free log subscription operation binding the contract event 0x23aad4e61744ece164130aa415c1616e80136b0f0770e56589438b90b269265e.
//
// Solidity: event QuorumUpdated(((address,uint96)[]) _old, ((address,uint96)[]) _new)
func (_Registry *RegistryFilterer) WatchQuorumUpdated(opts *bind.WatchOpts, sink chan<- *RegistryQuorumUpdated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "QuorumUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryQuorumUpdated)
				if err := _Registry.contract.UnpackLog(event, "QuorumUpdated", log); err != nil {
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

// ParseQuorumUpdated is a log parse operation binding the contract event 0x23aad4e61744ece164130aa415c1616e80136b0f0770e56589438b90b269265e.
//
// Solidity: event QuorumUpdated(((address,uint96)[]) _old, ((address,uint96)[]) _new)
func (_Registry *RegistryFilterer) ParseQuorumUpdated(log types.Log) (*RegistryQuorumUpdated, error) {
	event := new(RegistryQuorumUpdated)
	if err := _Registry.contract.UnpackLog(event, "QuorumUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryThresholdWeightUpdatedIterator is returned from FilterThresholdWeightUpdated and is used to iterate over the raw logs and unpacked data for ThresholdWeightUpdated events raised by the Registry contract.
type RegistryThresholdWeightUpdatedIterator struct {
	Event *RegistryThresholdWeightUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryThresholdWeightUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryThresholdWeightUpdated)
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
		it.Event = new(RegistryThresholdWeightUpdated)
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
func (it *RegistryThresholdWeightUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryThresholdWeightUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryThresholdWeightUpdated represents a ThresholdWeightUpdated event raised by the Registry contract.
type RegistryThresholdWeightUpdated struct {
	ThresholdWeight *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterThresholdWeightUpdated is a free log retrieval operation binding the contract event 0x9324f7e5a7c0288808a634ccde44b8e979676474b22e29ee9dd569b55e791a4b.
//
// Solidity: event ThresholdWeightUpdated(uint256 _thresholdWeight)
func (_Registry *RegistryFilterer) FilterThresholdWeightUpdated(opts *bind.FilterOpts) (*RegistryThresholdWeightUpdatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ThresholdWeightUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistryThresholdWeightUpdatedIterator{contract: _Registry.contract, event: "ThresholdWeightUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdWeightUpdated is a free log subscription operation binding the contract event 0x9324f7e5a7c0288808a634ccde44b8e979676474b22e29ee9dd569b55e791a4b.
//
// Solidity: event ThresholdWeightUpdated(uint256 _thresholdWeight)
func (_Registry *RegistryFilterer) WatchThresholdWeightUpdated(opts *bind.WatchOpts, sink chan<- *RegistryThresholdWeightUpdated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ThresholdWeightUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryThresholdWeightUpdated)
				if err := _Registry.contract.UnpackLog(event, "ThresholdWeightUpdated", log); err != nil {
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

// ParseThresholdWeightUpdated is a log parse operation binding the contract event 0x9324f7e5a7c0288808a634ccde44b8e979676474b22e29ee9dd569b55e791a4b.
//
// Solidity: event ThresholdWeightUpdated(uint256 _thresholdWeight)
func (_Registry *RegistryFilterer) ParseThresholdWeightUpdated(log types.Log) (*RegistryThresholdWeightUpdated, error) {
	event := new(RegistryThresholdWeightUpdated)
	if err := _Registry.contract.UnpackLog(event, "ThresholdWeightUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryTotalWeightUpdatedIterator is returned from FilterTotalWeightUpdated and is used to iterate over the raw logs and unpacked data for TotalWeightUpdated events raised by the Registry contract.
type RegistryTotalWeightUpdatedIterator struct {
	Event *RegistryTotalWeightUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryTotalWeightUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryTotalWeightUpdated)
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
		it.Event = new(RegistryTotalWeightUpdated)
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
func (it *RegistryTotalWeightUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryTotalWeightUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryTotalWeightUpdated represents a TotalWeightUpdated event raised by the Registry contract.
type RegistryTotalWeightUpdated struct {
	OldTotalWeight *big.Int
	NewTotalWeight *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalWeightUpdated is a free log retrieval operation binding the contract event 0x86dcf86b12dfeedea74ae9300dbdaa193bcce5809369c8177ea2f4eaaa65729b.
//
// Solidity: event TotalWeightUpdated(uint256 oldTotalWeight, uint256 newTotalWeight)
func (_Registry *RegistryFilterer) FilterTotalWeightUpdated(opts *bind.FilterOpts) (*RegistryTotalWeightUpdatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "TotalWeightUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistryTotalWeightUpdatedIterator{contract: _Registry.contract, event: "TotalWeightUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalWeightUpdated is a free log subscription operation binding the contract event 0x86dcf86b12dfeedea74ae9300dbdaa193bcce5809369c8177ea2f4eaaa65729b.
//
// Solidity: event TotalWeightUpdated(uint256 oldTotalWeight, uint256 newTotalWeight)
func (_Registry *RegistryFilterer) WatchTotalWeightUpdated(opts *bind.WatchOpts, sink chan<- *RegistryTotalWeightUpdated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "TotalWeightUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryTotalWeightUpdated)
				if err := _Registry.contract.UnpackLog(event, "TotalWeightUpdated", log); err != nil {
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

// ParseTotalWeightUpdated is a log parse operation binding the contract event 0x86dcf86b12dfeedea74ae9300dbdaa193bcce5809369c8177ea2f4eaaa65729b.
//
// Solidity: event TotalWeightUpdated(uint256 oldTotalWeight, uint256 newTotalWeight)
func (_Registry *RegistryFilterer) ParseTotalWeightUpdated(log types.Log) (*RegistryTotalWeightUpdated, error) {
	event := new(RegistryTotalWeightUpdated)
	if err := _Registry.contract.UnpackLog(event, "TotalWeightUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryUpdateMinimumWeightIterator is returned from FilterUpdateMinimumWeight and is used to iterate over the raw logs and unpacked data for UpdateMinimumWeight events raised by the Registry contract.
type RegistryUpdateMinimumWeightIterator struct {
	Event *RegistryUpdateMinimumWeight // Event containing the contract specifics and raw log

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
func (it *RegistryUpdateMinimumWeightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryUpdateMinimumWeight)
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
		it.Event = new(RegistryUpdateMinimumWeight)
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
func (it *RegistryUpdateMinimumWeightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryUpdateMinimumWeightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryUpdateMinimumWeight represents a UpdateMinimumWeight event raised by the Registry contract.
type RegistryUpdateMinimumWeight struct {
	OldMinimumWeight *big.Int
	NewMinimumWeight *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinimumWeight is a free log retrieval operation binding the contract event 0x1ea42186b305fa37310450d9fb87ea1e8f0c7f447e771479e3b27634bfe84dc1.
//
// Solidity: event UpdateMinimumWeight(uint256 oldMinimumWeight, uint256 newMinimumWeight)
func (_Registry *RegistryFilterer) FilterUpdateMinimumWeight(opts *bind.FilterOpts) (*RegistryUpdateMinimumWeightIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "UpdateMinimumWeight")
	if err != nil {
		return nil, err
	}
	return &RegistryUpdateMinimumWeightIterator{contract: _Registry.contract, event: "UpdateMinimumWeight", logs: logs, sub: sub}, nil
}

// WatchUpdateMinimumWeight is a free log subscription operation binding the contract event 0x1ea42186b305fa37310450d9fb87ea1e8f0c7f447e771479e3b27634bfe84dc1.
//
// Solidity: event UpdateMinimumWeight(uint256 oldMinimumWeight, uint256 newMinimumWeight)
func (_Registry *RegistryFilterer) WatchUpdateMinimumWeight(opts *bind.WatchOpts, sink chan<- *RegistryUpdateMinimumWeight) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "UpdateMinimumWeight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryUpdateMinimumWeight)
				if err := _Registry.contract.UnpackLog(event, "UpdateMinimumWeight", log); err != nil {
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

// ParseUpdateMinimumWeight is a log parse operation binding the contract event 0x1ea42186b305fa37310450d9fb87ea1e8f0c7f447e771479e3b27634bfe84dc1.
//
// Solidity: event UpdateMinimumWeight(uint256 oldMinimumWeight, uint256 newMinimumWeight)
func (_Registry *RegistryFilterer) ParseUpdateMinimumWeight(log types.Log) (*RegistryUpdateMinimumWeight, error) {
	event := new(RegistryUpdateMinimumWeight)
	if err := _Registry.contract.UnpackLog(event, "UpdateMinimumWeight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
