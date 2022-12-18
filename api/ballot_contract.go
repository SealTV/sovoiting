// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proposalNames\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isVotingActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposal\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"voted\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winnerNane\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winningProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"winningProposal_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610b73380380610b7383398101604081905261002f916100ec565b600080546001600160a01b03191633908117825581526001602081905260408220555b81518110156100c5576002604051806040016040528084848151811061007a5761007a6101a9565b602090810291909101810151825260009181018290528354600181810186559483529181902083516002909302019182559190910151910155806100bd816101bf565b915050610052565b50506003805460ff191690556101e6565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156100ff57600080fd5b82516001600160401b038082111561011657600080fd5b818501915085601f83011261012a57600080fd5b81518181111561013c5761013c6100d6565b8060051b604051601f19603f83011681018181108582111715610161576101616100d6565b60405291825284820192508381018501918883111561017f57600080fd5b938501935b8285101561019d57845184529385019392850192610184565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600182016101df57634e487b7160e01b600052601160045260246000fd5b5060010190565b61097e806101f56000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80635c19a95c116100715780635c19a95c1461013b578063609ff1bd1461014e5780636a832edb146101645780639e7b8d611461016c578063a3ec138d1461017f578063c3403ddf146101f057600080fd5b80630121b93f146100ae578063013cf08b146100c35780631ec6b60a146100f05780632e4176cf146100f857806347cf6f7614610123575b600080fd5b6100c16100bc36600461088b565b6101f8565b005b6100d66100d136600461088b565b610347565b604080519283526020830191909152015b60405180910390f35b6100c1610375565b60005461010b906001600160a01b031681565b6040516001600160a01b0390911681526020016100e7565b61012b6103e5565b60405190151581526020016100e7565b6100c16101493660046108a4565b610406565b610156610655565b6040519081526020016100e7565b6101566106d2565b6100c161017a3660046108a4565b610705565b6101c161018d3660046108a4565b600160208190526000918252604090912080549181015460029091015460ff82169161010090046001600160a01b03169084565b6040516100e7949392919093845291151560208401526001600160a01b03166040830152606082015260800190565b6100c161081d565b33600090815260016020526040812080549091036102555760405162461bcd60e51b81526020600482015260156024820152742430b9903737903934b3b43a103a37903b37ba329760591b60448201526064015b60405180910390fd5b600181015460ff161561029b5760405162461bcd60e51b815260206004820152600e60248201526d20b63932b0b23c903b37ba32b21760911b604482015260640161024c565b6102a36103e5565b6102ef5760405162461bcd60e51b815260206004820152601760248201527f566f74696e672073686f756c6420626520416374697665000000000000000000604482015260640161024c565b6001818101805460ff191690911790556002808201839055815481549091908490811061031e5761031e6108d4565b9060005260206000209060020201600101600082825461033e9190610900565b90915550505050565b6002818154811061035757600080fd5b60009182526020909120600290910201805460019091015490915082565b6000546001600160a01b031633146103cf5760405162461bcd60e51b815260206004820152601e60248201527f4f6e6c79206368616972706572736f6e20737461727420766f74696e672e0000604482015260640161024c565b600380546001919060ff191682805b0217905550565b6000600160035460ff16600281111561040057610400610919565b14905090565b33600090815260016020526040812080549091036104665760405162461bcd60e51b815260206004820152601a60248201527f596f752068617665206e6f20726967687420746f20766f74652e000000000000604482015260640161024c565b600181015460ff16156104b05760405162461bcd60e51b81526020600482015260126024820152712cb7ba9030b63932b0b23c903b37ba32b21760711b604482015260640161024c565b336001600160a01b038316036105085760405162461bcd60e51b815260206004820152601e60248201527f53656c662d64656c65676174696f6e20697320646973616c6c6f7765642e0000604482015260640161024c565b6001600160a01b0382811660009081526001602081905260409091200154610100900416156105ac576001600160a01b03918216600090815260016020819052604090912001546101009004909116903382036105a75760405162461bcd60e51b815260206004820152601960248201527f466f756e64206c6f6f7020696e2064656c65676174696f6e2e00000000000000604482015260640161024c565b610508565b6001600160a01b03821660009081526001602081905260409091208054909111156105d657600080fd5b600182810180546001600160a81b0319166101006001600160a01b03871602178217905581015460ff16156106375781546002828101548154811061061d5761061d6108d4565b906000526020600020906002020160010181905550505050565b81548154829060009061064b908490610900565b9091555050505050565b600080805b6002548110156106cd578160028281548110610678576106786108d4565b90600052602060002090600202016001015411156106bb57600281815481106106a3576106a36108d4565b90600052602060002090600202016001015491508092505b806106c58161092f565b91505061065a565b505090565b600060026106de610655565b815481106106ee576106ee6108d4565b906000526020600020906002020160000154905090565b6000546001600160a01b031633146107705760405162461bcd60e51b815260206004820152602860248201527f4f6e6c79206368616972706572736f6e2063616e2067697665207269676874206044820152673a37903b37ba329760c11b606482015260840161024c565b6001600160a01b0381166000908152600160208190526040909120015460ff16156107dd5760405162461bcd60e51b815260206004820152601860248201527f54686520766f74657220616c726561647920766f7465642e0000000000000000604482015260640161024c565b6001600160a01b0381166000908152600160205260409020541561080057600080fd5b6001600160a01b0316600090815260016020819052604090912055565b6000546001600160a01b031633146108775760405162461bcd60e51b815260206004820152601c60248201527f4f6e6c79206368616972706572736f6e20656e6420766f74696e672e00000000604482015260640161024c565b600380546002919060ff19166001836103de565b60006020828403121561089d57600080fd5b5035919050565b6000602082840312156108b657600080fd5b81356001600160a01b03811681146108cd57600080fd5b9392505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610913576109136108ea565b92915050565b634e487b7160e01b600052602160045260246000fd5b600060018201610941576109416108ea565b506001019056fea264697066735822122073cf6d72f79150e472426892f0dbe6b36e23c4552331828cd6c48829a733c01a64736f6c63430008110033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, proposalNames [][32]byte) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, proposalNames)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Api *ApiCaller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "chairperson")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Api *ApiSession) Chairperson() (common.Address, error) {
	return _Api.Contract.Chairperson(&_Api.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Api *ApiCallerSession) Chairperson() (common.Address, error) {
	return _Api.Contract.Chairperson(&_Api.CallOpts)
}

// IsVotingActive is a free data retrieval call binding the contract method 0x47cf6f76.
//
// Solidity: function isVotingActive() view returns(bool isActive)
func (_Api *ApiCaller) IsVotingActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "isVotingActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVotingActive is a free data retrieval call binding the contract method 0x47cf6f76.
//
// Solidity: function isVotingActive() view returns(bool isActive)
func (_Api *ApiSession) IsVotingActive() (bool, error) {
	return _Api.Contract.IsVotingActive(&_Api.CallOpts)
}

// IsVotingActive is a free data retrieval call binding the contract method 0x47cf6f76.
//
// Solidity: function isVotingActive() view returns(bool isActive)
func (_Api *ApiCallerSession) IsVotingActive() (bool, error) {
	return _Api.Contract.IsVotingActive(&_Api.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Api *ApiCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Name      [32]byte
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Api *ApiSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Api.Contract.Proposals(&_Api.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Api *ApiCallerSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Api.Contract.Proposals(&_Api.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Api *ApiCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		Weight   *big.Int
		Voted    bool
		Delegate common.Address
		Vote     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Weight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Voted = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Delegate = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Vote = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Api *ApiSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Api.Contract.Voters(&_Api.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Api *ApiCallerSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Api.Contract.Voters(&_Api.CallOpts, arg0)
}

// WinnerNane is a free data retrieval call binding the contract method 0x6a832edb.
//
// Solidity: function winnerNane() view returns(bytes32 name)
func (_Api *ApiCaller) WinnerNane(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "winnerNane")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinnerNane is a free data retrieval call binding the contract method 0x6a832edb.
//
// Solidity: function winnerNane() view returns(bytes32 name)
func (_Api *ApiSession) WinnerNane() ([32]byte, error) {
	return _Api.Contract.WinnerNane(&_Api.CallOpts)
}

// WinnerNane is a free data retrieval call binding the contract method 0x6a832edb.
//
// Solidity: function winnerNane() view returns(bytes32 name)
func (_Api *ApiCallerSession) WinnerNane() ([32]byte, error) {
	return _Api.Contract.WinnerNane(&_Api.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Api *ApiCaller) WinningProposal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "winningProposal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Api *ApiSession) WinningProposal() (*big.Int, error) {
	return _Api.Contract.WinningProposal(&_Api.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Api *ApiCallerSession) WinningProposal() (*big.Int, error) {
	return _Api.Contract.WinningProposal(&_Api.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Api *ApiTransactor) Delegate(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "delegate", to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Api *ApiSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Api.Contract.Delegate(&_Api.TransactOpts, to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Api *ApiTransactorSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Api.Contract.Delegate(&_Api.TransactOpts, to)
}

// EndVoting is a paid mutator transaction binding the contract method 0xc3403ddf.
//
// Solidity: function endVoting() returns()
func (_Api *ApiTransactor) EndVoting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "endVoting")
}

// EndVoting is a paid mutator transaction binding the contract method 0xc3403ddf.
//
// Solidity: function endVoting() returns()
func (_Api *ApiSession) EndVoting() (*types.Transaction, error) {
	return _Api.Contract.EndVoting(&_Api.TransactOpts)
}

// EndVoting is a paid mutator transaction binding the contract method 0xc3403ddf.
//
// Solidity: function endVoting() returns()
func (_Api *ApiTransactorSession) EndVoting() (*types.Transaction, error) {
	return _Api.Contract.EndVoting(&_Api.TransactOpts)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Api *ApiTransactor) GiveRightToVote(opts *bind.TransactOpts, voter common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "giveRightToVote", voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Api *ApiSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Api.Contract.GiveRightToVote(&_Api.TransactOpts, voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Api *ApiTransactorSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Api.Contract.GiveRightToVote(&_Api.TransactOpts, voter)
}

// StartVoting is a paid mutator transaction binding the contract method 0x1ec6b60a.
//
// Solidity: function startVoting() returns()
func (_Api *ApiTransactor) StartVoting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "startVoting")
}

// StartVoting is a paid mutator transaction binding the contract method 0x1ec6b60a.
//
// Solidity: function startVoting() returns()
func (_Api *ApiSession) StartVoting() (*types.Transaction, error) {
	return _Api.Contract.StartVoting(&_Api.TransactOpts)
}

// StartVoting is a paid mutator transaction binding the contract method 0x1ec6b60a.
//
// Solidity: function startVoting() returns()
func (_Api *ApiTransactorSession) StartVoting() (*types.Transaction, error) {
	return _Api.Contract.StartVoting(&_Api.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Api *ApiTransactor) Vote(opts *bind.TransactOpts, proposal *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "vote", proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Api *ApiSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Vote(&_Api.TransactOpts, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Api *ApiTransactorSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Vote(&_Api.TransactOpts, proposal)
}
