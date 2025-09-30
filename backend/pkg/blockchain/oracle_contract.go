// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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

// OracleContractMetaData contains all meta data concerning the OracleContract contract.
var OracleContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_updater\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MAX_AGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PRICE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_PRICE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRICE_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emergencyWithdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emergencyWithdrawToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getContractInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"updaterAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxAge\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRoundId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"roundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestPriceSafe\",\"inputs\":[],\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"roundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPriceAge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isStale\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"roundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUpdater\",\"inputs\":[{\"name\":\"_newUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePrice\",\"inputs\":[{\"name\":\"_newPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EmergencyWithdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceUpdated\",\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"roundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdaterChanged\",\"inputs\":[{\"name\":\"oldUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// OracleContractABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleContractMetaData.ABI instead.
var OracleContractABI = OracleContractMetaData.ABI

// OracleContract is an auto generated Go binding around an Ethereum contract.
type OracleContract struct {
	OracleContractCaller     // Read-only binding to the contract
	OracleContractTransactor // Write-only binding to the contract
	OracleContractFilterer   // Log filterer for contract events
}

// OracleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleContractSession struct {
	Contract     *OracleContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleContractCallerSession struct {
	Contract *OracleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OracleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleContractTransactorSession struct {
	Contract     *OracleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OracleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleContractRaw struct {
	Contract *OracleContract // Generic contract binding to access the raw methods on
}

// OracleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleContractCallerRaw struct {
	Contract *OracleContractCaller // Generic read-only contract binding to access the raw methods on
}

// OracleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleContractTransactorRaw struct {
	Contract *OracleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleContract creates a new instance of OracleContract, bound to a specific deployed contract.
func NewOracleContract(address common.Address, backend bind.ContractBackend) (*OracleContract, error) {
	contract, err := bindOracleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleContract{OracleContractCaller: OracleContractCaller{contract: contract}, OracleContractTransactor: OracleContractTransactor{contract: contract}, OracleContractFilterer: OracleContractFilterer{contract: contract}}, nil
}

// NewOracleContractCaller creates a new read-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractCaller(address common.Address, caller bind.ContractCaller) (*OracleContractCaller, error) {
	contract, err := bindOracleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractCaller{contract: contract}, nil
}

// NewOracleContractTransactor creates a new write-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleContractTransactor, error) {
	contract, err := bindOracleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractTransactor{contract: contract}, nil
}

// NewOracleContractFilterer creates a new log filterer instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleContractFilterer, error) {
	contract, err := bindOracleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleContractFilterer{contract: contract}, nil
}

// bindOracleContract binds a generic wrapper to an already deployed contract.
func bindOracleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.OracleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transact(opts, method, params...)
}

// MAXAGE is a free data retrieval call binding the contract method 0x0dcaeaf2.
//
// Solidity: function MAX_AGE() view returns(uint256)
func (_OracleContract *OracleContractCaller) MAXAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "MAX_AGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXAGE is a free data retrieval call binding the contract method 0x0dcaeaf2.
//
// Solidity: function MAX_AGE() view returns(uint256)
func (_OracleContract *OracleContractSession) MAXAGE() (*big.Int, error) {
	return _OracleContract.Contract.MAXAGE(&_OracleContract.CallOpts)
}

// MAXAGE is a free data retrieval call binding the contract method 0x0dcaeaf2.
//
// Solidity: function MAX_AGE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) MAXAGE() (*big.Int, error) {
	return _OracleContract.Contract.MAXAGE(&_OracleContract.CallOpts)
}

// MAXPRICE is a free data retrieval call binding the contract method 0x01c11d96.
//
// Solidity: function MAX_PRICE() view returns(uint256)
func (_OracleContract *OracleContractCaller) MAXPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "MAX_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPRICE is a free data retrieval call binding the contract method 0x01c11d96.
//
// Solidity: function MAX_PRICE() view returns(uint256)
func (_OracleContract *OracleContractSession) MAXPRICE() (*big.Int, error) {
	return _OracleContract.Contract.MAXPRICE(&_OracleContract.CallOpts)
}

// MAXPRICE is a free data retrieval call binding the contract method 0x01c11d96.
//
// Solidity: function MAX_PRICE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) MAXPRICE() (*big.Int, error) {
	return _OracleContract.Contract.MAXPRICE(&_OracleContract.CallOpts)
}

// MINPRICE is a free data retrieval call binding the contract method 0xad9f20a6.
//
// Solidity: function MIN_PRICE() view returns(uint256)
func (_OracleContract *OracleContractCaller) MINPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "MIN_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINPRICE is a free data retrieval call binding the contract method 0xad9f20a6.
//
// Solidity: function MIN_PRICE() view returns(uint256)
func (_OracleContract *OracleContractSession) MINPRICE() (*big.Int, error) {
	return _OracleContract.Contract.MINPRICE(&_OracleContract.CallOpts)
}

// MINPRICE is a free data retrieval call binding the contract method 0xad9f20a6.
//
// Solidity: function MIN_PRICE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) MINPRICE() (*big.Int, error) {
	return _OracleContract.Contract.MINPRICE(&_OracleContract.CallOpts)
}

// PRICEDECIMALS is a free data retrieval call binding the contract method 0xf1a640f8.
//
// Solidity: function PRICE_DECIMALS() view returns(uint256)
func (_OracleContract *OracleContractCaller) PRICEDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "PRICE_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEDECIMALS is a free data retrieval call binding the contract method 0xf1a640f8.
//
// Solidity: function PRICE_DECIMALS() view returns(uint256)
func (_OracleContract *OracleContractSession) PRICEDECIMALS() (*big.Int, error) {
	return _OracleContract.Contract.PRICEDECIMALS(&_OracleContract.CallOpts)
}

// PRICEDECIMALS is a free data retrieval call binding the contract method 0xf1a640f8.
//
// Solidity: function PRICE_DECIMALS() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) PRICEDECIMALS() (*big.Int, error) {
	return _OracleContract.Contract.PRICEDECIMALS(&_OracleContract.CallOpts)
}

// GetContractInfo is a free data retrieval call binding the contract method 0x7cc1f867.
//
// Solidity: function getContractInfo() view returns(address updaterAddress, bool isPaused, uint256 maxAge, uint256 minPrice, uint256 maxPrice)
func (_OracleContract *OracleContractCaller) GetContractInfo(opts *bind.CallOpts) (struct {
	UpdaterAddress common.Address
	IsPaused       bool
	MaxAge         *big.Int
	MinPrice       *big.Int
	MaxPrice       *big.Int
}, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "getContractInfo")

	outstruct := new(struct {
		UpdaterAddress common.Address
		IsPaused       bool
		MaxAge         *big.Int
		MinPrice       *big.Int
		MaxPrice       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UpdaterAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsPaused = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.MaxAge = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MinPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetContractInfo is a free data retrieval call binding the contract method 0x7cc1f867.
//
// Solidity: function getContractInfo() view returns(address updaterAddress, bool isPaused, uint256 maxAge, uint256 minPrice, uint256 maxPrice)
func (_OracleContract *OracleContractSession) GetContractInfo() (struct {
	UpdaterAddress common.Address
	IsPaused       bool
	MaxAge         *big.Int
	MinPrice       *big.Int
	MaxPrice       *big.Int
}, error) {
	return _OracleContract.Contract.GetContractInfo(&_OracleContract.CallOpts)
}

// GetContractInfo is a free data retrieval call binding the contract method 0x7cc1f867.
//
// Solidity: function getContractInfo() view returns(address updaterAddress, bool isPaused, uint256 maxAge, uint256 minPrice, uint256 maxPrice)
func (_OracleContract *OracleContractCallerSession) GetContractInfo() (struct {
	UpdaterAddress common.Address
	IsPaused       bool
	MaxAge         *big.Int
	MinPrice       *big.Int
	MaxPrice       *big.Int
}, error) {
	return _OracleContract.Contract.GetContractInfo(&_OracleContract.CallOpts)
}

// GetCurrentRoundId is a free data retrieval call binding the contract method 0x5727e25d.
//
// Solidity: function getCurrentRoundId() view returns(uint256)
func (_OracleContract *OracleContractCaller) GetCurrentRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "getCurrentRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRoundId is a free data retrieval call binding the contract method 0x5727e25d.
//
// Solidity: function getCurrentRoundId() view returns(uint256)
func (_OracleContract *OracleContractSession) GetCurrentRoundId() (*big.Int, error) {
	return _OracleContract.Contract.GetCurrentRoundId(&_OracleContract.CallOpts)
}

// GetCurrentRoundId is a free data retrieval call binding the contract method 0x5727e25d.
//
// Solidity: function getCurrentRoundId() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) GetCurrentRoundId() (*big.Int, error) {
	return _OracleContract.Contract.GetCurrentRoundId(&_OracleContract.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(uint256 price, uint256 timestamp, uint256 roundId)
func (_OracleContract *OracleContractCaller) GetLatestPrice(opts *bind.CallOpts) (struct {
	Price     *big.Int
	Timestamp *big.Int
	RoundId   *big.Int
}, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "getLatestPrice")

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
		RoundId   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RoundId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(uint256 price, uint256 timestamp, uint256 roundId)
func (_OracleContract *OracleContractSession) GetLatestPrice() (struct {
	Price     *big.Int
	Timestamp *big.Int
	RoundId   *big.Int
}, error) {
	return _OracleContract.Contract.GetLatestPrice(&_OracleContract.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(uint256 price, uint256 timestamp, uint256 roundId)
func (_OracleContract *OracleContractCallerSession) GetLatestPrice() (struct {
	Price     *big.Int
	Timestamp *big.Int
	RoundId   *big.Int
}, error) {
	return _OracleContract.Contract.GetLatestPrice(&_OracleContract.CallOpts)
}

// GetLatestPriceSafe is a free data retrieval call binding the contract method 0x09e03e29.
//
// Solidity: function getLatestPriceSafe() view returns(uint256 price, uint256 timestamp, uint256 roundId)
func (_OracleContract *OracleContractCaller) GetLatestPriceSafe(opts *bind.CallOpts) (struct {
	Price     *big.Int
	Timestamp *big.Int
	RoundId   *big.Int
}, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "getLatestPriceSafe")

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
		RoundId   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RoundId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLatestPriceSafe is a free data retrieval call binding the contract method 0x09e03e29.
//
// Solidity: function getLatestPriceSafe() view returns(uint256 price, uint256 timestamp, uint256 roundId)
func (_OracleContract *OracleContractSession) GetLatestPriceSafe() (struct {
	Price     *big.Int
	Timestamp *big.Int
	RoundId   *big.Int
}, error) {
	return _OracleContract.Contract.GetLatestPriceSafe(&_OracleContract.CallOpts)
}

// GetLatestPriceSafe is a free data retrieval call binding the contract method 0x09e03e29.
//
// Solidity: function getLatestPriceSafe() view returns(uint256 price, uint256 timestamp, uint256 roundId)
func (_OracleContract *OracleContractCallerSession) GetLatestPriceSafe() (struct {
	Price     *big.Int
	Timestamp *big.Int
	RoundId   *big.Int
}, error) {
	return _OracleContract.Contract.GetLatestPriceSafe(&_OracleContract.CallOpts)
}

// GetPriceAge is a free data retrieval call binding the contract method 0xa4ed446d.
//
// Solidity: function getPriceAge() view returns(uint256)
func (_OracleContract *OracleContractCaller) GetPriceAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "getPriceAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceAge is a free data retrieval call binding the contract method 0xa4ed446d.
//
// Solidity: function getPriceAge() view returns(uint256)
func (_OracleContract *OracleContractSession) GetPriceAge() (*big.Int, error) {
	return _OracleContract.Contract.GetPriceAge(&_OracleContract.CallOpts)
}

// GetPriceAge is a free data retrieval call binding the contract method 0xa4ed446d.
//
// Solidity: function getPriceAge() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) GetPriceAge() (*big.Int, error) {
	return _OracleContract.Contract.GetPriceAge(&_OracleContract.CallOpts)
}

// IsStale is a free data retrieval call binding the contract method 0x1a26f447.
//
// Solidity: function isStale() view returns(bool)
func (_OracleContract *OracleContractCaller) IsStale(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "isStale")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStale is a free data retrieval call binding the contract method 0x1a26f447.
//
// Solidity: function isStale() view returns(bool)
func (_OracleContract *OracleContractSession) IsStale() (bool, error) {
	return _OracleContract.Contract.IsStale(&_OracleContract.CallOpts)
}

// IsStale is a free data retrieval call binding the contract method 0x1a26f447.
//
// Solidity: function isStale() view returns(bool)
func (_OracleContract *OracleContractCallerSession) IsStale() (bool, error) {
	return _OracleContract.Contract.IsStale(&_OracleContract.CallOpts)
}

// LatestPrice is a free data retrieval call binding the contract method 0xa3e6ba94.
//
// Solidity: function latestPrice() view returns(uint256 price, uint256 timestamp, uint256 roundId)
func (_OracleContract *OracleContractCaller) LatestPrice(opts *bind.CallOpts) (struct {
	Price     *big.Int
	Timestamp *big.Int
	RoundId   *big.Int
}, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "latestPrice")

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
		RoundId   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RoundId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestPrice is a free data retrieval call binding the contract method 0xa3e6ba94.
//
// Solidity: function latestPrice() view returns(uint256 price, uint256 timestamp, uint256 roundId)
func (_OracleContract *OracleContractSession) LatestPrice() (struct {
	Price     *big.Int
	Timestamp *big.Int
	RoundId   *big.Int
}, error) {
	return _OracleContract.Contract.LatestPrice(&_OracleContract.CallOpts)
}

// LatestPrice is a free data retrieval call binding the contract method 0xa3e6ba94.
//
// Solidity: function latestPrice() view returns(uint256 price, uint256 timestamp, uint256 roundId)
func (_OracleContract *OracleContractCallerSession) LatestPrice() (struct {
	Price     *big.Int
	Timestamp *big.Int
	RoundId   *big.Int
}, error) {
	return _OracleContract.Contract.LatestPrice(&_OracleContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleContract *OracleContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleContract *OracleContractSession) Owner() (common.Address, error) {
	return _OracleContract.Contract.Owner(&_OracleContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleContract *OracleContractCallerSession) Owner() (common.Address, error) {
	return _OracleContract.Contract.Owner(&_OracleContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OracleContract *OracleContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OracleContract *OracleContractSession) Paused() (bool, error) {
	return _OracleContract.Contract.Paused(&_OracleContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OracleContract *OracleContractCallerSession) Paused() (bool, error) {
	return _OracleContract.Contract.Paused(&_OracleContract.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_OracleContract *OracleContractCaller) Updater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "updater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_OracleContract *OracleContractSession) Updater() (common.Address, error) {
	return _OracleContract.Contract.Updater(&_OracleContract.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_OracleContract *OracleContractCallerSession) Updater() (common.Address, error) {
	return _OracleContract.Contract.Updater(&_OracleContract.CallOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_OracleContract *OracleContractTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_OracleContract *OracleContractSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _OracleContract.Contract.EmergencyWithdraw(&_OracleContract.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_OracleContract *OracleContractTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _OracleContract.Contract.EmergencyWithdraw(&_OracleContract.TransactOpts)
}

// EmergencyWithdrawToken is a paid mutator transaction binding the contract method 0x1af03203.
//
// Solidity: function emergencyWithdrawToken(address _token) returns()
func (_OracleContract *OracleContractTransactor) EmergencyWithdrawToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "emergencyWithdrawToken", _token)
}

// EmergencyWithdrawToken is a paid mutator transaction binding the contract method 0x1af03203.
//
// Solidity: function emergencyWithdrawToken(address _token) returns()
func (_OracleContract *OracleContractSession) EmergencyWithdrawToken(_token common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.EmergencyWithdrawToken(&_OracleContract.TransactOpts, _token)
}

// EmergencyWithdrawToken is a paid mutator transaction binding the contract method 0x1af03203.
//
// Solidity: function emergencyWithdrawToken(address _token) returns()
func (_OracleContract *OracleContractTransactorSession) EmergencyWithdrawToken(_token common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.EmergencyWithdrawToken(&_OracleContract.TransactOpts, _token)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleContract *OracleContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleContract *OracleContractSession) Pause() (*types.Transaction, error) {
	return _OracleContract.Contract.Pause(&_OracleContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleContract *OracleContractTransactorSession) Pause() (*types.Transaction, error) {
	return _OracleContract.Contract.Pause(&_OracleContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleContract *OracleContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleContract *OracleContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _OracleContract.Contract.RenounceOwnership(&_OracleContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleContract *OracleContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OracleContract.Contract.RenounceOwnership(&_OracleContract.TransactOpts)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _newUpdater) returns()
func (_OracleContract *OracleContractTransactor) SetUpdater(opts *bind.TransactOpts, _newUpdater common.Address) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "setUpdater", _newUpdater)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _newUpdater) returns()
func (_OracleContract *OracleContractSession) SetUpdater(_newUpdater common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.SetUpdater(&_OracleContract.TransactOpts, _newUpdater)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _newUpdater) returns()
func (_OracleContract *OracleContractTransactorSession) SetUpdater(_newUpdater common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.SetUpdater(&_OracleContract.TransactOpts, _newUpdater)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OracleContract *OracleContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OracleContract *OracleContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.TransferOwnership(&_OracleContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OracleContract *OracleContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.TransferOwnership(&_OracleContract.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleContract *OracleContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleContract *OracleContractSession) Unpause() (*types.Transaction, error) {
	return _OracleContract.Contract.Unpause(&_OracleContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleContract *OracleContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _OracleContract.Contract.Unpause(&_OracleContract.TransactOpts)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 _newPrice) returns()
func (_OracleContract *OracleContractTransactor) UpdatePrice(opts *bind.TransactOpts, _newPrice *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "updatePrice", _newPrice)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 _newPrice) returns()
func (_OracleContract *OracleContractSession) UpdatePrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.UpdatePrice(&_OracleContract.TransactOpts, _newPrice)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 _newPrice) returns()
func (_OracleContract *OracleContractTransactorSession) UpdatePrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.UpdatePrice(&_OracleContract.TransactOpts, _newPrice)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OracleContract *OracleContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OracleContract *OracleContractSession) Receive() (*types.Transaction, error) {
	return _OracleContract.Contract.Receive(&_OracleContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OracleContract *OracleContractTransactorSession) Receive() (*types.Transaction, error) {
	return _OracleContract.Contract.Receive(&_OracleContract.TransactOpts)
}

// OracleContractEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the OracleContract contract.
type OracleContractEmergencyWithdrawIterator struct {
	Event *OracleContractEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *OracleContractEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractEmergencyWithdraw)
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
		it.Event = new(OracleContractEmergencyWithdraw)
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
func (it *OracleContractEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractEmergencyWithdraw represents a EmergencyWithdraw event raised by the OracleContract contract.
type OracleContractEmergencyWithdraw struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed token, uint256 amount)
func (_OracleContract *OracleContractFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, token []common.Address) (*OracleContractEmergencyWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "EmergencyWithdraw", tokenRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractEmergencyWithdrawIterator{contract: _OracleContract.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed token, uint256 amount)
func (_OracleContract *OracleContractFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *OracleContractEmergencyWithdraw, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "EmergencyWithdraw", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractEmergencyWithdraw)
				if err := _OracleContract.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed token, uint256 amount)
func (_OracleContract *OracleContractFilterer) ParseEmergencyWithdraw(log types.Log) (*OracleContractEmergencyWithdraw, error) {
	event := new(OracleContractEmergencyWithdraw)
	if err := _OracleContract.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OracleContract contract.
type OracleContractOwnershipTransferredIterator struct {
	Event *OracleContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OracleContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractOwnershipTransferred)
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
		it.Event = new(OracleContractOwnershipTransferred)
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
func (it *OracleContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractOwnershipTransferred represents a OwnershipTransferred event raised by the OracleContract contract.
type OracleContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OracleContract *OracleContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OracleContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractOwnershipTransferredIterator{contract: _OracleContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OracleContract *OracleContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OracleContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractOwnershipTransferred)
				if err := _OracleContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OracleContract *OracleContractFilterer) ParseOwnershipTransferred(log types.Log) (*OracleContractOwnershipTransferred, error) {
	event := new(OracleContractOwnershipTransferred)
	if err := _OracleContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OracleContract contract.
type OracleContractPausedIterator struct {
	Event *OracleContractPaused // Event containing the contract specifics and raw log

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
func (it *OracleContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractPaused)
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
		it.Event = new(OracleContractPaused)
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
func (it *OracleContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractPaused represents a Paused event raised by the OracleContract contract.
type OracleContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OracleContract *OracleContractFilterer) FilterPaused(opts *bind.FilterOpts) (*OracleContractPausedIterator, error) {

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OracleContractPausedIterator{contract: _OracleContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OracleContract *OracleContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OracleContractPaused) (event.Subscription, error) {

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractPaused)
				if err := _OracleContract.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OracleContract *OracleContractFilterer) ParsePaused(log types.Log) (*OracleContractPaused, error) {
	event := new(OracleContractPaused)
	if err := _OracleContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the OracleContract contract.
type OracleContractPriceUpdatedIterator struct {
	Event *OracleContractPriceUpdated // Event containing the contract specifics and raw log

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
func (it *OracleContractPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractPriceUpdated)
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
		it.Event = new(OracleContractPriceUpdated)
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
func (it *OracleContractPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractPriceUpdated represents a PriceUpdated event raised by the OracleContract contract.
type OracleContractPriceUpdated struct {
	Price     *big.Int
	Timestamp *big.Int
	RoundId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0x15819dd2fd9f6418b142e798d08a18d0bf06ea368f4480b7b0d3f75bd966bc48.
//
// Solidity: event PriceUpdated(uint256 indexed price, uint256 indexed timestamp, uint256 indexed roundId)
func (_OracleContract *OracleContractFilterer) FilterPriceUpdated(opts *bind.FilterOpts, price []*big.Int, timestamp []*big.Int, roundId []*big.Int) (*OracleContractPriceUpdatedIterator, error) {

	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "PriceUpdated", priceRule, timestampRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractPriceUpdatedIterator{contract: _OracleContract.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0x15819dd2fd9f6418b142e798d08a18d0bf06ea368f4480b7b0d3f75bd966bc48.
//
// Solidity: event PriceUpdated(uint256 indexed price, uint256 indexed timestamp, uint256 indexed roundId)
func (_OracleContract *OracleContractFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *OracleContractPriceUpdated, price []*big.Int, timestamp []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "PriceUpdated", priceRule, timestampRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractPriceUpdated)
				if err := _OracleContract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
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

// ParsePriceUpdated is a log parse operation binding the contract event 0x15819dd2fd9f6418b142e798d08a18d0bf06ea368f4480b7b0d3f75bd966bc48.
//
// Solidity: event PriceUpdated(uint256 indexed price, uint256 indexed timestamp, uint256 indexed roundId)
func (_OracleContract *OracleContractFilterer) ParsePriceUpdated(log types.Log) (*OracleContractPriceUpdated, error) {
	event := new(OracleContractPriceUpdated)
	if err := _OracleContract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OracleContract contract.
type OracleContractUnpausedIterator struct {
	Event *OracleContractUnpaused // Event containing the contract specifics and raw log

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
func (it *OracleContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractUnpaused)
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
		it.Event = new(OracleContractUnpaused)
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
func (it *OracleContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractUnpaused represents a Unpaused event raised by the OracleContract contract.
type OracleContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OracleContract *OracleContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OracleContractUnpausedIterator, error) {

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OracleContractUnpausedIterator{contract: _OracleContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OracleContract *OracleContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OracleContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractUnpaused)
				if err := _OracleContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OracleContract *OracleContractFilterer) ParseUnpaused(log types.Log) (*OracleContractUnpaused, error) {
	event := new(OracleContractUnpaused)
	if err := _OracleContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractUpdaterChangedIterator is returned from FilterUpdaterChanged and is used to iterate over the raw logs and unpacked data for UpdaterChanged events raised by the OracleContract contract.
type OracleContractUpdaterChangedIterator struct {
	Event *OracleContractUpdaterChanged // Event containing the contract specifics and raw log

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
func (it *OracleContractUpdaterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractUpdaterChanged)
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
		it.Event = new(OracleContractUpdaterChanged)
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
func (it *OracleContractUpdaterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractUpdaterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractUpdaterChanged represents a UpdaterChanged event raised by the OracleContract contract.
type OracleContractUpdaterChanged struct {
	OldUpdater common.Address
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdaterChanged is a free log retrieval operation binding the contract event 0x662a4a4a892f5f13cf7ee050fdaa045f8641601fdbc843e8a71f418099cacd4e.
//
// Solidity: event UpdaterChanged(address indexed oldUpdater, address indexed newUpdater)
func (_OracleContract *OracleContractFilterer) FilterUpdaterChanged(opts *bind.FilterOpts, oldUpdater []common.Address, newUpdater []common.Address) (*OracleContractUpdaterChangedIterator, error) {

	var oldUpdaterRule []interface{}
	for _, oldUpdaterItem := range oldUpdater {
		oldUpdaterRule = append(oldUpdaterRule, oldUpdaterItem)
	}
	var newUpdaterRule []interface{}
	for _, newUpdaterItem := range newUpdater {
		newUpdaterRule = append(newUpdaterRule, newUpdaterItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "UpdaterChanged", oldUpdaterRule, newUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractUpdaterChangedIterator{contract: _OracleContract.contract, event: "UpdaterChanged", logs: logs, sub: sub}, nil
}

// WatchUpdaterChanged is a free log subscription operation binding the contract event 0x662a4a4a892f5f13cf7ee050fdaa045f8641601fdbc843e8a71f418099cacd4e.
//
// Solidity: event UpdaterChanged(address indexed oldUpdater, address indexed newUpdater)
func (_OracleContract *OracleContractFilterer) WatchUpdaterChanged(opts *bind.WatchOpts, sink chan<- *OracleContractUpdaterChanged, oldUpdater []common.Address, newUpdater []common.Address) (event.Subscription, error) {

	var oldUpdaterRule []interface{}
	for _, oldUpdaterItem := range oldUpdater {
		oldUpdaterRule = append(oldUpdaterRule, oldUpdaterItem)
	}
	var newUpdaterRule []interface{}
	for _, newUpdaterItem := range newUpdater {
		newUpdaterRule = append(newUpdaterRule, newUpdaterItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "UpdaterChanged", oldUpdaterRule, newUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractUpdaterChanged)
				if err := _OracleContract.contract.UnpackLog(event, "UpdaterChanged", log); err != nil {
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

// ParseUpdaterChanged is a log parse operation binding the contract event 0x662a4a4a892f5f13cf7ee050fdaa045f8641601fdbc843e8a71f418099cacd4e.
//
// Solidity: event UpdaterChanged(address indexed oldUpdater, address indexed newUpdater)
func (_OracleContract *OracleContractFilterer) ParseUpdaterChanged(log types.Log) (*OracleContractUpdaterChanged, error) {
	event := new(OracleContractUpdaterChanged)
	if err := _OracleContract.contract.UnpackLog(event, "UpdaterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
