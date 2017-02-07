## Musicoin Blockcchain in Go

Modified version of the Ethereum protocol in Go

## Building the source

Building GMC(go-musicoin) requires both a Go and a C compiler.
You can install them using your favorite package manager.
Once the dependencies are installed, run

    make gmc

## Setup cross platform build
    ```
    apt update
    apt install -y docker.io
    apt install -y golang
    git clone https://github.com/Musicoin/go-musicoin.git
    cd go-musicoin
    make gmc-cross
    ```

## Run Musicoin client (gmc)

As we havn't set default dir in code yet, you will still need to set &lt;datadir&gt; manually to prevent from fusing into your local Ethereum data. Please also remove your previous data if there's some generated from previous version of Musicoin client.

`gmc --identity Musicoin --datadir ~/.musicoin console`

If you want to run with RPC, please add more flags accordingly:

`gmc --identity Musicoin --datadir ~/.musicoin --rpc --rpcapi="db,eth,net,web3,personal" --rpcport "8545" --rpcaddr "127.0.0.1" --rpccorsdomain "localhost" console`

## Can't find peers?
Musicoin has set up some default nodes that you can try to connect default boostrap nodes:

In GMC console(either of the following 2 nodes):
> admin.addPeer("{enode info}")

"enode://ba2f6409f9894c12f5aad3471b9c4a2e7999b246af775c39f99d85b020cfc95d0b0dc6dd0985895bb2c4149cb45e4a3c17f7585be326cec293176bf81802a987@104.196.160.105:30303"

"enode://58703aae65e704576a95079b9814cf98f4e20294d9d21f444540c1b83d98353f1228caa7eec3fabfe52c094deffd0e33f0004e8631415a4f0d7c6fba6788bec0@104.197.56.66:30303"
