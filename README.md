## Musicoin Blockcchain in Go

Musicoin Go client(GMC) is a modified blockchain based on Ethereum protocol in Go language.

## Building the source

Building GMC(go-musicoin) requires both a Go and a C compiler.
You can install them using your favorite package manager.
Once the dependencies are installed, run:

`make gmc`

## Known issues with Go1.8+

While compiling with Go1.8+, you will notice the following error:
```
build/env.sh go run build/ci.go install ./cmd/gmc
unexpected directory layout:
	import path:
  /go-musicoin/build/_workspace/src/github.com/ethereum/go-ethereum/internal/build
	root:
  /go-musicoin/build/_workspace/src
	dir:
  /go-musicoin/build/_workspace/src/github.com/ethereum/go-ethereum/internal/build
	expand root: /go-musicoin/build/_workspace/src
	expand dir: /go-musicoin/internal/build
	separator: /
make: *** [gmc] Error 1
```
This is because of the unique directory structure that gmc follows which would not compile with the strict regulations imposed by Goo1.8 and above. In order to solve this, please re-compile using Go1.7 or lower. A successful build will create a `build/` subdirectory under which the `gmc` binary will be located.

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

`gmc console`

A common error is that `gmc` is not executable. Please do `chmod +x <path-to-gmc>` to avoid this.

If you want to run with RPC, please add more flags accordingly:

`gmc --rpc --rpcapi="db,eth,net,web3,personal" --rpcport "8545" --rpcaddr "127.0.0.1" --rpccorsdomain "localhost" console`

## Can't find peers to sync?
Musicoin has set up some default nodes that you can try to connect as bootstrap nodes. Once connected, the console will start syncing automatically. In case you can't see syncing after a long time, you may have to add peer(s) manually.

In GMC console, add knowing peer with its enode information:

`> admin.addPeer("{enode info}")`
