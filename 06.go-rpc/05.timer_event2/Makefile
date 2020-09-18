#!MAKEFILE
# Refer to https://gitlab.com/pantomath-io/demo-grpc/blob/init-makefile/Makefile

SERVER_OUT := server.exe
CLIENT_OUT := client.exe
PROTO_OUT := api/proto/eventTimer.pb.go
PKG := "06.go-rpc/05.timer_event2"
SERVER_PKG_BUILD := "${PKG}/cmd/server"
CLIENT_PKG_BUILD := "${PKG}/cmd/client"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all eventTimerApi build_server build_client

all: build_server build_client

api/proto/eventTimer.pb.go: api/proto/eventTimer.proto
	@protoc --proto_path=api/proto/ \
		--go_opt=paths=source_relative \
		--go_out=plugins=grpc:api/proto \
		api/proto/eventTimer.proto

eventTimerApi: api/proto/eventTimer.pb.go ## Auto-generate grpc go sources

dep: ## Get the dependencies
	@go get -v -d ./...

build_server: dep eventTimerApi ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

build_client: dep eventTimerApi ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(PROTO_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'