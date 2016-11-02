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
