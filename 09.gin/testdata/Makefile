#!MAKEFILE
# Refer to https://gitlab.com/pantomath-io/demo-grpc/blob/init-makefile/Makefile

PROTO_OUT := "protoexample/test.pb.go"
PKG := "09.gin/testdata"
#PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all  

all: protoexample

protoexample/test.pb.go: protoexample/test.proto
	@protoc -I protoexample/ \
		-I${GOPATH}/src \
		--go_out=plugins=grpc:protoexample \
		protoexample/test.proto

protoexample: protoexample/test.pb.go ## Auto-generate grpc go sources

clean: ## Remove previous builds
	@rm $(PROTO_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
