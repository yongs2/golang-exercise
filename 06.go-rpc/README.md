# gRPC Examples

## 01.helloworld

### 1. Install protobuf compiler

```sh
cd ~/
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.0/protoc-3.11.0-linux-x86_64.zip
unzip protoc-3.11.0-linux-x86_64.zip -d ~/protoc-3.11.0-linux-x86_64
cp ~/protoc-3.11.0-linux-x86_64/bin/protoc /go/bin/

go get -d -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go
ls -al $GOPATH/bin/
```

### 2. generate go code

```sh
cd /go/src/06.go-rpc/01.helloworld
protoc -I helloworld helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```
