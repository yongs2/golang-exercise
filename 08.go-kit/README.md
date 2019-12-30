# Go-kit Examples

- [Go-kit Example](https://gokit.io/examples/)

## stringsvc

### stringsvc1

```sh
go run 01.stringsvc1/main.go

curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
```

### stringsvc2

```sh
go run 02.stringsvc2/*.go

curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
```

### stringsvc3

- add extra package

```sh
go get github.com/streadway/handy/breaker
go get github.com/afex/hystrix-go/hystrix
go get golang.org/x/time/rate
```

- run server and proxy

```sh
go run 03.stringsvc3/*.go -listen=:8001 &
go run 03.stringsvc3/*.go -listen=:8002 &
go run 03.stringsvc3/*.go -listen=:8003 &
go run 03.stringsvc3/*.go -listen=:8080 -proxy=localhost:8001,localhost:8002,localhost:8003
```

- run client

```sh
for s in foo bar baz ; do curl -d"{\"s\":\"$s\"}" localhost:8080/uppercase ; done
for s in foo bar baz ; do curl -d"{\"s\":\"$s\"}" localhost:8080/count ; done
curl localhost:8080/metrics
```

## addsvc

- [install apache-thrift](https://thrift.apache.org/)

```sh
apt-get -y install flex bison
wget https://github.com/apache/thrift/archive/v0.13.0.tar.gz
tar -zxvf v0.13.0.tar.gz
cd thrift-0.13.0/
./bootstrap.sh
./configure  --without-qt4 --without-qt5 --without-c_glib --without-csharp --without-java --without-erlang --without-nodejs --without-lua --without-python --without-perl --without-php --without-php_extension --without-dart --without-ruby --without-haskell --without-rs --without-cl --without-haxe --without-dotnetcore --without-d
```

- build

```sh
make pb thrift
make
```

- run server.bin for test

```sh
./server.bin 
```

- run client.bin for test

```sh
./client.bin -http-addr=localhost:8081 -method=sum 1 2
./client.bin -http-addr=localhost:8081 -method=concat 1 2
./client.bin -grpc-addr=localhost:8082 -method=sum 1 2
./client.bin -grpc-addr=localhost:8082 -method=concat 1 2
./client.bin -thrift-addr=localhost:8083 -method=sum 1 2
./client.bin -thrift-addr=localhost:8083 -method=concat 1 2
./client.bin -jsonrpc-addr=localhost:8084 -method=sum 1 2
./client.bin -jsonrpc-addr=localhost:8084 -method=concat 1 2
```

## profilesvc

- run profilesvc

```sh
go run ./cmd/profilesvc/main.go
```

- create a profile

```sh
curl -d '{"id":"1234","Name":"Go Kit"}' -H "Content-Type: application/json" -X POST http://localhost:8080/profiles/
curl localhost:8080/profiles/1234
curl -d '{"id": "1", "location": "location1"}' -H "Content-Type: application/json" -X POST http://localhost:8080/profiles/1234/addresses/
curl -X GET localhost:8080/profiles/1234/addresses/
curl -X DELETE localhost:8080/profiles/1234/addresses/1
```

## shipping

## apigateway