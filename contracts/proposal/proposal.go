// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proposal

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
const ContractsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogCreateProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogPassProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogRejectProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogSetUnPassed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"auth\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogVote\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FeeRecoder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MaxValidators\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinimalStakingCoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ProposalContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PunishContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StakingLockPeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValidatorContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WithdrawProfitPeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vals\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pass\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalLastingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"agree\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reject\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"resultExist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"setUnPassed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"auth\",\"type\":\"bool\"}],\"name\":\"voteProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"voteTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auth\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// Pass is a free data retrieval call binding the contract method 0x82c4b3b2.
//
// Solidity: function pass(address ) view returns(bool)
func (_Contracts *ContractsCaller) Pass(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "pass", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pass is a free data retrieval call binding the contract method 0x82c4b3b2.
//
// Solidity: function pass(address ) view returns(bool)
func (_Contracts *ContractsSession) Pass(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Pass(&_Contracts.CallOpts, arg0)
}

// Pass is a free data retrieval call binding the contract method 0x82c4b3b2.
//
// Solidity: function pass(address ) view returns(bool)
func (_Contracts *ContractsCallerSession) Pass(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Pass(&_Contracts.CallOpts, arg0)
}

// ProposalLastingPeriod is a free data retrieval call binding the contract method 0xe823c814.
//
// Solidity: function proposalLastingPeriod() view returns(uint256)
func (_Contracts *ContractsCaller) ProposalLastingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "proposalLastingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalLastingPeriod is a free data retrieval call binding the contract method 0xe823c814.
//
// Solidity: function proposalLastingPeriod() view returns(uint256)
func (_Contracts *ContractsSession) ProposalLastingPeriod() (*big.Int, error) {
	return _Contracts.Contract.ProposalLastingPeriod(&_Contracts.CallOpts)
}

// ProposalLastingPeriod is a free data retrieval call binding the contract method 0xe823c814.
//
// Solidity: function proposalLastingPeriod() view returns(uint256)
func (_Contracts *ContractsCallerSession) ProposalLastingPeriod() (*big.Int, error) {
	return _Contracts.Contract.ProposalLastingPeriod(&_Contracts.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(address proposer, address dst, string details, uint256 createTime, uint16 agree, uint16 reject, bool resultExist)
func (_Contracts *ContractsCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Proposer    common.Address
	Dst         common.Address
	Details     string
	CreateTime  *big.Int
	Agree       uint16
	Reject      uint16
	ResultExist bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Proposer    common.Address
		Dst         common.Address
		Details     string
		CreateTime  *big.Int
		Agree       uint16
		Reject      uint16
		ResultExist bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proposer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Dst = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Details = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.CreateTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Agree = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.Reject = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.ResultExist = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(address proposer, address dst, string details, uint256 createTime, uint16 agree, uint16 reject, bool resultExist)
func (_Contracts *ContractsSession) Proposals(arg0 [32]byte) (struct {
	Proposer    common.Address
	Dst         common.Address
	Details     string
	CreateTime  *big.Int
	Agree       uint16
	Reject      uint16
	ResultExist bool
}, error) {
	return _Contracts.Contract.Proposals(&_Contracts.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(address proposer, address dst, string details, uint256 createTime, uint16 agree, uint16 reject, bool resultExist)
func (_Contracts *ContractsCallerSession) Proposals(arg0 [32]byte) (struct {
	Proposer    common.Address
	Dst         common.Address
	Details     string
	CreateTime  *big.Int
	Agree       uint16
	Reject      uint16
	ResultExist bool
}, error) {
	return _Contracts.Contract.Proposals(&_Contracts.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0x1db5ade8.
//
// Solidity: function votes(address , bytes32 ) view returns(address voter, uint256 voteTime, bool auth)
func (_Contracts *ContractsCaller) Votes(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	Voter    common.Address
	VoteTime *big.Int
	Auth     bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "votes", arg0, arg1)

	outstruct := new(struct {
		Voter    common.Address
		VoteTime *big.Int
		Auth     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Voter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.VoteTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Auth = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Votes is a free data retrieval call binding the contract method 0x1db5ade8.
//
// Solidity: function votes(address , bytes32 ) view returns(address voter, uint256 voteTime, bool auth)
func (_Contracts *ContractsSession) Votes(arg0 common.Address, arg1 [32]byte) (struct {
	Voter    common.Address
	VoteTime *big.Int
	Auth     bool
}, error) {
	return _Contracts.Contract.Votes(&_Contracts.CallOpts, arg0, arg1)
}

// Votes is a free data retrieval call binding the contract method 0x1db5ade8.
//
// Solidity: function votes(address , bytes32 ) view returns(address voter, uint256 voteTime, bool auth)
func (_Contracts *ContractsCallerSession) Votes(arg0 common.Address, arg1 [32]byte) (struct {
	Voter    common.Address
	VoteTime *big.Int
	Auth     bool
}, error) {
	return _Contracts.Contract.Votes(&_Contracts.CallOpts, arg0, arg1)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x1f4f7d29.
//
// Solidity: function createProposal(address dst, string details) returns(bool)
func (_Contracts *ContractsTransactor) CreateProposal(opts *bind.TransactOpts, dst common.Address, details string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createProposal", dst, details)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x1f4f7d29.
//
// Solidity: function createProposal(address dst, string details) returns(bool)
func (_Contracts *ContractsSession) CreateProposal(dst common.Address, details string) (*types.Transaction, error) {
	return _Contracts.Contract.CreateProposal(&_Contracts.TransactOpts, dst, details)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x1f4f7d29.
//
// Solidity: function createProposal(address dst, string details) returns(bool)
func (_Contracts *ContractsTransactorSession) CreateProposal(dst common.Address, details string) (*types.Transaction, error) {
	return _Contracts.Contract.CreateProposal(&_Contracts.TransactOpts, dst, details)
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

// SetUnPassed is a paid mutator transaction binding the contract method 0xfb71c430.
//
// Solidity: function setUnPassed(address val) returns(bool)
func (_Contracts *ContractsTransactor) SetUnPassed(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setUnPassed", val)
}

// SetUnPassed is a paid mutator transaction binding the contract method 0xfb71c430.
//
// Solidity: function setUnPassed(address val) returns(bool)
func (_Contracts *ContractsSession) SetUnPassed(val common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetUnPassed(&_Contracts.TransactOpts, val)
}

// SetUnPassed is a paid mutator transaction binding the contract method 0xfb71c430.
//
// Solidity: function setUnPassed(address val) returns(bool)
func (_Contracts *ContractsTransactorSession) SetUnPassed(val common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetUnPassed(&_Contracts.TransactOpts, val)
}

// VoteProposal is a paid mutator transaction binding the contract method 0xa4c4d922.
//
// Solidity: function voteProposal(bytes32 id, bool auth) returns(bool)
func (_Contracts *ContractsTransactor) VoteProposal(opts *bind.TransactOpts, id [32]byte, auth bool) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "voteProposal", id, auth)
}

// VoteProposal is a paid mutator transaction binding the contract method 0xa4c4d922.
//
// Solidity: function voteProposal(bytes32 id, bool auth) returns(bool)
func (_Contracts *ContractsSession) VoteProposal(id [32]byte, auth bool) (*types.Transaction, error) {
	return _Contracts.Contract.VoteProposal(&_Contracts.TransactOpts, id, auth)
}

// VoteProposal is a paid mutator transaction binding the contract method 0xa4c4d922.
//
// Solidity: function voteProposal(bytes32 id, bool auth) returns(bool)
func (_Contracts *ContractsTransactorSession) VoteProposal(id [32]byte, auth bool) (*types.Transaction, error) {
	return _Contracts.Contract.VoteProposal(&_Contracts.TransactOpts, id, auth)
}

// ContractsLogCreateProposalIterator is returned from FilterLogCreateProposal and is used to iterate over the raw logs and unpacked data for LogCreateProposal events raised by the Contracts contract.
type ContractsLogCreateProposalIterator struct {
	Event *ContractsLogCreateProposal // Event containing the contract specifics and raw log

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
func (it *ContractsLogCreateProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogCreateProposal)
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
		it.Event = new(ContractsLogCreateProposal)
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
func (it *ContractsLogCreateProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogCreateProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogCreateProposal represents a LogCreateProposal event raised by the Contracts contract.
type ContractsLogCreateProposal struct {
	Id       [32]byte
	Proposer common.Address
	Dst      common.Address
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogCreateProposal is a free log retrieval operation binding the contract event 0xc10f2f4d53a0e342536c6af3cce9c6ee25c32dbb323521ce0e1d4494a3e362e8.
//
// Solidity: event LogCreateProposal(bytes32 indexed id, address indexed proposer, address indexed dst, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogCreateProposal(opts *bind.FilterOpts, id [][32]byte, proposer []common.Address, dst []common.Address) (*ContractsLogCreateProposalIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogCreateProposal", idRule, proposerRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogCreateProposalIterator{contract: _Contracts.contract, event: "LogCreateProposal", logs: logs, sub: sub}, nil
}

// WatchLogCreateProposal is a free log subscription operation binding the contract event 0xc10f2f4d53a0e342536c6af3cce9c6ee25c32dbb323521ce0e1d4494a3e362e8.
//
// Solidity: event LogCreateProposal(bytes32 indexed id, address indexed proposer, address indexed dst, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogCreateProposal(opts *bind.WatchOpts, sink chan<- *ContractsLogCreateProposal, id [][32]byte, proposer []common.Address, dst []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogCreateProposal", idRule, proposerRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogCreateProposal)
				if err := _Contracts.contract.UnpackLog(event, "LogCreateProposal", log); err != nil {
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

// ParseLogCreateProposal is a log parse operation binding the contract event 0xc10f2f4d53a0e342536c6af3cce9c6ee25c32dbb323521ce0e1d4494a3e362e8.
//
// Solidity: event LogCreateProposal(bytes32 indexed id, address indexed proposer, address indexed dst, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogCreateProposal(log types.Log) (*ContractsLogCreateProposal, error) {
	event := new(ContractsLogCreateProposal)
	if err := _Contracts.contract.UnpackLog(event, "LogCreateProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogPassProposalIterator is returned from FilterLogPassProposal and is used to iterate over the raw logs and unpacked data for LogPassProposal events raised by the Contracts contract.
type ContractsLogPassProposalIterator struct {
	Event *ContractsLogPassProposal // Event containing the contract specifics and raw log

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
func (it *ContractsLogPassProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogPassProposal)
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
		it.Event = new(ContractsLogPassProposal)
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
func (it *ContractsLogPassProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogPassProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogPassProposal represents a LogPassProposal event raised by the Contracts contract.
type ContractsLogPassProposal struct {
	Id   [32]byte
	Dst  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogPassProposal is a free log retrieval operation binding the contract event 0xc9d96d61eb62031865c523ae107f3c22f5ed445af237636bcd88bea1705c70d5.
//
// Solidity: event LogPassProposal(bytes32 indexed id, address indexed dst, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogPassProposal(opts *bind.FilterOpts, id [][32]byte, dst []common.Address) (*ContractsLogPassProposalIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogPassProposal", idRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogPassProposalIterator{contract: _Contracts.contract, event: "LogPassProposal", logs: logs, sub: sub}, nil
}

// WatchLogPassProposal is a free log subscription operation binding the contract event 0xc9d96d61eb62031865c523ae107f3c22f5ed445af237636bcd88bea1705c70d5.
//
// Solidity: event LogPassProposal(bytes32 indexed id, address indexed dst, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogPassProposal(opts *bind.WatchOpts, sink chan<- *ContractsLogPassProposal, id [][32]byte, dst []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogPassProposal", idRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogPassProposal)
				if err := _Contracts.contract.UnpackLog(event, "LogPassProposal", log); err != nil {
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

// ParseLogPassProposal is a log parse operation binding the contract event 0xc9d96d61eb62031865c523ae107f3c22f5ed445af237636bcd88bea1705c70d5.
//
// Solidity: event LogPassProposal(bytes32 indexed id, address indexed dst, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogPassProposal(log types.Log) (*ContractsLogPassProposal, error) {
	event := new(ContractsLogPassProposal)
	if err := _Contracts.contract.UnpackLog(event, "LogPassProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogRejectProposalIterator is returned from FilterLogRejectProposal and is used to iterate over the raw logs and unpacked data for LogRejectProposal events raised by the Contracts contract.
type ContractsLogRejectProposalIterator struct {
	Event *ContractsLogRejectProposal // Event containing the contract specifics and raw log

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
func (it *ContractsLogRejectProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogRejectProposal)
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
		it.Event = new(ContractsLogRejectProposal)
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
func (it *ContractsLogRejectProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogRejectProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogRejectProposal represents a LogRejectProposal event raised by the Contracts contract.
type ContractsLogRejectProposal struct {
	Id   [32]byte
	Dst  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRejectProposal is a free log retrieval operation binding the contract event 0xec955d77e6e7d74e18b1c91977ef0f6fd5a6d02a28d1979686339fe693997825.
//
// Solidity: event LogRejectProposal(bytes32 indexed id, address indexed dst, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogRejectProposal(opts *bind.FilterOpts, id [][32]byte, dst []common.Address) (*ContractsLogRejectProposalIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogRejectProposal", idRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogRejectProposalIterator{contract: _Contracts.contract, event: "LogRejectProposal", logs: logs, sub: sub}, nil
}

// WatchLogRejectProposal is a free log subscription operation binding the contract event 0xec955d77e6e7d74e18b1c91977ef0f6fd5a6d02a28d1979686339fe693997825.
//
// Solidity: event LogRejectProposal(bytes32 indexed id, address indexed dst, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogRejectProposal(opts *bind.WatchOpts, sink chan<- *ContractsLogRejectProposal, id [][32]byte, dst []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogRejectProposal", idRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogRejectProposal)
				if err := _Contracts.contract.UnpackLog(event, "LogRejectProposal", log); err != nil {
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

// ParseLogRejectProposal is a log parse operation binding the contract event 0xec955d77e6e7d74e18b1c91977ef0f6fd5a6d02a28d1979686339fe693997825.
//
// Solidity: event LogRejectProposal(bytes32 indexed id, address indexed dst, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogRejectProposal(log types.Log) (*ContractsLogRejectProposal, error) {
	event := new(ContractsLogRejectProposal)
	if err := _Contracts.contract.UnpackLog(event, "LogRejectProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogSetUnPassedIterator is returned from FilterLogSetUnPassed and is used to iterate over the raw logs and unpacked data for LogSetUnPassed events raised by the Contracts contract.
type ContractsLogSetUnPassedIterator struct {
	Event *ContractsLogSetUnPassed // Event containing the contract specifics and raw log

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
func (it *ContractsLogSetUnPassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogSetUnPassed)
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
		it.Event = new(ContractsLogSetUnPassed)
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
func (it *ContractsLogSetUnPassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogSetUnPassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogSetUnPassed represents a LogSetUnPassed event raised by the Contracts contract.
type ContractsLogSetUnPassed struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogSetUnPassed is a free log retrieval operation binding the contract event 0x4b6d81f9601252010f1c1bcfacb86729b7e7898d2abc627bcefd1991dddc395e.
//
// Solidity: event LogSetUnPassed(address indexed val, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogSetUnPassed(opts *bind.FilterOpts, val []common.Address) (*ContractsLogSetUnPassedIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogSetUnPassed", valRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogSetUnPassedIterator{contract: _Contracts.contract, event: "LogSetUnPassed", logs: logs, sub: sub}, nil
}

// WatchLogSetUnPassed is a free log subscription operation binding the contract event 0x4b6d81f9601252010f1c1bcfacb86729b7e7898d2abc627bcefd1991dddc395e.
//
// Solidity: event LogSetUnPassed(address indexed val, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogSetUnPassed(opts *bind.WatchOpts, sink chan<- *ContractsLogSetUnPassed, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogSetUnPassed", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogSetUnPassed)
				if err := _Contracts.contract.UnpackLog(event, "LogSetUnPassed", log); err != nil {
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

// ParseLogSetUnPassed is a log parse operation binding the contract event 0x4b6d81f9601252010f1c1bcfacb86729b7e7898d2abc627bcefd1991dddc395e.
//
// Solidity: event LogSetUnPassed(address indexed val, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogSetUnPassed(log types.Log) (*ContractsLogSetUnPassed, error) {
	event := new(ContractsLogSetUnPassed)
	if err := _Contracts.contract.UnpackLog(event, "LogSetUnPassed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogVoteIterator is returned from FilterLogVote and is used to iterate over the raw logs and unpacked data for LogVote events raised by the Contracts contract.
type ContractsLogVoteIterator struct {
	Event *ContractsLogVote // Event containing the contract specifics and raw log

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
func (it *ContractsLogVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogVote)
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
		it.Event = new(ContractsLogVote)
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
func (it *ContractsLogVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogVote represents a LogVote event raised by the Contracts contract.
type ContractsLogVote struct {
	Id    [32]byte
	Voter common.Address
	Auth  bool
	Time  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogVote is a free log retrieval operation binding the contract event 0x6c59bda68cac318717c60c7c9635a78a0f0613f9887cc18a7157f5745a86d14e.
//
// Solidity: event LogVote(bytes32 indexed id, address indexed voter, bool auth, uint256 time)
func (_Contracts *ContractsFilterer) FilterLogVote(opts *bind.FilterOpts, id [][32]byte, voter []common.Address) (*ContractsLogVoteIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogVote", idRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogVoteIterator{contract: _Contracts.contract, event: "LogVote", logs: logs, sub: sub}, nil
}

// WatchLogVote is a free log subscription operation binding the contract event 0x6c59bda68cac318717c60c7c9635a78a0f0613f9887cc18a7157f5745a86d14e.
//
// Solidity: event LogVote(bytes32 indexed id, address indexed voter, bool auth, uint256 time)
func (_Contracts *ContractsFilterer) WatchLogVote(opts *bind.WatchOpts, sink chan<- *ContractsLogVote, id [][32]byte, voter []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogVote", idRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogVote)
				if err := _Contracts.contract.UnpackLog(event, "LogVote", log); err != nil {
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

// ParseLogVote is a log parse operation binding the contract event 0x6c59bda68cac318717c60c7c9635a78a0f0613f9887cc18a7157f5745a86d14e.
//
// Solidity: event LogVote(bytes32 indexed id, address indexed voter, bool auth, uint256 time)
func (_Contracts *ContractsFilterer) ParseLogVote(log types.Log) (*ContractsLogVote, error) {
	event := new(ContractsLogVote)
	if err := _Contracts.contract.UnpackLog(event, "LogVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
