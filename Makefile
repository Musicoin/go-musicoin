# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gmc android ios gmc-cross swarm evm all test clean
.PHONY: gmc-linux gmc-linux-386 gmc-linux-amd64 gmc-linux-mips64 gmc-linux-mips64le
.PHONY: gmc-linux-arm gmc-linux-arm-5 gmc-linux-arm-6 gmc-linux-arm-7 gmc-linux-arm64
.PHONY: gmc-darwin gmc-darwin-386 gmc-darwin-amd64
.PHONY: gmc-windows gmc-windows-386 gmc-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

gmc:
	build/env.sh go run build/ci.go install ./cmd/gmc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gmc\" to launch gmc."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gmc.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/GMC.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./cmd/abigen

# Cross Compilation Targets (xgo)

gmc-cross: gmc-linux gmc-darwin gmc-windows gmc-android gmc-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gmc-*

gmc-linux: gmc-linux-386 gmc-linux-amd64 gmc-linux-arm gmc-linux-mips64 gmc-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-*

gmc-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gmc
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep 386

gmc-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gmc
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep amd64

gmc-linux-arm: gmc-linux-arm-5 gmc-linux-arm-6 gmc-linux-arm-7 gmc-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep arm

gmc-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gmc
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep arm-5

gmc-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gmc
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep arm-6

gmc-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gmc
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep arm-7

gmc-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gmc
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep arm64

gmc-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gmc
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep mips

gmc-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gmc
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep mipsle

gmc-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gmc
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep mips64

gmc-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gmc
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep mips64le

gmc-darwin: gmc-darwin-386 gmc-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gmc-darwin-*

gmc-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gmc
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-darwin-* | grep 386

gmc-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gmc
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-darwin-* | grep amd64

gmc-windows: gmc-windows-386 gmc-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gmc-windows-*

gmc-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gmc
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-windows-* | grep 386

gmc-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gmc
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-windows-* | grep amd64
