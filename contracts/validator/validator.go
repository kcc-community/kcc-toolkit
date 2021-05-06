// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validator

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

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogAddToTopValidators\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogCreateValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogDistributeBlockReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogEditValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogReactive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogRemoveFromTopValidators\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogRemoveValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogRemoveValidatorIncoming\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"staking\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogUnStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newSet\",\"type\":\"address[]\"}],\"name\":\"LogUpdateValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawProfits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawStaking\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FeeRecoder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MaxValidators\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinimalStakingCoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ProposalContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PunishContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StakingLockPeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValidatorContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WithdrawProfitPeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"feeAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"createOrEditValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeBlockReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getStakingInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStakeOfActiveValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getValidatorDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumValidators.Status\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestValidatorsSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vals\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"isActiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"isTopValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"removeValidatorIncoming\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalJailed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"tryReactive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"unStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"updateActiveValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"validateDescription\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawProfits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawStaking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// FeeRecoder is a free data retrieval call binding the contract method 0xabfcb13f.
//
// Solidity: function FeeRecoder() view returns(address)
func (_Contracts *ContractsCaller) FeeRecoder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "FeeRecoder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecoder is a free data retrieval call binding the contract method 0xabfcb13f.
//
// Solidity: function FeeRecoder() view returns(address)
func (_Contracts *ContractsSession) FeeRecoder() (common.Address, error) {
	return _Contracts.Contract.FeeRecoder(&_Contracts.CallOpts)
}

// FeeRecoder is a free data retrieval call binding the contract method 0xabfcb13f.
//
// Solidity: function FeeRecoder() view returns(address)
func (_Contracts *ContractsCallerSession) FeeRecoder() (common.Address, error) {
	return _Contracts.Contract.FeeRecoder(&_Contracts.CallOpts)
}

// MaxValidators is a free data retrieval call binding the contract method 0xc967f90f.
//
// Solidity: function MaxValidators() view returns(uint16)
func (_Contracts *ContractsCaller) MaxValidators(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "MaxValidators")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxValidators is a free data retrieval call binding the contract method 0xc967f90f.
//
// Solidity: function MaxValidators() view returns(uint16)
func (_Contracts *ContractsSession) MaxValidators() (uint16, error) {
	return _Contracts.Contract.MaxValidators(&_Contracts.CallOpts)
}

// MaxValidators is a free data retrieval call binding the contract method 0xc967f90f.
//
// Solidity: function MaxValidators() view returns(uint16)
func (_Contracts *ContractsCallerSession) MaxValidators() (uint16, error) {
	return _Contracts.Contract.MaxValidators(&_Contracts.CallOpts)
}

// MinimalStakingCoin is a free data retrieval call binding the contract method 0xbe645692.
//
// Solidity: function MinimalStakingCoin() view returns(uint256)
func (_Contracts *ContractsCaller) MinimalStakingCoin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "MinimalStakingCoin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimalStakingCoin is a free data retrieval call binding the contract method 0xbe645692.
//
// Solidity: function MinimalStakingCoin() view returns(uint256)
func (_Contracts *ContractsSession) MinimalStakingCoin() (*big.Int, error) {
	return _Contracts.Contract.MinimalStakingCoin(&_Contracts.CallOpts)
}

// MinimalStakingCoin is a free data retrieval call binding the contract method 0xbe645692.
//
// Solidity: function MinimalStakingCoin() view returns(uint256)
func (_Contracts *ContractsCallerSession) MinimalStakingCoin() (*big.Int, error) {
	return _Contracts.Contract.MinimalStakingCoin(&_Contracts.CallOpts)
}

// ProposalContractAddr is a free data retrieval call binding the contract method 0x60ae528c.
//
// Solidity: function ProposalContractAddr() view returns(address)
func (_Contracts *ContractsCaller) ProposalContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ProposalContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalContractAddr is a free data retrieval call binding the contract method 0x60ae528c.
//
// Solidity: function ProposalContractAddr() view returns(address)
func (_Contracts *ContractsSession) ProposalContractAddr() (common.Address, error) {
	return _Contracts.Contract.ProposalContractAddr(&_Contracts.CallOpts)
}

// ProposalContractAddr is a free data retrieval call binding the contract method 0x60ae528c.
//
// Solidity: function ProposalContractAddr() view returns(address)
func (_Contracts *ContractsCallerSession) ProposalContractAddr() (common.Address, error) {
	return _Contracts.Contract.ProposalContractAddr(&_Contracts.CallOpts)
}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Contracts *ContractsCaller) PunishContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "PunishContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Contracts *ContractsSession) PunishContractAddr() (common.Address, error) {
	return _Contracts.Contract.PunishContractAddr(&_Contracts.CallOpts)
}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Contracts *ContractsCallerSession) PunishContractAddr() (common.Address, error) {
	return _Contracts.Contract.PunishContractAddr(&_Contracts.CallOpts)
}

// StakingLockPeriod is a free data retrieval call binding the contract method 0xdb78dd28.
//
// Solidity: function StakingLockPeriod() view returns(uint64)
func (_Contracts *ContractsCaller) StakingLockPeriod(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "StakingLockPeriod")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// StakingLockPeriod is a free data retrieval call binding the contract method 0xdb78dd28.
//
// Solidity: function StakingLockPeriod() view returns(uint64)
func (_Contracts *ContractsSession) StakingLockPeriod() (uint64, error) {
	return _Contracts.Contract.StakingLockPeriod(&_Contracts.CallOpts)
}

// StakingLockPeriod is a free data retrieval call binding the contract method 0xdb78dd28.
//
// Solidity: function StakingLockPeriod() view returns(uint64)
func (_Contracts *ContractsCallerSession) StakingLockPeriod() (uint64, error) {
	return _Contracts.Contract.StakingLockPeriod(&_Contracts.CallOpts)
}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Contracts *ContractsCaller) ValidatorContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ValidatorContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Contracts *ContractsSession) ValidatorContractAddr() (common.Address, error) {
	return _Contracts.Contract.ValidatorContractAddr(&_Contracts.CallOpts)
}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Contracts *ContractsCallerSession) ValidatorContractAddr() (common.Address, error) {
	return _Contracts.Contract.ValidatorContractAddr(&_Contracts.CallOpts)
}

// WithdrawProfitPeriod is a free data retrieval call binding the contract method 0xefd8d8e2.
//
// Solidity: function WithdrawProfitPeriod() view returns(uint64)
func (_Contracts *ContractsCaller) WithdrawProfitPeriod(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "WithdrawProfitPeriod")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WithdrawProfitPeriod is a free data retrieval call binding the contract method 0xefd8d8e2.
//
// Solidity: function WithdrawProfitPeriod() view returns(uint64)
func (_Contracts *ContractsSession) WithdrawProfitPeriod() (uint64, error) {
	return _Contracts.Contract.WithdrawProfitPeriod(&_Contracts.CallOpts)
}

// WithdrawProfitPeriod is a free data retrieval call binding the contract method 0xefd8d8e2.
//
// Solidity: function WithdrawProfitPeriod() view returns(uint64)
func (_Contracts *ContractsCallerSession) WithdrawProfitPeriod() (uint64, error) {
	return _Contracts.Contract.WithdrawProfitPeriod(&_Contracts.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address)
func (_Contracts *ContractsCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "currentValidatorSet", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address)
func (_Contracts *ContractsSession) CurrentValidatorSet(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.CurrentValidatorSet(&_Contracts.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address)
func (_Contracts *ContractsCallerSession) CurrentValidatorSet(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.CurrentValidatorSet(&_Contracts.CallOpts, arg0)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_Contracts *ContractsCaller) GetActiveValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getActiveValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_Contracts *ContractsSession) GetActiveValidators() ([]common.Address, error) {
	return _Contracts.Contract.GetActiveValidators(&_Contracts.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_Contracts *ContractsCallerSession) GetActiveValidators() ([]common.Address, error) {
	return _Contracts.Contract.GetActiveValidators(&_Contracts.CallOpts)
}

// GetStakingInfo is a free data retrieval call binding the contract method 0x7f4f95fa.
//
// Solidity: function getStakingInfo(address staker, address val) view returns(uint256, uint256, uint256)
func (_Contracts *ContractsCaller) GetStakingInfo(opts *bind.CallOpts, staker common.Address, val common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getStakingInfo", staker, val)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetStakingInfo is a free data retrieval call binding the contract method 0x7f4f95fa.
//
// Solidity: function getStakingInfo(address staker, address val) view returns(uint256, uint256, uint256)
func (_Contracts *ContractsSession) GetStakingInfo(staker common.Address, val common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Contracts.Contract.GetStakingInfo(&_Contracts.CallOpts, staker, val)
}

// GetStakingInfo is a free data retrieval call binding the contract method 0x7f4f95fa.
//
// Solidity: function getStakingInfo(address staker, address val) view returns(uint256, uint256, uint256)
func (_Contracts *ContractsCallerSession) GetStakingInfo(staker common.Address, val common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Contracts.Contract.GetStakingInfo(&_Contracts.CallOpts, staker, val)
}

// GetTopValidators is a free data retrieval call binding the contract method 0xafeea115.
//
// Solidity: function getTopValidators() view returns(address[])
func (_Contracts *ContractsCaller) GetTopValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getTopValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTopValidators is a free data retrieval call binding the contract method 0xafeea115.
//
// Solidity: function getTopValidators() view returns(address[])
func (_Contracts *ContractsSession) GetTopValidators() ([]common.Address, error) {
	return _Contracts.Contract.GetTopValidators(&_Contracts.CallOpts)
}

// GetTopValidators is a free data retrieval call binding the contract method 0xafeea115.
//
// Solidity: function getTopValidators() view returns(address[])
func (_Contracts *ContractsCallerSession) GetTopValidators() ([]common.Address, error) {
	return _Contracts.Contract.GetTopValidators(&_Contracts.CallOpts)
}

// GetTotalStakeOfActiveValidators is a free data retrieval call binding the contract method 0xc253c384.
//
// Solidity: function getTotalStakeOfActiveValidators() view returns(uint256 total, uint256 len)
func (_Contracts *ContractsCaller) GetTotalStakeOfActiveValidators(opts *bind.CallOpts) (struct {
	Total *big.Int
	Len   *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getTotalStakeOfActiveValidators")

	outstruct := new(struct {
		Total *big.Int
		Len   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Len = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTotalStakeOfActiveValidators is a free data retrieval call binding the contract method 0xc253c384.
//
// Solidity: function getTotalStakeOfActiveValidators() view returns(uint256 total, uint256 len)
func (_Contracts *ContractsSession) GetTotalStakeOfActiveValidators() (struct {
	Total *big.Int
	Len   *big.Int
}, error) {
	return _Contracts.Contract.GetTotalStakeOfActiveValidators(&_Contracts.CallOpts)
}

// GetTotalStakeOfActiveValidators is a free data retrieval call binding the contract method 0xc253c384.
//
// Solidity: function getTotalStakeOfActiveValidators() view returns(uint256 total, uint256 len)
func (_Contracts *ContractsCallerSession) GetTotalStakeOfActiveValidators() (struct {
	Total *big.Int
	Len   *big.Int
}, error) {
	return _Contracts.Contract.GetTotalStakeOfActiveValidators(&_Contracts.CallOpts)
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address val) view returns(string, string, string, string, string)
func (_Contracts *ContractsCaller) GetValidatorDescription(opts *bind.CallOpts, val common.Address) (string, string, string, string, string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getValidatorDescription", val)

	if err != nil {
		return *new(string), *new(string), *new(string), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)

	return out0, out1, out2, out3, out4, err

}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address val) view returns(string, string, string, string, string)
func (_Contracts *ContractsSession) GetValidatorDescription(val common.Address) (string, string, string, string, string, error) {
	return _Contracts.Contract.GetValidatorDescription(&_Contracts.CallOpts, val)
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address val) view returns(string, string, string, string, string)
func (_Contracts *ContractsCallerSession) GetValidatorDescription(val common.Address) (string, string, string, string, string, error) {
	return _Contracts.Contract.GetValidatorDescription(&_Contracts.CallOpts, val)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address val) view returns(address, uint8, uint256, uint256, uint256, uint256, address[])
func (_Contracts *ContractsCaller) GetValidatorInfo(opts *bind.CallOpts, val common.Address) (common.Address, uint8, *big.Int, *big.Int, *big.Int, *big.Int, []common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getValidatorInfo", val)

	if err != nil {
		return *new(common.Address), *new(uint8), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new([]common.Address)).(*[]common.Address)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address val) view returns(address, uint8, uint256, uint256, uint256, uint256, address[])
func (_Contracts *ContractsSession) GetValidatorInfo(val common.Address) (common.Address, uint8, *big.Int, *big.Int, *big.Int, *big.Int, []common.Address, error) {
	return _Contracts.Contract.GetValidatorInfo(&_Contracts.CallOpts, val)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address val) view returns(address, uint8, uint256, uint256, uint256, uint256, address[])
func (_Contracts *ContractsCallerSession) GetValidatorInfo(val common.Address) (common.Address, uint8, *big.Int, *big.Int, *big.Int, *big.Int, []common.Address, error) {
	return _Contracts.Contract.GetValidatorInfo(&_Contracts.CallOpts, val)
}

// HighestValidatorsSet is a free data retrieval call binding the contract method 0x4b3d500b.
//
// Solidity: function highestValidatorsSet(uint256 ) view returns(address)
func (_Contracts *ContractsCaller) HighestValidatorsSet(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "highestValidatorsSet", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestValidatorsSet is a free data retrieval call binding the contract method 0x4b3d500b.
//
// Solidity: function highestValidatorsSet(uint256 ) view returns(address)
func (_Contracts *ContractsSession) HighestValidatorsSet(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.HighestValidatorsSet(&_Contracts.CallOpts, arg0)
}

// HighestValidatorsSet is a free data retrieval call binding the contract method 0x4b3d500b.
//
// Solidity: function highestValidatorsSet(uint256 ) view returns(address)
func (_Contracts *ContractsCallerSession) HighestValidatorsSet(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.HighestValidatorsSet(&_Contracts.CallOpts, arg0)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contracts *ContractsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contracts *ContractsSession) Initialized() (bool, error) {
	return _Contracts.Contract.Initialized(&_Contracts.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contracts *ContractsCallerSession) Initialized() (bool, error) {
	return _Contracts.Contract.Initialized(&_Contracts.CallOpts)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address who) view returns(bool)
func (_Contracts *ContractsCaller) IsActiveValidator(opts *bind.CallOpts, who common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isActiveValidator", who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address who) view returns(bool)
func (_Contracts *ContractsSession) IsActiveValidator(who common.Address) (bool, error) {
	return _Contracts.Contract.IsActiveValidator(&_Contracts.CallOpts, who)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address who) view returns(bool)
func (_Contracts *ContractsCallerSession) IsActiveValidator(who common.Address) (bool, error) {
	return _Contracts.Contract.IsActiveValidator(&_Contracts.CallOpts, who)
}

// IsTopValidator is a free data retrieval call binding the contract method 0x98e3b626.
//
// Solidity: function isTopValidator(address who) view returns(bool)
func (_Contracts *ContractsCaller) IsTopValidator(opts *bind.CallOpts, who common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isTopValidator", who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTopValidator is a free data retrieval call binding the contract method 0x98e3b626.
//
// Solidity: function isTopValidator(address who) view returns(bool)
func (_Contracts *ContractsSession) IsTopValidator(who common.Address) (bool, error) {
	return _Contracts.Contract.IsTopValidator(&_Contracts.CallOpts, who)
}

// IsTopValidator is a free data retrieval call binding the contract method 0x98e3b626.
//
// Solidity: function isTopValidator(address who) view returns(bool)
func (_Contracts *ContractsCallerSession) IsTopValidator(who common.Address) (bool, error) {
	return _Contracts.Contract.IsTopValidator(&_Contracts.CallOpts, who)
}

// TotalJailed is a free data retrieval call binding the contract method 0x534dacd6.
//
// Solidity: function totalJailed() view returns(uint256)
func (_Contracts *ContractsCaller) TotalJailed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalJailed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalJailed is a free data retrieval call binding the contract method 0x534dacd6.
//
// Solidity: function totalJailed() view returns(uint256)
func (_Contracts *ContractsSession) TotalJailed() (*big.Int, error) {
	return _Contracts.Contract.TotalJailed(&_Contracts.CallOpts)
}

// TotalJailed is a free data retrieval call binding the contract method 0x534dacd6.
//
// Solidity: function totalJailed() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalJailed() (*big.Int, error) {
	return _Contracts.Contract.TotalJailed(&_Contracts.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Contracts *ContractsCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Contracts *ContractsSession) TotalStake() (*big.Int, error) {
	return _Contracts.Contract.TotalStake(&_Contracts.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalStake() (*big.Int, error) {
	return _Contracts.Contract.TotalStake(&_Contracts.CallOpts)
}

// ValidateDescription is a free data retrieval call binding the contract method 0xb6c88519.
//
// Solidity: function validateDescription(string moniker, string identity, string website, string email, string details) pure returns(bool)
func (_Contracts *ContractsCaller) ValidateDescription(opts *bind.CallOpts, moniker string, identity string, website string, email string, details string) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "validateDescription", moniker, identity, website, email, details)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateDescription is a free data retrieval call binding the contract method 0xb6c88519.
//
// Solidity: function validateDescription(string moniker, string identity, string website, string email, string details) pure returns(bool)
func (_Contracts *ContractsSession) ValidateDescription(moniker string, identity string, website string, email string, details string) (bool, error) {
	return _Contracts.Contract.ValidateDescription(&_Contracts.CallOpts, moniker, identity, website, email, details)
}

// ValidateDescription is a free data retrieval call binding the contract method 0xb6c88519.
//
// Solidity: function validateDescription(string moniker, string identity, string website, string email, string details) pure returns(bool)
func (_Contracts *ContractsCallerSession) ValidateDescription(moniker string, identity string, website string, email string, details string) (bool, error) {
	return _Contracts.Contract.ValidateDescription(&_Contracts.CallOpts, moniker, identity, website, email, details)
}

// CreateOrEditValidator is a paid mutator transaction binding the contract method 0xa406fcb7.
//
// Solidity: function createOrEditValidator(address feeAddr, string moniker, string identity, string website, string email, string details) returns(bool)
func (_Contracts *ContractsTransactor) CreateOrEditValidator(opts *bind.TransactOpts, feeAddr common.Address, moniker string, identity string, website string, email string, details string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createOrEditValidator", feeAddr, moniker, identity, website, email, details)
}

// CreateOrEditValidator is a paid mutator transaction binding the contract method 0xa406fcb7.
//
// Solidity: function createOrEditValidator(address feeAddr, string moniker, string identity, string website, string email, string details) returns(bool)
func (_Contracts *ContractsSession) CreateOrEditValidator(feeAddr common.Address, moniker string, identity string, website string, email string, details string) (*types.Transaction, error) {
	return _Contracts.Contract.CreateOrEditValidator(&_Contracts.TransactOpts, feeAddr, moniker, identity, website, email, details)
}

// CreateOrEditValidator is a paid mutator transaction binding the contract method 0xa406fcb7.
//
// Solidity: function createOrEditValidator(address feeAddr, string moniker, string identity, string website, string email, string details) returns(bool)
func (_Contracts *ContractsTransactorSession) CreateOrEditValidator(feeAddr common.Address, moniker string, identity string, website string, email string, details string) (*types.Transaction, error) {
	return _Contracts.Contract.CreateOrEditValidator(&_Contracts.TransactOpts, feeAddr, moniker, identity, website, email, details)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0xd6c0edad.
//
// Solidity: function distributeBlockReward() payable returns()
func (_Contracts *ContractsTransactor) DistributeBlockReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "distributeBlockReward")
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0xd6c0edad.
//
// Solidity: function distributeBlockReward() payable returns()
func (_Contracts *ContractsSession) DistributeBlockReward() (*types.Transaction, error) {
	return _Contracts.Contract.DistributeBlockReward(&_Contracts.TransactOpts)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0xd6c0edad.
//
// Solidity: function distributeBlockReward() payable returns()
func (_Contracts *ContractsTransactorSession) DistributeBlockReward() (*types.Transaction, error) {
	return _Contracts.Contract.DistributeBlockReward(&_Contracts.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] vals) returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts, vals []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize", vals)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] vals) returns()
func (_Contracts *ContractsSession) Initialize(vals []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, vals)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] vals) returns()
func (_Contracts *ContractsTransactorSession) Initialize(vals []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, vals)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address val) returns()
func (_Contracts *ContractsTransactor) RemoveValidator(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeValidator", val)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address val) returns()
func (_Contracts *ContractsSession) RemoveValidator(val common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveValidator(&_Contracts.TransactOpts, val)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address val) returns()
func (_Contracts *ContractsTransactorSession) RemoveValidator(val common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveValidator(&_Contracts.TransactOpts, val)
}

// RemoveValidatorIncoming is a paid mutator transaction binding the contract method 0x5dd09590.
//
// Solidity: function removeValidatorIncoming(address val) returns()
func (_Contracts *ContractsTransactor) RemoveValidatorIncoming(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeValidatorIncoming", val)
}

// RemoveValidatorIncoming is a paid mutator transaction binding the contract method 0x5dd09590.
//
// Solidity: function removeValidatorIncoming(address val) returns()
func (_Contracts *ContractsSession) RemoveValidatorIncoming(val common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveValidatorIncoming(&_Contracts.TransactOpts, val)
}

// RemoveValidatorIncoming is a paid mutator transaction binding the contract method 0x5dd09590.
//
// Solidity: function removeValidatorIncoming(address val) returns()
func (_Contracts *ContractsTransactorSession) RemoveValidatorIncoming(val common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveValidatorIncoming(&_Contracts.TransactOpts, val)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address validator) payable returns(bool)
func (_Contracts *ContractsTransactor) Stake(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "stake", validator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address validator) payable returns(bool)
func (_Contracts *ContractsSession) Stake(validator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Stake(&_Contracts.TransactOpts, validator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address validator) payable returns(bool)
func (_Contracts *ContractsTransactorSession) Stake(validator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Stake(&_Contracts.TransactOpts, validator)
}

// TryReactive is a paid mutator transaction binding the contract method 0x82bd3d92.
//
// Solidity: function tryReactive(address validator) returns(bool)
func (_Contracts *ContractsTransactor) TryReactive(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "tryReactive", validator)
}

// TryReactive is a paid mutator transaction binding the contract method 0x82bd3d92.
//
// Solidity: function tryReactive(address validator) returns(bool)
func (_Contracts *ContractsSession) TryReactive(validator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TryReactive(&_Contracts.TransactOpts, validator)
}

// TryReactive is a paid mutator transaction binding the contract method 0x82bd3d92.
//
// Solidity: function tryReactive(address validator) returns(bool)
func (_Contracts *ContractsTransactorSession) TryReactive(validator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TryReactive(&_Contracts.TransactOpts, validator)
}

// UnStake is a paid mutator transaction binding the contract method 0x387e4948.
//
// Solidity: function unStake(address validator) returns(bool)
func (_Contracts *ContractsTransactor) UnStake(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unStake", validator)
}

// UnStake is a paid mutator transaction binding the contract method 0x387e4948.
//
// Solidity: function unStake(address validator) returns(bool)
func (_Contracts *ContractsSession) UnStake(validator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnStake(&_Contracts.TransactOpts, validator)
}

// UnStake is a paid mutator transaction binding the contract method 0x387e4948.
//
// Solidity: function unStake(address validator) returns(bool)
func (_Contracts *ContractsTransactorSession) UnStake(validator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnStake(&_Contracts.TransactOpts, validator)
}

// UpdateActiveValidatorSet is a paid mutator transaction binding the contract method 0x056dd786.
//
// Solidity: function updateActiveValidatorSet(uint256 epoch) returns()
func (_Contracts *ContractsTransactor) UpdateActiveValidatorSet(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateActiveValidatorSet", epoch)
}

// UpdateActiveValidatorSet is a paid mutator transaction binding the contract method 0x056dd786.
//
// Solidity: function updateActiveValidatorSet(uint256 epoch) returns()
func (_Contracts *ContractsSession) UpdateActiveValidatorSet(epoch *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateActiveValidatorSet(&_Contracts.TransactOpts, epoch)
}

// UpdateActiveValidatorSet is a paid mutator transaction binding the contract method 0x056dd786.
//
// Solidity: function updateActiveValidatorSet(uint256 epoch) returns()
func (_Contracts *ContractsTransactorSession) UpdateActiveValidatorSet(epoch *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateActiveValidatorSet(&_Contracts.TransactOpts, epoch)
}

// WithdrawProfits is a paid mutator transaction binding the contract method 0x00362a77.
//
// Solidity: function withdrawProfits(address validator) returns(bool)
func (_Contracts *ContractsTransactor) WithdrawProfits(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawProfits", validator)
}

// WithdrawProfits is a paid mutator transaction binding the contract method 0x00362a77.
//
// Solidity: function withdrawProfits(address validator) returns(bool)
func (_Contracts *ContractsSession) WithdrawProfits(validator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawProfits(&_Contracts.TransactOpts, validator)
}

// WithdrawProfits is a paid mutator transaction binding the contract method 0x00362a77.
//
// Solidity: function withdrawProfits(address validator) returns(bool)
func (_Contracts *ContractsTransactorSession) WithdrawProfits(validator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawProfits(&_Contracts.TransactOpts, validator)
}

// WithdrawStaking is a paid mutator transaction binding the contract method 0x222d3b05.
//
// Solidity: function withdrawStaking(address validator) returns(bool)
func (_Contracts *ContractsTransactor) WithdrawStaking(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawStaking", validator)
}

// WithdrawStaking is a paid mutator transaction binding the contract method 0x222d3b05.
//
// Solidity: function withdrawStaking(address validator) returns(bool)
func (_Contracts *ContractsSession) WithdrawStaking(validator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawStaking(&_Contracts.TransactOpts, validator)
}

// WithdrawStaking is a paid mutator transaction binding the contract method 0x222d3b05.
//
// Solidity: function withdrawStaking(address validator) returns(bool)
func (_Contracts *ContractsTransactorSession) WithdrawStaking(validator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawStaking(&_Contracts.TransactOpts, validator)
}

// ContractsLogAddToTopValidatorsIterator is returned from FilterLogAddToTopValidators and is used to iterate over the raw logs and unpacked data for LogAddToTopValidators events raised by the Contracts contract.
type ContractsLogAddToTopValidatorsIterator struct {
	Event *ContractsLogAddToTopValidators // Event containing the contract specifics and raw log

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
func (it *ContractsLogAddToTopValidatorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogAddToTopValidators)
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
		it.Event = new(ContractsLogAddToTopValidators)
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
func (it *ContractsLogAddToTopValidatorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogAddToTopValidatorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogAddToTopValidators represents a LogAddToTopValidators event raised by the Contracts contract.
type ContractsLogAddToTopValidators struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddToTopValidators is a free log retrieval operation binding the contract event 0x1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a.
//
// Solidity: event LogAddToTopValidators(address indexed val, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogAddToTopValidators(opts *bind.FilterOpts, val []common.Address) (*ContractsLogAddToTopValidatorsIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogAddToTopValidators", valRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogAddToTopValidatorsIterator{contract: _Contracts.contract, event: "LogAddToTopValidators", logs: logs, sub: sub}, nil
}

// WatchLogAddToTopValidators is a free log subscription operation binding the contract event 0x1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a.
//
// Solidity: event LogAddToTopValidators(address indexed val, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogAddToTopValidators(opts *bind.WatchOpts, sink chan<- *ContractsLogAddToTopValidators, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogAddToTopValidators", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogAddToTopValidators)
				if err := _Contracts.contract.UnpackLog(event, "LogAddToTopValidators", log); err != nil {
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

// ParseLogAddToTopValidators is a log parse operation binding the contract event 0x1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a.
//
// Solidity: event LogAddToTopValidators(address indexed val, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogAddToTopValidators(log types.Log) (*ContractsLogAddToTopValidators, error) {
	event := new(ContractsLogAddToTopValidators)
	if err := _Contracts.contract.UnpackLog(event, "LogAddToTopValidators", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogCreateValidatorIterator is returned from FilterLogCreateValidator and is used to iterate over the raw logs and unpacked data for LogCreateValidator events raised by the Contracts contract.
type ContractsLogCreateValidatorIterator struct {
	Event *ContractsLogCreateValidator // Event containing the contract specifics and raw log

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
func (it *ContractsLogCreateValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogCreateValidator)
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
		it.Event = new(ContractsLogCreateValidator)
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
func (it *ContractsLogCreateValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogCreateValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogCreateValidator represents a LogCreateValidator event raised by the Contracts contract.
type ContractsLogCreateValidator struct {
	Val  common.Address
	Fee  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogCreateValidator is a free log retrieval operation binding the contract event 0x887eec9d757b7247dd8e51198f9d1b8f27979bceb34bdcc1bffd4ec5ec736c22.
//
// Solidity: event LogCreateValidator(address indexed val, address indexed fee, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogCreateValidator(opts *bind.FilterOpts, val []common.Address, fee []common.Address) (*ContractsLogCreateValidatorIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogCreateValidator", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogCreateValidatorIterator{contract: _Contracts.contract, event: "LogCreateValidator", logs: logs, sub: sub}, nil
}

// WatchLogCreateValidator is a free log subscription operation binding the contract event 0x887eec9d757b7247dd8e51198f9d1b8f27979bceb34bdcc1bffd4ec5ec736c22.
//
// Solidity: event LogCreateValidator(address indexed val, address indexed fee, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogCreateValidator(opts *bind.WatchOpts, sink chan<- *ContractsLogCreateValidator, val []common.Address, fee []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogCreateValidator", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogCreateValidator)
				if err := _Contracts.contract.UnpackLog(event, "LogCreateValidator", log); err != nil {
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

// ParseLogCreateValidator is a log parse operation binding the contract event 0x887eec9d757b7247dd8e51198f9d1b8f27979bceb34bdcc1bffd4ec5ec736c22.
//
// Solidity: event LogCreateValidator(address indexed val, address indexed fee, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogCreateValidator(log types.Log) (*ContractsLogCreateValidator, error) {
	event := new(ContractsLogCreateValidator)
	if err := _Contracts.contract.UnpackLog(event, "LogCreateValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogDistributeBlockRewardIterator is returned from FilterLogDistributeBlockReward and is used to iterate over the raw logs and unpacked data for LogDistributeBlockReward events raised by the Contracts contract.
type ContractsLogDistributeBlockRewardIterator struct {
	Event *ContractsLogDistributeBlockReward // Event containing the contract specifics and raw log

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
func (it *ContractsLogDistributeBlockRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogDistributeBlockReward)
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
		it.Event = new(ContractsLogDistributeBlockReward)
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
func (it *ContractsLogDistributeBlockRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogDistributeBlockRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogDistributeBlockReward represents a LogDistributeBlockReward event raised by the Contracts contract.
type ContractsLogDistributeBlockReward struct {
	Coinbase    common.Address
	BlockReward *big.Int
	Time        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogDistributeBlockReward is a free log retrieval operation binding the contract event 0x7dc4e5df59513708dca355b8706273a5df7b810a4cec8019f2a4b9bb166a1a04.
//
// Solidity: event LogDistributeBlockReward(address indexed coinbase, uint256 blockReward, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogDistributeBlockReward(opts *bind.FilterOpts, coinbase []common.Address) (*ContractsLogDistributeBlockRewardIterator, error) {

	var coinbaseRule []interface{}
	for _, coinbaseItem := range coinbase {
		coinbaseRule = append(coinbaseRule, coinbaseItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogDistributeBlockReward", coinbaseRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogDistributeBlockRewardIterator{contract: _Contracts.contract, event: "LogDistributeBlockReward", logs: logs, sub: sub}, nil
}

// WatchLogDistributeBlockReward is a free log subscription operation binding the contract event 0x7dc4e5df59513708dca355b8706273a5df7b810a4cec8019f2a4b9bb166a1a04.
//
// Solidity: event LogDistributeBlockReward(address indexed coinbase, uint256 blockReward, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogDistributeBlockReward(opts *bind.WatchOpts, sink chan<- *ContractsLogDistributeBlockReward, coinbase []common.Address) (event.Subscription, error) {

	var coinbaseRule []interface{}
	for _, coinbaseItem := range coinbase {
		coinbaseRule = append(coinbaseRule, coinbaseItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogDistributeBlockReward", coinbaseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogDistributeBlockReward)
				if err := _Contracts.contract.UnpackLog(event, "LogDistributeBlockReward", log); err != nil {
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

// ParseLogDistributeBlockReward is a log parse operation binding the contract event 0x7dc4e5df59513708dca355b8706273a5df7b810a4cec8019f2a4b9bb166a1a04.
//
// Solidity: event LogDistributeBlockReward(address indexed coinbase, uint256 blockReward, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogDistributeBlockReward(log types.Log) (*ContractsLogDistributeBlockReward, error) {
	event := new(ContractsLogDistributeBlockReward)
	if err := _Contracts.contract.UnpackLog(event, "LogDistributeBlockReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogEditValidatorIterator is returned from FilterLogEditValidator and is used to iterate over the raw logs and unpacked data for LogEditValidator events raised by the Contracts contract.
type ContractsLogEditValidatorIterator struct {
	Event *ContractsLogEditValidator // Event containing the contract specifics and raw log

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
func (it *ContractsLogEditValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogEditValidator)
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
		it.Event = new(ContractsLogEditValidator)
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
func (it *ContractsLogEditValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogEditValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogEditValidator represents a LogEditValidator event raised by the Contracts contract.
type ContractsLogEditValidator struct {
	Val  common.Address
	Fee  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogEditValidator is a free log retrieval operation binding the contract event 0xb8421f65501371f54d58de1937ff1e1ccdb76423ef6f84acea1814a0f6362ca0.
//
// Solidity: event LogEditValidator(address indexed val, address indexed fee, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogEditValidator(opts *bind.FilterOpts, val []common.Address, fee []common.Address) (*ContractsLogEditValidatorIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogEditValidator", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogEditValidatorIterator{contract: _Contracts.contract, event: "LogEditValidator", logs: logs, sub: sub}, nil
}

// WatchLogEditValidator is a free log subscription operation binding the contract event 0xb8421f65501371f54d58de1937ff1e1ccdb76423ef6f84acea1814a0f6362ca0.
//
// Solidity: event LogEditValidator(address indexed val, address indexed fee, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogEditValidator(opts *bind.WatchOpts, sink chan<- *ContractsLogEditValidator, val []common.Address, fee []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogEditValidator", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogEditValidator)
				if err := _Contracts.contract.UnpackLog(event, "LogEditValidator", log); err != nil {
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

// ParseLogEditValidator is a log parse operation binding the contract event 0xb8421f65501371f54d58de1937ff1e1ccdb76423ef6f84acea1814a0f6362ca0.
//
// Solidity: event LogEditValidator(address indexed val, address indexed fee, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogEditValidator(log types.Log) (*ContractsLogEditValidator, error) {
	event := new(ContractsLogEditValidator)
	if err := _Contracts.contract.UnpackLog(event, "LogEditValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogReactiveIterator is returned from FilterLogReactive and is used to iterate over the raw logs and unpacked data for LogReactive events raised by the Contracts contract.
type ContractsLogReactiveIterator struct {
	Event *ContractsLogReactive // Event containing the contract specifics and raw log

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
func (it *ContractsLogReactiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogReactive)
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
		it.Event = new(ContractsLogReactive)
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
func (it *ContractsLogReactiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogReactiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogReactive represents a LogReactive event raised by the Contracts contract.
type ContractsLogReactive struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogReactive is a free log retrieval operation binding the contract event 0xd8b2c426ec1be69ca7583d26b1e893946e3227430d3ebc3bd64d9e1c469cb400.
//
// Solidity: event LogReactive(address indexed val, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogReactive(opts *bind.FilterOpts, val []common.Address) (*ContractsLogReactiveIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogReactive", valRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogReactiveIterator{contract: _Contracts.contract, event: "LogReactive", logs: logs, sub: sub}, nil
}

// WatchLogReactive is a free log subscription operation binding the contract event 0xd8b2c426ec1be69ca7583d26b1e893946e3227430d3ebc3bd64d9e1c469cb400.
//
// Solidity: event LogReactive(address indexed val, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogReactive(opts *bind.WatchOpts, sink chan<- *ContractsLogReactive, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogReactive", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogReactive)
				if err := _Contracts.contract.UnpackLog(event, "LogReactive", log); err != nil {
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

// ParseLogReactive is a log parse operation binding the contract event 0xd8b2c426ec1be69ca7583d26b1e893946e3227430d3ebc3bd64d9e1c469cb400.
//
// Solidity: event LogReactive(address indexed val, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogReactive(log types.Log) (*ContractsLogReactive, error) {
	event := new(ContractsLogReactive)
	if err := _Contracts.contract.UnpackLog(event, "LogReactive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogRemoveFromTopValidatorsIterator is returned from FilterLogRemoveFromTopValidators and is used to iterate over the raw logs and unpacked data for LogRemoveFromTopValidators events raised by the Contracts contract.
type ContractsLogRemoveFromTopValidatorsIterator struct {
	Event *ContractsLogRemoveFromTopValidators // Event containing the contract specifics and raw log

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
func (it *ContractsLogRemoveFromTopValidatorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogRemoveFromTopValidators)
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
		it.Event = new(ContractsLogRemoveFromTopValidators)
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
func (it *ContractsLogRemoveFromTopValidatorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogRemoveFromTopValidatorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogRemoveFromTopValidators represents a LogRemoveFromTopValidators event raised by the Contracts contract.
type ContractsLogRemoveFromTopValidators struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRemoveFromTopValidators is a free log retrieval operation binding the contract event 0x7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed.
//
// Solidity: event LogRemoveFromTopValidators(address indexed val, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogRemoveFromTopValidators(opts *bind.FilterOpts, val []common.Address) (*ContractsLogRemoveFromTopValidatorsIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogRemoveFromTopValidators", valRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogRemoveFromTopValidatorsIterator{contract: _Contracts.contract, event: "LogRemoveFromTopValidators", logs: logs, sub: sub}, nil
}

// WatchLogRemoveFromTopValidators is a free log subscription operation binding the contract event 0x7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed.
//
// Solidity: event LogRemoveFromTopValidators(address indexed val, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogRemoveFromTopValidators(opts *bind.WatchOpts, sink chan<- *ContractsLogRemoveFromTopValidators, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogRemoveFromTopValidators", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogRemoveFromTopValidators)
				if err := _Contracts.contract.UnpackLog(event, "LogRemoveFromTopValidators", log); err != nil {
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

// ParseLogRemoveFromTopValidators is a log parse operation binding the contract event 0x7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed.
//
// Solidity: event LogRemoveFromTopValidators(address indexed val, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogRemoveFromTopValidators(log types.Log) (*ContractsLogRemoveFromTopValidators, error) {
	event := new(ContractsLogRemoveFromTopValidators)
	if err := _Contracts.contract.UnpackLog(event, "LogRemoveFromTopValidators", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogRemoveValidatorIterator is returned from FilterLogRemoveValidator and is used to iterate over the raw logs and unpacked data for LogRemoveValidator events raised by the Contracts contract.
type ContractsLogRemoveValidatorIterator struct {
	Event *ContractsLogRemoveValidator // Event containing the contract specifics and raw log

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
func (it *ContractsLogRemoveValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogRemoveValidator)
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
		it.Event = new(ContractsLogRemoveValidator)
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
func (it *ContractsLogRemoveValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogRemoveValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogRemoveValidator represents a LogRemoveValidator event raised by the Contracts contract.
type ContractsLogRemoveValidator struct {
	Val  common.Address
	Amt  *big.Int
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRemoveValidator is a free log retrieval operation binding the contract event 0xa26de7ab324eac08c596549f421e5c8741213d237d2e9a2c9c0ebde0a7a849fe.
//
// Solidity: event LogRemoveValidator(address indexed val, uint256 amt, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogRemoveValidator(opts *bind.FilterOpts, val []common.Address) (*ContractsLogRemoveValidatorIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogRemoveValidator", valRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogRemoveValidatorIterator{contract: _Contracts.contract, event: "LogRemoveValidator", logs: logs, sub: sub}, nil
}

// WatchLogRemoveValidator is a free log subscription operation binding the contract event 0xa26de7ab324eac08c596549f421e5c8741213d237d2e9a2c9c0ebde0a7a849fe.
//
// Solidity: event LogRemoveValidator(address indexed val, uint256 amt, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogRemoveValidator(opts *bind.WatchOpts, sink chan<- *ContractsLogRemoveValidator, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogRemoveValidator", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogRemoveValidator)
				if err := _Contracts.contract.UnpackLog(event, "LogRemoveValidator", log); err != nil {
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

// ParseLogRemoveValidator is a log parse operation binding the contract event 0xa26de7ab324eac08c596549f421e5c8741213d237d2e9a2c9c0ebde0a7a849fe.
//
// Solidity: event LogRemoveValidator(address indexed val, uint256 amt, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogRemoveValidator(log types.Log) (*ContractsLogRemoveValidator, error) {
	event := new(ContractsLogRemoveValidator)
	if err := _Contracts.contract.UnpackLog(event, "LogRemoveValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogRemoveValidatorIncomingIterator is returned from FilterLogRemoveValidatorIncoming and is used to iterate over the raw logs and unpacked data for LogRemoveValidatorIncoming events raised by the Contracts contract.
type ContractsLogRemoveValidatorIncomingIterator struct {
	Event *ContractsLogRemoveValidatorIncoming // Event containing the contract specifics and raw log

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
func (it *ContractsLogRemoveValidatorIncomingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogRemoveValidatorIncoming)
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
		it.Event = new(ContractsLogRemoveValidatorIncoming)
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
func (it *ContractsLogRemoveValidatorIncomingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogRemoveValidatorIncomingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogRemoveValidatorIncoming represents a LogRemoveValidatorIncoming event raised by the Contracts contract.
type ContractsLogRemoveValidatorIncoming struct {
	Val  common.Address
	Amt  *big.Int
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRemoveValidatorIncoming is a free log retrieval operation binding the contract event 0xe294e9d73f8eee23e21b2e1567960625a6b5d339cb127b55d0d09473a9951235.
//
// Solidity: event LogRemoveValidatorIncoming(address indexed val, uint256 amt, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogRemoveValidatorIncoming(opts *bind.FilterOpts, val []common.Address) (*ContractsLogRemoveValidatorIncomingIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogRemoveValidatorIncoming", valRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogRemoveValidatorIncomingIterator{contract: _Contracts.contract, event: "LogRemoveValidatorIncoming", logs: logs, sub: sub}, nil
}

// WatchLogRemoveValidatorIncoming is a free log subscription operation binding the contract event 0xe294e9d73f8eee23e21b2e1567960625a6b5d339cb127b55d0d09473a9951235.
//
// Solidity: event LogRemoveValidatorIncoming(address indexed val, uint256 amt, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogRemoveValidatorIncoming(opts *bind.WatchOpts, sink chan<- *ContractsLogRemoveValidatorIncoming, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogRemoveValidatorIncoming", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogRemoveValidatorIncoming)
				if err := _Contracts.contract.UnpackLog(event, "LogRemoveValidatorIncoming", log); err != nil {
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

// ParseLogRemoveValidatorIncoming is a log parse operation binding the contract event 0xe294e9d73f8eee23e21b2e1567960625a6b5d339cb127b55d0d09473a9951235.
//
// Solidity: event LogRemoveValidatorIncoming(address indexed val, uint256 amt, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogRemoveValidatorIncoming(log types.Log) (*ContractsLogRemoveValidatorIncoming, error) {
	event := new(ContractsLogRemoveValidatorIncoming)
	if err := _Contracts.contract.UnpackLog(event, "LogRemoveValidatorIncoming", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogStakeIterator is returned from FilterLogStake and is used to iterate over the raw logs and unpacked data for LogStake events raised by the Contracts contract.
type ContractsLogStakeIterator struct {
	Event *ContractsLogStake // Event containing the contract specifics and raw log

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
func (it *ContractsLogStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogStake)
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
		it.Event = new(ContractsLogStake)
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
func (it *ContractsLogStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogStake represents a LogStake event raised by the Contracts contract.
type ContractsLogStake struct {
	Staker  common.Address
	Val     common.Address
	Staking *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogStake is a free log retrieval operation binding the contract event 0xb9ba725934532316cffe10975da6eb25ad49c2d1c294d982c46c9f8d684ee075.
//
// Solidity: event LogStake(address indexed staker, address indexed val, uint256 staking, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogStake(opts *bind.FilterOpts, staker []common.Address, val []common.Address) (*ContractsLogStakeIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogStake", stakerRule, valRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogStakeIterator{contract: _Contracts.contract, event: "LogStake", logs: logs, sub: sub}, nil
}

// WatchLogStake is a free log subscription operation binding the contract event 0xb9ba725934532316cffe10975da6eb25ad49c2d1c294d982c46c9f8d684ee075.
//
// Solidity: event LogStake(address indexed staker, address indexed val, uint256 staking, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogStake(opts *bind.WatchOpts, sink chan<- *ContractsLogStake, staker []common.Address, val []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogStake", stakerRule, valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogStake)
				if err := _Contracts.contract.UnpackLog(event, "LogStake", log); err != nil {
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

// ParseLogStake is a log parse operation binding the contract event 0xb9ba725934532316cffe10975da6eb25ad49c2d1c294d982c46c9f8d684ee075.
//
// Solidity: event LogStake(address indexed staker, address indexed val, uint256 staking, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogStake(log types.Log) (*ContractsLogStake, error) {
	event := new(ContractsLogStake)
	if err := _Contracts.contract.UnpackLog(event, "LogStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogUnStakeIterator is returned from FilterLogUnStake and is used to iterate over the raw logs and unpacked data for LogUnStake events raised by the Contracts contract.
type ContractsLogUnStakeIterator struct {
	Event *ContractsLogUnStake // Event containing the contract specifics and raw log

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
func (it *ContractsLogUnStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogUnStake)
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
		it.Event = new(ContractsLogUnStake)
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
func (it *ContractsLogUnStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogUnStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogUnStake represents a LogUnStake event raised by the Contracts contract.
type ContractsLogUnStake struct {
	Staker common.Address
	Val    common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogUnStake is a free log retrieval operation binding the contract event 0x1b1cb26a0e501aae1495d2137de6fad492b3c33d5f7dfa73db02e1e58ae6f3cf.
//
// Solidity: event LogUnStake(address indexed staker, address indexed val, uint256 amount, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogUnStake(opts *bind.FilterOpts, staker []common.Address, val []common.Address) (*ContractsLogUnStakeIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogUnStake", stakerRule, valRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogUnStakeIterator{contract: _Contracts.contract, event: "LogUnStake", logs: logs, sub: sub}, nil
}

// WatchLogUnStake is a free log subscription operation binding the contract event 0x1b1cb26a0e501aae1495d2137de6fad492b3c33d5f7dfa73db02e1e58ae6f3cf.
//
// Solidity: event LogUnStake(address indexed staker, address indexed val, uint256 amount, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogUnStake(opts *bind.WatchOpts, sink chan<- *ContractsLogUnStake, staker []common.Address, val []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogUnStake", stakerRule, valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogUnStake)
				if err := _Contracts.contract.UnpackLog(event, "LogUnStake", log); err != nil {
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

// ParseLogUnStake is a log parse operation binding the contract event 0x1b1cb26a0e501aae1495d2137de6fad492b3c33d5f7dfa73db02e1e58ae6f3cf.
//
// Solidity: event LogUnStake(address indexed staker, address indexed val, uint256 amount, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogUnStake(log types.Log) (*ContractsLogUnStake, error) {
	event := new(ContractsLogUnStake)
	if err := _Contracts.contract.UnpackLog(event, "LogUnStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogUpdateValidatorIterator is returned from FilterLogUpdateValidator and is used to iterate over the raw logs and unpacked data for LogUpdateValidator events raised by the Contracts contract.
type ContractsLogUpdateValidatorIterator struct {
	Event *ContractsLogUpdateValidator // Event containing the contract specifics and raw log

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
func (it *ContractsLogUpdateValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogUpdateValidator)
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
		it.Event = new(ContractsLogUpdateValidator)
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
func (it *ContractsLogUpdateValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogUpdateValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogUpdateValidator represents a LogUpdateValidator event raised by the Contracts contract.
type ContractsLogUpdateValidator struct {
	NewSet []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogUpdateValidator is a free log retrieval operation binding the contract event 0xeacea8f3c22f06c0b18306bdb04d0a967255129e8ce0094debb0a0ff89d006b5.
//
// Solidity: event LogUpdateValidator(address[] newSet)
func (_Contracts *ContractsFilterer) FilterLogUpdateValidator(opts *bind.FilterOpts) (*ContractsLogUpdateValidatorIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogUpdateValidator")
	if err != nil {
		return nil, err
	}
	return &ContractsLogUpdateValidatorIterator{contract: _Contracts.contract, event: "LogUpdateValidator", logs: logs, sub: sub}, nil
}

// WatchLogUpdateValidator is a free log subscription operation binding the contract event 0xeacea8f3c22f06c0b18306bdb04d0a967255129e8ce0094debb0a0ff89d006b5.
//
// Solidity: event LogUpdateValidator(address[] newSet)
func (_Contracts *ContractsFilterer) WatchLogUpdateValidator(opts *bind.WatchOpts, sink chan<- *ContractsLogUpdateValidator) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogUpdateValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogUpdateValidator)
				if err := _Contracts.contract.UnpackLog(event, "LogUpdateValidator", log); err != nil {
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

// ParseLogUpdateValidator is a log parse operation binding the contract event 0xeacea8f3c22f06c0b18306bdb04d0a967255129e8ce0094debb0a0ff89d006b5.
//
// Solidity: event LogUpdateValidator(address[] newSet)
func (_Contracts *ContractsFilterer) ParseLogUpdateValidator(log types.Log) (*ContractsLogUpdateValidator, error) {
	event := new(ContractsLogUpdateValidator)
	if err := _Contracts.contract.UnpackLog(event, "LogUpdateValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogWithdrawProfitsIterator is returned from FilterLogWithdrawProfits and is used to iterate over the raw logs and unpacked data for LogWithdrawProfits events raised by the Contracts contract.
type ContractsLogWithdrawProfitsIterator struct {
	Event *ContractsLogWithdrawProfits // Event containing the contract specifics and raw log

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
func (it *ContractsLogWithdrawProfitsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogWithdrawProfits)
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
		it.Event = new(ContractsLogWithdrawProfits)
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
func (it *ContractsLogWithdrawProfitsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogWithdrawProfitsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogWithdrawProfits represents a LogWithdrawProfits event raised by the Contracts contract.
type ContractsLogWithdrawProfits struct {
	Val  common.Address
	Fee  common.Address
	Amt  *big.Int
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawProfits is a free log retrieval operation binding the contract event 0x51a69b4502f660774c9339825c7b5adbf0b8622289134647e29728ec5d9b3bb9.
//
// Solidity: event LogWithdrawProfits(address indexed val, address indexed fee, uint256 amt, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogWithdrawProfits(opts *bind.FilterOpts, val []common.Address, fee []common.Address) (*ContractsLogWithdrawProfitsIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogWithdrawProfits", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogWithdrawProfitsIterator{contract: _Contracts.contract, event: "LogWithdrawProfits", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawProfits is a free log subscription operation binding the contract event 0x51a69b4502f660774c9339825c7b5adbf0b8622289134647e29728ec5d9b3bb9.
//
// Solidity: event LogWithdrawProfits(address indexed val, address indexed fee, uint256 amt, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogWithdrawProfits(opts *bind.WatchOpts, sink chan<- *ContractsLogWithdrawProfits, val []common.Address, fee []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogWithdrawProfits", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogWithdrawProfits)
				if err := _Contracts.contract.UnpackLog(event, "LogWithdrawProfits", log); err != nil {
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

// ParseLogWithdrawProfits is a log parse operation binding the contract event 0x51a69b4502f660774c9339825c7b5adbf0b8622289134647e29728ec5d9b3bb9.
//
// Solidity: event LogWithdrawProfits(address indexed val, address indexed fee, uint256 amt, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogWithdrawProfits(log types.Log) (*ContractsLogWithdrawProfits, error) {
	event := new(ContractsLogWithdrawProfits)
	if err := _Contracts.contract.UnpackLog(event, "LogWithdrawProfits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogWithdrawStakingIterator is returned from FilterLogWithdrawStaking and is used to iterate over the raw logs and unpacked data for LogWithdrawStaking events raised by the Contracts contract.
type ContractsLogWithdrawStakingIterator struct {
	Event *ContractsLogWithdrawStaking // Event containing the contract specifics and raw log

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
func (it *ContractsLogWithdrawStakingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogWithdrawStaking)
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
		it.Event = new(ContractsLogWithdrawStaking)
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
func (it *ContractsLogWithdrawStakingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogWithdrawStakingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogWithdrawStaking represents a LogWithdrawStaking event raised by the Contracts contract.
type ContractsLogWithdrawStaking struct {
	Staker common.Address
	Val    common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawStaking is a free log retrieval operation binding the contract event 0xa70cd94070cd852339a76b32cf2d95a3c8f2a322269163d276071c1c14955619.
//
// Solidity: event LogWithdrawStaking(address indexed staker, address indexed val, uint256 amount, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogWithdrawStaking(opts *bind.FilterOpts, staker []common.Address, val []common.Address) (*ContractsLogWithdrawStakingIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogWithdrawStaking", stakerRule, valRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogWithdrawStakingIterator{contract: _Contracts.contract, event: "LogWithdrawStaking", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawStaking is a free log subscription operation binding the contract event 0xa70cd94070cd852339a76b32cf2d95a3c8f2a322269163d276071c1c14955619.
//
// Solidity: event LogWithdrawStaking(address indexed staker, address indexed val, uint256 amount, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogWithdrawStaking(opts *bind.WatchOpts, sink chan<- *ContractsLogWithdrawStaking, staker []common.Address, val []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogWithdrawStaking", stakerRule, valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogWithdrawStaking)
				if err := _Contracts.contract.UnpackLog(event, "LogWithdrawStaking", log); err != nil {
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

// ParseLogWithdrawStaking is a log parse operation binding the contract event 0xa70cd94070cd852339a76b32cf2d95a3c8f2a322269163d276071c1c14955619.
//
// Solidity: event LogWithdrawStaking(address indexed staker, address indexed val, uint256 amount, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogWithdrawStaking(log types.Log) (*ContractsLogWithdrawStaking, error) {
	event := new(ContractsLogWithdrawStaking)
	if err := _Contracts.contract.UnpackLog(event, "LogWithdrawStaking", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
