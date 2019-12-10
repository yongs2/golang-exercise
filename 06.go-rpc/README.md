# gRPC Examples

## 01.helloworld

### 1. Install protobuf compiler

```sh
apt-get update; apt-get -y install unzip;
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

### 3. [gomock](https://github.com/golang/mock)

```sh
go get github.com/golang/mock/mockgen
cd /go/src/06.go-rpc/01.helloworld
mkdir -p mock_helloworld
mockgen -destination=mock_helloworld/hw_mock.go -package=mock_helloworld -source=helloworld/helloworld.pb.go
```

### 4. generate cpp code

- grpc 다운로드 

```sh
cd ~/;
git clone -b $(curl -L https://grpc.io/release) https://github.com/grpc/grpc
cd ~/grpc;
git submodule update --init
make;
make install;
ls -la /usr/local/include
```

- make 에러 발생시 Makefile 에서 -Werror 를 -Wno-tautological-compare -Wno-class-memaccess -Wno-ignored-qualifiers 으로 변경

- protobuf 설치

```sh
cd third-party/protobuf;
make; make install;
```

- greeter_client 빌드

```sh
cd cpp_helloworld;
make greeter_client
```

- proto 변환 시험

```sh
cd /go/src/06.go-rpc/01.helloworld/
protoc -I helloworld helloworld/helloworld.proto --cpp_out=./helloworld
```

## 02. [route_guide](https://github.com/grpc/grpc-go/tree/master/examples/route_guide)

### 1. Define a service in a .proto file(route_guide.proto).
### 2. Generate server and client code using the protocol buffer compiler.

- generate go code

```sh
cd /go/src/06.go-rpc/02.routeguide
protoc -I routeguide routeguide/route_guide.proto --go_out=plugins=grpc:routeguide
```

### 3. Use the Go gRPC API to write a simple client and server for your service

go generate 06.go-rpc/02.route_guide/routeguide

### 4. go-mock

```sh
go get github.com/golang/mock/mockgen
cd /go/src/06.go-rpc/02.route_guide
mkdir -p mock_routeguide

mockgen -destination=mock_routeguide/rg_mock.go -package=mock_routeguide -source=routeguide/route_guide.pb.go
```

## 03. features/proto/echo

### 1. Define a service in a .proto file(echo/echo.proto).
### 2. Generate server and client code using the protocol buffer compiler.

- generate go code

```sh
cd /go/src/06.go-rpc/03.features/proto
go generate 06.go-rpc/03.features/proto
```

- or run protoc
```sh
cd /go/src/06.go-rpc/03.features/proto
protoc -I ./echo ./echo/echo.proto --go_out=plugins=grpc,paths=source_relative:./echo
```

### features/authentication

- golang.org/x/oauth2 를 포함시켰는데도, undefined: metadata.OnGCE 라고 표시되면, 
- https://github.com/golang/oauth2 를 참조하여, metadata 참조하는 부분을 cloud.google.com/go/compute/metadata 로 수정

### features/compression

- UseCompressor 로 gzip 으로 전송 데이터를 압축시킴
- res, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: msg}, grpc.UseCompressor(gzip.Name))

### features/deadline
### features/encryption/TLS
### features/errors
### features/interceptor
### features/load_balancing
### features/metadata
### features/multiplex
### features/name_resolving