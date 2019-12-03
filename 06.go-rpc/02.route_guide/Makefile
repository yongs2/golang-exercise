#!MAKEFILE
# Refer to https://gitlab.com/pantomath-io/demo-grpc/blob/init-makefile/Makefile

SERVER_OUT := "server.bin"
CLIENT_OUT := "client.bin"
PROTO_OUT := "routeguide/route_guide.pb.go"
PKG := "06.go-rpc/02.route_guide"
SERVER_PKG_BUILD := "${PKG}/server"
CLIENT_PKG_BUILD := "${PKG}/client"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all routeguide build_server build_client

all: build_server build_client

routeguide/route_guide.pb.go: routeguide/route_guide.proto
	@protoc -I routeguide/ \
		-I${GOPATH}/src \
		--go_out=plugins=grpc:routeguide \
		routeguide/route_guide.proto

routeguide: routeguide/route_guide.pb.go ## Auto-generate grpc go sources

dep: ## Get the dependencies
	@go get -v -d ./...

build_server: dep routeguide ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

build_client: dep routeguide ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(PROTO_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
