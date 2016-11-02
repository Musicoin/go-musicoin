// Copyright 2015 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package utils

import "github.com/ethereum/go-ethereum/p2p/discover"

// FrontierBootNodes are the enode URLs of the P2P bootstrap nodes running on
// the Frontier network.
var FrontierBootNodes = []*discover.Node{
	// ETH/DEV Go Bootnodes
	discover.MustParseNode("enode://96eb532a30efaf5e968d837207da5b3545316966ac06d30aad1a74d7e68c787e8405f6cf5dc218a5528cb2d619435bdaeb9f8cb8955df9ab9dd12ef085fe9009@104.198.70.194:30303"), // IE
	discover.MustParseNode("enode://6c262c6e95f65ab0199fdd4832eab17512372a145fb057602fa322d72d331b384881cc1a0bd6e690367bc309cb38f5dbaafe9e93cee26246d83ef23f22379bef@104.197.75.174:30303"),  // BR

	// ETH/DEV Cpp Bootnodes
	//discover.MustParseNode("enode://979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9@5.1.83.226:30303"),
}

// TestNetBootNodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestNetBootNodes = []*discover.Node{
	// ETH/DEV Go Bootnodes
	discover.MustParseNode("enode://e4533109cc9bd7604e4ff6c095f7a1d807e15b38e9bfeb05d3b7c423ba86af0a9e89abbf40bd9dde4250fef114cd09270fa4e224cbeef8b7bf05a51e8260d6b8@94.242.229.4:40404"),
	discover.MustParseNode("enode://8c336ee6f03e99613ad21274f269479bf4413fb294d697ef15ab897598afb931f56beb8e97af530aee20ce2bcba5776f4a312bc168545de4d43736992c814592@94.242.229.203:30303"),

	// ETH/DEV Cpp Bootnodes
}
