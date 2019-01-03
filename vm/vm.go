// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"math/big"

	"github.com/LambdaIM/lambda-libs/common"
)

// NewEVMContext creates a new context for use in the E
func NewEVMContext(
	blockCoinbase *common.Address,
	blockFromAddr common.Address,
	blockHash common.Hash,
	blockNumber *big.Int,
	blockTime *big.Int,
	blockDifficulty *big.Int,
	gasLimit uint64,
	gasPrice *big.Int,
	orderFunc OrderFunc) Context {

	return Context{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		GetHash:     GetHashFn(blockHash),
		Order:       orderFunc,
		Origin:      blockFromAddr,
		Coinbase:    *blockCoinbase,
		BlockNumber: new(big.Int).Set(blockNumber),
		Time:        new(big.Int).Set(blockTime),
		Difficulty:  new(big.Int).Set(blockDifficulty),
		GasLimit:    gasLimit,
		GasPrice:    new(big.Int).Set(gasPrice),
	}
}

// GetHashFn returns a GetHashFunc which retrieves header hashes by number
func GetHashFn(blockHash common.Hash) func(n uint64) common.Hash {
	return func(n uint64) common.Hash {
		return blockHash
	}
}

// CanTransfer checks whether there are enough funds in the address' account to make a transfer.
// This does not take the necessary gas in to account to make the transfer valid.
func CanTransfer(db StateDB, addr common.Address, amount *big.Int) bool {
	return db.GetBalance(addr).Cmp(amount) >= 0
}

// Transfer subtracts amount from sender and adds amount to recipient using the given Db
func Transfer(db StateDB, sender, recipient common.Address, amount *big.Int) {
	db.SubBalance(sender, amount)
	db.AddBalance(recipient, amount)
}
