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
	"enode://7c62387a4f0a02603b6dcbc7d1223521dad6837e6f11d9da84ac2a4802fe855531c26ac3b1bfb895c15d1e3ec5882661bb95c4b98d9124407e72881e86b8c1b1@104.196.1.188:30303",
	"enode://05592bc5d209c2b0eee10ad47cf3324898d6d8a34e9ac0746431d52eb5fc5d415e9bd42ecb42f44b29955dafdb760fd84817e23c3e3bc041783d6315be6aa996@35.243.225.31:30303",
	"enode://5812ed9c80cbe08bad47ebd01d754116b23d553e9154b30177504efd57db31c98ec2bbc69f63194799da3d646c11a75b56b60c75509f202d1981de4ba31dc97f@35.237.165.10:30303",
	"enode://ad5288a8082141ecb6f5b6687835fddac31ec1fd94ca904c21438fec65b4b1d5f4481a30d2e93677acec331ad445739cf6d271adb7c89e26955db6bd64885c04@35.185.67.35:30303",
	"enode://b58c0c71f08864c0cf7fa9dea2c4cbefae5ae7a36cc30d286603b24982d25f3ccc056b589119324c51768fc2054b8c529ecf682e06e1e9980170b93ff194ed7a@132.148.132.9:30303",
	"enode://001f2b5d2f73884e4a7cd8e47307ca47445401e817e41d8a4273c5e0e959ea67bfe110372810c1e5fa990e8342c963209054290e7e7e7dff749b10100c333c8b@211.149.192.7:54615",
	"enode://d302f52c8789ad87ee528f1431a67f1aa646c9bec17babb4665dfb3d61de5b9119a70aa77b2147a5f28854092ba09769323c1c552a6ac6f6a34cbcf767e2d2fe@158.69.248.48:30303",
	"enode://005dca0dfc195fc3729205299a90e6ce4938b8b3ec420c985ebb5852c7b5cfae11f67b0d3f2de47b21c7d067588542dfecc516f28678838e378517ace3845bc5@180.167.190.204:42228",
	"enode://0084e6cd8a63e33b6d283018c7a4d72e6c450a8a845958f65a1f46520e955bc07a410bbdb476c3d927d2dd9269fa498a3e7af613525d1850b3be5131b153f6ff@119.145.254.18:46582",
	"enode://c72564bce8331ae298fb8ece113a456e3927d7e5989c2be3e445678b3600579f722410ef9bbfe339335d676af77343cb21b5b1703b7bebc32be85fce937a2220@191.252.185.71:30303",
	"enode://01326803b85a8199ff83adac5e46808f90d617f6dccce829885dd698963d5cedd2c46685322f237e8c7b95dbb6971ce6dbfc8e9bb3991473fa5bb8bca67d0b41@111.230.171.185:38288",
	"enode://e3ae4d25ee64791ff98bf17c37acf90933359f2505c00f65c84f6863231a32a94153cadb0a462e428f18f35ded6bd91cd91033d26576a28558c22678be9cfaee@5.63.158.137:35555",
	"enode://a979fb575495b8d6db44f750317d0f4622bf4c2aa3365d6af7c284339968eef29b69ad0dce72a4d8db5ebb4968de0e3bec910127f134779fbcb0cb6d3331163c@52.16.188.185:30303",
	"enode://3f1d12044546b76342d59d4a05532c14b85aa669704bfe1f864fe079415aa2c02d743e03218e57a33fb94523adb54032871a6c51b2cc5514cb7c7e35b3ed0a99@13.93.211.84:30303",
	"enode://78de8a0916848093c73790ead81d1928bec737d565119932b98c6b100d944b7a95e94f847f689fc723399d2e31129d182f7ef3863f2b4c820abbf3ab2722344d@191.235.84.50:30303",
	"enode://158f8aab45f6d19c6cbf4a089c2670541a8da11978a2f90dbf6a502a4a3bab80d288afdbeb7ec0ef6d92de563767f3b1ea9e8e334ca711e9f8e2df5a0385e8e6@13.75.154.138:30303",
	"enode://1118980bf48b0a3640bdba04e0fe78b1add18e1cd99bf22d53daac1fd9972ad650df52176e7c7d89d1114cfef2bc23a2959aa54998a46afcf7d91809f0855082@52.74.57.123:30303",
	"enode://979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9@5.1.83.226:30303",
	"enode://d302f52c8789ad87ee528f1431a67f1aa646c9bec17babb4665dfb3d61de5b9119a70aa77b2147a5f28854092ba09769323c1c552a6ac6f6a34cbcf767e2d2fe@158.69.248.48:30303",
	"enode://c72564bce8331ae298fb8ece113a456e3927d7e5989c2be3e445678b3600579f722410ef9bbfe339335d676af77343cb21b5b1703b7bebc32be85fce937a2220@191.252.185.71:30303",
	"enode://e3ae4d25ee64791ff98bf17c37acf90933359f2505c00f65c84f6863231a32a94153cadb0a462e428f18f35ded6bd91cd91033d26576a28558c22678be9cfaee@5.63.158.137:35555",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://6ce05930c72abc632c58e2e4324f7c7ea478cec0ed4fa2528982cf34483094e9cbc9216e7aa349691242576d552a2a56aaeae426c5303ded677ce455ba1acd9d@13.84.180.240:30303", // US-TX
	"enode://20c9ad97c081d63397d7b685a412227a40e23c8bdc6688c6f37e97cfbc22d2b4d1db1510d8f61e6a8866ad7f0e17c02b14182d37ea7c3c8b9c2683aeb6b733a1@52.169.14.227:30303", // IE
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303?discport=30304", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303?discport=30304",  // INFURA
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30305",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30308",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30309",
}
