// Copyright 2015 The go-ethereum Authors
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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://09fcd36d553044c8b499b9b9e13a228ffd99572c513f77073d41f009717c464cd4399c0e665d6aff1590324254ee4e698b2b2533b1998dd04d896b9d6aff7895@35.185.67.35:30303",
	"enode://89e51a34770a0badf8ea18c4c4d2c361cde707abd60031d99b1ab3010363e1898230a516ddb37d974af8d8db1b322779d7fe0caae0617bed4924d1b4968cf92b@35.231.48.142:30303",
	"enode://b58c0c71f08864c0cf7fa9dea2c4cbefae5ae7a36cc30d286603b24982d25f3ccc056b589119324c51768fc2054b8c529ecf682e06e1e9980170b93ff194ed7a@132.148.132.9:30303",
	"enode://d302f52c8789ad87ee528f1431a67f1aa646c9bec17babb4665dfb3d61de5b9119a70aa77b2147a5f28854092ba09769323c1c552a6ac6f6a34cbcf767e2d2fe@158.69.248.48:30303",
	"enode://c72564bce8331ae298fb8ece113a456e3927d7e5989c2be3e445678b3600579f722410ef9bbfe339335d676af77343cb21b5b1703b7bebc32be85fce937a2220@191.252.185.71:30303",
	"enode://e3ae4d25ee64791ff98bf17c37acf90933359f2505c00f65c84f6863231a32a94153cadb0a462e428f18f35ded6bd91cd91033d26576a28558c22678be9cfaee@5.63.158.137:35555",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// PoW test network.
var TestnetBootnodes = []string{
	"enode://6ce05930c72abc632c58e2e4324f7c7ea478cec0ed4fa2528982cf34483094e9cbc9216e7aa349691242576d552a2a56aaeae426c5303ded677ce455ba1acd9d@13.84.180.240:30303", // US-TX
	"enode://20c9ad97c081d63397d7b685a412227a40e23c8bdc6688c6f37e97cfbc22d2b4d1db1510d8f61e6a8866ad7f0e17c02b14182d37ea7c3c8b9c2683aeb6b733a1@52.169.14.227:30303", // IE
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// PoA test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}
