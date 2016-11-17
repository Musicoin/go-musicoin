## Musicoin Blockcchain in Go

Modified version of the Ethereum protocol in Go

## Building the source

Building geth requires both a Go and a C compiler.
You can install them using your favourite package manager.
Once the dependencies are installed, run

    make geth

or, to build the full suite of utilities:

    make all

## Run Musicoin version of geth

As we havn't set default dir in code yet, you will still need to set &lt;datadir&gt; manually to prevent from fusing into your local Ethereum data. Please also remove your previous data if there's some generated from previous verison of Musicoin client.

`geth --identity Musicoin --networkid 55313716 --datadir ~/.musicoin console`

If you want to run with RPC, please add more flags accordingly:

`geth --identity Musicoin --networkid 55313716 --datadir ~/.musicoin --rpc --rpcapi="db,eth,net,web3,personal" --rpcport "8545" --rpcaddr "127.0.0.1" --rpccorsdomain "localhost" console`

## Can't find peers?
Musicoin has set up some default nodes that you can try to connect default boostrap nodes:

In geth console(either of the following 2 nodes): 
> admin.addPeer("{enode info}")

node 1:
enode://d383ed7364a61b49c8532cd96ed95bfaee49388779f04dd5baaf63769817595c6c2c16324d11598b4e90c1069e467a7aaf44f10ea20ae85da31752de60e0fc47@104.197.75.174:30303

nodoe 2: enode://b5cd2f4e925f5eaa026e3965a89888e156d9e190f7d16cd685687c3ff0abbafe0a29a45d1ea2a0096f9454a07bdabf38f2fce3df108a33ce5620055016d7abac@104.198.70.194:30303
