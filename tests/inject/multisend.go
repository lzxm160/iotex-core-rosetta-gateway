// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package inject

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

// MultisendABI is the input ABI used to generate the binding from.
const MultisendABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"payload\",\"type\":\"string\"}],\"name\":\"Payload\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"payload\",\"type\":\"string\"}],\"name\":\"multiSend\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// MultisendFuncSigs maps the 4-byte function signature to its string representation.
var MultisendFuncSigs = map[string]string{
	"e3b48f48": "multiSend(address[],uint256[],string)",
}

// MultisendBin is the compiled bytecode used for deploying new contracts.
var MultisendBin = "0x608060405234801561001057600080fd5b50610525806100206000396000f3fe60806040526004361061001e5760003560e01c8063e3b48f4814610023575b600080fd5b6101d16004803603606081101561003957600080fd5b81019060208101813564010000000081111561005457600080fd5b82018360208201111561006657600080fd5b8035906020019184602083028401116401000000008311171561008857600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092959493602081019350359150506401000000008111156100d857600080fd5b8201836020820111156100ea57600080fd5b8035906020019184602083028401116401000000008311171561010c57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561015c57600080fd5b82018360208201111561016e57600080fd5b8035906020019184600183028401116401000000008311171561019057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506101d3945050505050565b005b61012c835111156102155760405162461bcd60e51b81526004018080602001828103825260278152602001806104ca6027913960400191505060405180910390fd5b8151835114610262576040805162461bcd60e51b81526020600482015260146024820152730e0c2e4c2dacae8cae4e640dcdee840dac2e8c6d60631b604482015290519081900360640190fd5b6000805b84518110156102945783818151811061027b57fe5b6020026020010151820191508080600101915050610266565b50803410156102dd576040805162461bcd60e51b815260206004820152601060248201526f3737ba1032b737bab3b4103a37b5b2b760811b604482015290519081900360640190fd5b3481900360005b85518110156103c0578581815181106102f957fe5b60200260200101516001600160a01b03166108fc86838151811061031957fe5b60200260200101519081150290604051600060405180830381858888f1935050505015801561034c573d6000803e3d6000fd5b507f69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de286828151811061037a57fe5b602002602001015186838151811061038e57fe5b602090810291909101810151604080516001600160a01b039094168452918301528051918290030190a16001016102e4565b50801561042957604051339082156108fc029083906000818181858888f193505050501580156103f4573d6000803e3d6000fd5b506040805182815290517f2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b979181900360200190a15b7f53a85291e316c24064ff2c7668d99f35ecbb40ef4e24794ff9d8abe901c7e62c836040518080602001828103825283818151815260200191508051906020019080838360005b83811015610488578181015183820152602001610470565b50505050905090810190601f1680156104b55780820380516001836020036101000a031916815260200191505b509250505060405180910390a1505050505056fe6e756d626572206f6620726563697069656e7473206973206c6172676572207468616e20333030a265627a7a7231582020cfda08ba8deef2a926a0d7dd2db36d5dfa38a398f59885da2efd2367b6919b64736f6c63430005100032"

// DeployMultisend deploys a new Ethereum contract, binding an instance of Multisend to it.
func DeployMultisend(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Multisend, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisendABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultisendBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multisend{MultisendCaller: MultisendCaller{contract: contract}, MultisendTransactor: MultisendTransactor{contract: contract}, MultisendFilterer: MultisendFilterer{contract: contract}}, nil
}

// Multisend is an auto generated Go binding around an Ethereum contract.
type Multisend struct {
	MultisendCaller     // Read-only binding to the contract
	MultisendTransactor // Write-only binding to the contract
	MultisendFilterer   // Log filterer for contract events
}

// MultisendCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultisendCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisendTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultisendTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisendFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultisendFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisendSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultisendSession struct {
	Contract     *Multisend        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultisendCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultisendCallerSession struct {
	Contract *MultisendCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MultisendTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultisendTransactorSession struct {
	Contract     *MultisendTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MultisendRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultisendRaw struct {
	Contract *Multisend // Generic contract binding to access the raw methods on
}

// MultisendCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultisendCallerRaw struct {
	Contract *MultisendCaller // Generic read-only contract binding to access the raw methods on
}

// MultisendTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultisendTransactorRaw struct {
	Contract *MultisendTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultisend creates a new instance of Multisend, bound to a specific deployed contract.
func NewMultisend(address common.Address, backend bind.ContractBackend) (*Multisend, error) {
	contract, err := bindMultisend(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multisend{MultisendCaller: MultisendCaller{contract: contract}, MultisendTransactor: MultisendTransactor{contract: contract}, MultisendFilterer: MultisendFilterer{contract: contract}}, nil
}

// NewMultisendCaller creates a new read-only instance of Multisend, bound to a specific deployed contract.
func NewMultisendCaller(address common.Address, caller bind.ContractCaller) (*MultisendCaller, error) {
	contract, err := bindMultisend(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultisendCaller{contract: contract}, nil
}

// NewMultisendTransactor creates a new write-only instance of Multisend, bound to a specific deployed contract.
func NewMultisendTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisendTransactor, error) {
	contract, err := bindMultisend(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultisendTransactor{contract: contract}, nil
}

// NewMultisendFilterer creates a new log filterer instance of Multisend, bound to a specific deployed contract.
func NewMultisendFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisendFilterer, error) {
	contract, err := bindMultisend(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultisendFilterer{contract: contract}, nil
}

// bindMultisend binds a generic wrapper to an already deployed contract.
func bindMultisend(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisendABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisend *MultisendRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Multisend.Contract.MultisendCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisend *MultisendRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisend.Contract.MultisendTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisend *MultisendRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisend.Contract.MultisendTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisend *MultisendCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Multisend.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisend *MultisendTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisend.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisend *MultisendTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisend.Contract.contract.Transact(opts, method, params...)
}

// MultiSend is a paid mutator transaction binding the contract method 0xe3b48f48.
//
// Solidity: function multiSend(address[] recipients, uint256[] amounts, string payload) payable returns()
func (_Multisend *MultisendTransactor) MultiSend(opts *bind.TransactOpts, recipients []common.Address, amounts []*big.Int, payload string) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "multiSend", recipients, amounts, payload)
}

// MultiSend is a paid mutator transaction binding the contract method 0xe3b48f48.
//
// Solidity: function multiSend(address[] recipients, uint256[] amounts, string payload) payable returns()
func (_Multisend *MultisendSession) MultiSend(recipients []common.Address, amounts []*big.Int, payload string) (*types.Transaction, error) {
	return _Multisend.Contract.MultiSend(&_Multisend.TransactOpts, recipients, amounts, payload)
}

// MultiSend is a paid mutator transaction binding the contract method 0xe3b48f48.
//
// Solidity: function multiSend(address[] recipients, uint256[] amounts, string payload) payable returns()
func (_Multisend *MultisendTransactorSession) MultiSend(recipients []common.Address, amounts []*big.Int, payload string) (*types.Transaction, error) {
	return _Multisend.Contract.MultiSend(&_Multisend.TransactOpts, recipients, amounts, payload)
}

// MultisendPayloadIterator is returned from FilterPayload and is used to iterate over the raw logs and unpacked data for Payload events raised by the Multisend contract.
type MultisendPayloadIterator struct {
	Event *MultisendPayload // Event containing the contract specifics and raw log

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
func (it *MultisendPayloadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendPayload)
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
		it.Event = new(MultisendPayload)
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
func (it *MultisendPayloadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendPayloadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendPayload represents a Payload event raised by the Multisend contract.
type MultisendPayload struct {
	Payload string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPayload is a free log retrieval operation binding the contract event 0x53a85291e316c24064ff2c7668d99f35ecbb40ef4e24794ff9d8abe901c7e62c.
//
// Solidity: event Payload(string payload)
func (_Multisend *MultisendFilterer) FilterPayload(opts *bind.FilterOpts) (*MultisendPayloadIterator, error) {

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "Payload")
	if err != nil {
		return nil, err
	}
	return &MultisendPayloadIterator{contract: _Multisend.contract, event: "Payload", logs: logs, sub: sub}, nil
}

// WatchPayload is a free log subscription operation binding the contract event 0x53a85291e316c24064ff2c7668d99f35ecbb40ef4e24794ff9d8abe901c7e62c.
//
// Solidity: event Payload(string payload)
func (_Multisend *MultisendFilterer) WatchPayload(opts *bind.WatchOpts, sink chan<- *MultisendPayload) (event.Subscription, error) {

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "Payload")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendPayload)
				if err := _Multisend.contract.UnpackLog(event, "Payload", log); err != nil {
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

// ParsePayload is a log parse operation binding the contract event 0x53a85291e316c24064ff2c7668d99f35ecbb40ef4e24794ff9d8abe901c7e62c.
//
// Solidity: event Payload(string payload)
func (_Multisend *MultisendFilterer) ParsePayload(log types.Log) (*MultisendPayload, error) {
	event := new(MultisendPayload)
	if err := _Multisend.contract.UnpackLog(event, "Payload", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultisendRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the Multisend contract.
type MultisendRefundIterator struct {
	Event *MultisendRefund // Event containing the contract specifics and raw log

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
func (it *MultisendRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendRefund)
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
		it.Event = new(MultisendRefund)
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
func (it *MultisendRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendRefund represents a Refund event raised by the Multisend contract.
type MultisendRefund struct {
	Refund *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97.
//
// Solidity: event Refund(uint256 refund)
func (_Multisend *MultisendFilterer) FilterRefund(opts *bind.FilterOpts) (*MultisendRefundIterator, error) {

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return &MultisendRefundIterator{contract: _Multisend.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97.
//
// Solidity: event Refund(uint256 refund)
func (_Multisend *MultisendFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *MultisendRefund) (event.Subscription, error) {

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendRefund)
				if err := _Multisend.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0x2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97.
//
// Solidity: event Refund(uint256 refund)
func (_Multisend *MultisendFilterer) ParseRefund(log types.Log) (*MultisendRefund, error) {
	event := new(MultisendRefund)
	if err := _Multisend.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultisendTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Multisend contract.
type MultisendTransferIterator struct {
	Event *MultisendTransfer // Event containing the contract specifics and raw log

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
func (it *MultisendTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendTransfer)
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
		it.Event = new(MultisendTransfer)
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
func (it *MultisendTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendTransfer represents a Transfer event raised by the Multisend contract.
type MultisendTransfer struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address recipient, uint256 amount)
func (_Multisend *MultisendFilterer) FilterTransfer(opts *bind.FilterOpts) (*MultisendTransferIterator, error) {

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &MultisendTransferIterator{contract: _Multisend.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address recipient, uint256 amount)
func (_Multisend *MultisendFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MultisendTransfer) (event.Subscription, error) {

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendTransfer)
				if err := _Multisend.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address recipient, uint256 amount)
func (_Multisend *MultisendFilterer) ParseTransfer(log types.Log) (*MultisendTransfer, error) {
	event := new(MultisendTransfer)
	if err := _Multisend.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
