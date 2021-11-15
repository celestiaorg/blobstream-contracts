package util

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Contract struct {
	Name            string
	SourcePath      string
	CompilerVersion string
	Address         common.Address

	ABI []byte
	Bin string
}

type TransactFunc func(opts *bind.TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error)

type BoundContract struct {
	*bind.BoundContract

	transactFn TransactFunc
	client     *ethclient.Client
	address    common.Address
	src        *Contract
	abi        abi.ABI
}

func BindContract(client *ethclient.Client, contract *Contract) (*BoundContract, error) {
	if contract == nil {
		err := errors.New("contract must not be nil")
		return nil, err
	}
	parsedABI, err := abi.JSON(bytes.NewReader(contract.ABI))
	if err != nil {
		err = fmt.Errorf("failed to parse contract ABI: %v", err)
		return nil, err
	}
	bound := &BoundContract{
		BoundContract: bind.NewBoundContract(contract.Address, parsedABI, client, client, client),
		client:        client,

		address: contract.Address,
		abi:     parsedABI,
		src:     contract,
	}
	return bound, nil
}

func (c *BoundContract) SetTransact(fn TransactFunc) {
	c.transactFn = fn
}

func (c *BoundContract) SetClient(client *ethclient.Client) {
	c.client = client
	c.BoundContract = bind.NewBoundContract(
		c.address, c.abi, client, client, client)
}

func (c *BoundContract) Client() *ethclient.Client {
	return c.client
}

func (c *BoundContract) Address() common.Address {
	return c.address
}

func (c *BoundContract) SetAddress(address common.Address) {
	c.address = address
	c.BoundContract = bind.NewBoundContract(
		address, c.abi, c.client, c.client, c.client)
}

func (c *BoundContract) Source() *Contract {
	return c.src
}

func (c *BoundContract) ABI() abi.ABI {
	return c.abi
}

func (c *BoundContract) DeployContract(
	opts *bind.TransactOpts,
	params ...interface{},
) (
	common.Address,
	*types.Transaction,
	error,
) {
	panic("not implemented")
}

func (c *BoundContract) Transact(
	opts *bind.TransactOpts,
	method string,
	params ...interface{},
) (
	*types.Transaction,
	error,
) {
	if c.transactFn == nil {
		return c.BoundContract.Transact(opts, method, params...)
	}
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return nil, err
	}
	return c.transactFn(opts, &c.address, input)
}

func (c *BoundContract) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	if c.transactFn == nil {
		return c.BoundContract.Transfer(opts)
	}
	return c.transactFn(opts, &c.address, nil)
}
