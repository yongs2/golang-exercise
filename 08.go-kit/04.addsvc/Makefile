#!MAKEFILE
# Refer to https://gitlab.com/pantomath-io/demo-grpc/blob/init-makefile/Makefile

SERVER_OUT := "server.bin"
CLIENT_OUT := "client.bin"
PROTO_OUT := "pb/addsvc.pb.go"
PKG := "08.go-kit/04.addsvc"
SERVER_PKG_BUILD := "${PKG}/cmd/addsvc"
CLIENT_PKG_BUILD := "${PKG}/cmd/addcli"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all pb build_server build_client

all: build_server build_client

pb/addsvc.pb.go: pb/addsvc.proto
	@protoc -I pb/ \
		-I${GOPATH}/src \
		--go_out=plugins=grpc:pb \
		pb/addsvc.proto

pb: pb/addsvc.pb.go ## Auto-generate grpc go sources

dep: ## Get the dependencies
	@go get -v -d ./...

build_server: dep pb ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

build_client: dep pb ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(PROTO_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
