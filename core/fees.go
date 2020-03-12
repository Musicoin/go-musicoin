// Copyright 2014 The Go-Ethereum Authors
// This file is part of the Go-Ethereum library.
//
// Copyright 2017 The Go-Musicoin Authors
// This file is part of the Go-Musicoin library.
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

package core

import (
	"math/big"
)

var BlockReward *big.Int = new(big.Int).Mul(big.NewInt(314), big.NewInt(1e+18))
var NewBlockReward *big.Int = new(big.Int).Mul(big.NewInt(250), big.NewInt(1e+18))
var NMBlockReward *big.Int = new(big.Int).Mul(big.NewInt(0), big.NewInt(1e+18))
var QTBlockReward *big.Int = new(big.Int).Mul(big.NewInt(50), big.NewInt(1e+18))
var UBIReward *big.Int = new(big.Int).Mul(big.NewInt(50), big.NewInt(1e+18))
var DevReward *big.Int = new(big.Int).Mul(big.NewInt(14), big.NewInt(1e+18))
