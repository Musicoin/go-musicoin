## Musicoin Blockcchain in Go

Musicoin Go client(GMC) is a modified blockchain based on Ethereum protocol in Go language.

## Building the source

Building GMC(go-musicoin) requires both a Go and a C compiler.
You can install them using your favorite package manager.
Once the dependencies are installed, run:

    `make gmc`

## Setup cross platform build
    ```
    apt update
    apt install -y docker.io
    apt install -y golang
    git clone https://github.com/Musicoin/go-musicoin.git
    cd go-musicoin
    make gmc-cross
    ```

## Run Musicoin client (GMC)

As we havn't set default dir in code yet, you will still need to set &lt;datadir&gt; manually to prevent from fusing into your local Ethereum data. Please also remove your previous data if there's some generated from previous version of Musicoin client.

`gmc console`

If you want to run with RPC, please add more flags accordingly:

`gmc --rpc --rpcapi="db,eth,net,web3,personal" --rpcport "8545" --rpcaddr "127.0.0.1" --rpccorsdomain "localhost" console`

## Can't find peers to sync?
Musicoin has set up some default nodes that you can try to connect as bootstrap nodes. Once connected, the console will start syncing automatically. In case you can't see syncing after a long time, you may have to add peer(s) manually.

In GMC console, add knowing peer with its enode information:
`> admin.addPeer("{enode info}")`
