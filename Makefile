# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gmc gmc-cross evm all test clean
.PHONY: gmc-linux gmc-linux-386 gmc-linux-amd64 gmc-linux-mips64 gmc-linux-mips64le
.PHONY: gmc-linux-arm gmc-linux-arm-5 gmc-linux-arm-6 gmc-linux-arm-7 gmc-linux-arm64
.PHONY: gmc-darwin gmc-darwin-386 gmc-darwin-amd64
.PHONY: gmc-windows gmc-windows-386 gmc-windows-amd64
.PHONY: gmc-android gmc-ios

GOBIN = build/bin
GO ?= latest

gmc:
	build/env.sh go run build/ci.go install ./cmd/gmc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gmc\" to launch musicoin."

evm:
	build/env.sh go run build/ci.go install ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm to start the evm."

all:
	build/env.sh go run build/ci.go install

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ Godeps/_workspace/pkg $(GOBIN)/*

# Cross Compilation Targets (xgo)

gmc-cross: gmc-linux gmc-darwin gmc-windows gmc-android gmc-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gmc-*

gmc-linux: gmc-linux-386 gmc-linux-amd64 gmc-linux-arm gmc-linux-mips64 gmc-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-*

gmc-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=linux/386 -v ./cmd/gmc
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep 386

gmc-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=linux/amd64 -v ./cmd/gmc
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep amd64

gmc-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=linux/mips64 -v ./cmd/gmc
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep mips64

gmc-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=linux/mips64le -v ./cmd/gmc
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gmc-linux-* | grep mips64le

gmc-darwin: gmc-darwin-386 gmc-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gmc-darwin-*

gmc-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=darwin/386 -v ./cmd/gmc
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-darwin-* | grep 386

gmc-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=darwin/amd64 -v ./cmd/gmc
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-darwin-* | grep amd64

gmc-windows: gmc-windows-386 gmc-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gmc-windows-*

gmc-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=windows/386 -v ./cmd/gmc
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-windows-* | grep 386

gmc-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=windows/amd64 -v ./cmd/gmc
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gmc-windows-* | grep amd64
