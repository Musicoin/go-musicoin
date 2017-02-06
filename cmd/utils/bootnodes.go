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
	//discover.MustParseNode("enode://f943403fae29cc82938b4d100ca3fb27f8b123a51d2c9f4dfbada6ae8dc536544552cb8d53e2f81c28951ec33c9f2bbb99be36fb0079cf242a7412cc9232e26e@104.196.233.5:30301"), // IE
	discover.MustParseNode("enode://c9d07b236fc02eb2005ff8121eebc9b5a2820ce68590038796b51900f43a09c7f165026c3ac30ac1aaf5b5fa5febc96dbb10204cd0ff87e8dfab3327013357cb@35.185.39.34:30301"),  // BR
  discover.MustParseNode("enode://e17d7d4d383cff3d95618ac12166be015abf18fd949d3b207f671daeed6a4e872dce9e92a96015916f044ecfa0f9c0a71327f89ec89e8d9a65baf4a3ccd24d0b@104.199.167.97:30301"),
	discover.MustParseNode("enode://2ed2f2d86f5eea288af77cd634b63c7a3e99c78506bc8c95b87e31a2794e502a2f5c7048ed3199bc13b85a7a6eda4f289be501cb69f3bfd68794e3c397e61039@104.155.43.197:30301")
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
