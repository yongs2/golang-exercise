# Go-kit Examples

- [Go-kit Example](https://gokit.io/examples/)

## stringsvc

### [stringsvc1](https://github.com/go-kit/kit/tree/master/examples/stringsvc1)

```sh
go run 01.stringsvc1/main.go

curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
```

### [stringsvc2](https://github.com/go-kit/kit/tree/master/examples/stringsvc2)

```sh
go run 02.stringsvc2/*.go

curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
```

### [stringsvc3](https://github.com/go-kit/kit/tree/master/examples/stringsvc3)

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

## [addsvc](https://github.com/go-kit/kit/tree/master/examples/addsvc)

- [install apache-thrift](https://thrift.apache.org/)

```sh
apt-get -y install flex bison
wget https://github.com/apache/thrift/archive/v0.13.0.tar.gz
tar -zxvf v0.13.0.tar.gz
cd thrift-0.13.0/
./bootstrap.sh
./configure  --without-qt4 --without-qt5 --without-c_glib --without-csharp --without-java --without-erlang --without-nodejs --without-lua --without-python --without-perl --without-php --without-php_extension --without-dart --without-ruby --without-haskell --without-rs --without-cl --without-haxe --without-dotnetcore --without-d
make; make install
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

## [profilesvc](https://github.com/go-kit/kit/tree/master/examples/profilesvc)

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

## [shipping](https://github.com/go-kit/kit/tree/master/examples/shipping)
- reference : [GoDDD](https://github.com/marcusolsson/goddd)
- test booking
  - Check out the sample cargos
  ```sh
  curl -X GET http://localhost:8080/booking/v1/cargos
  ```

  - Book new cargo
  ```sh
  curl -X POST -d '{"origin": "SESTO", "destination": "FIHEL", "arrival_deadline": "2016-03-21T19:50:24Z"}' -H "Content-Type: application/json" http://localhost:8080/booking/v1/cargos 
  curl -X POST -d '{"origin": "SESTO", "destination": "CNHKG", "arrival_deadline": "2021-01-19T09:28:00Z"}' -H "Content-Type: application/json" http://localhost:8080/booking/v1/cargos
  ```

  - Get cargo ABC123
  ```sh
  curl -X GET http://localhost:8080/booking/v1/cargos/ABC123
  ```

  - Request possible routes for sample cargo ABC123
  ```sh
  curl -X GET http://localhost:8080/booking/v1/cargos/ABC123/request_routes
  ```

  - change destination
  ```sh
  curl -d '{}' -X POST http://localhost:8080/booking/v1/cargos/ABC123/change_destination
  ```

  - Get location code
  ```sh
  curl -X GET http://localhost:8080/booking/v1/locations
  ```

  - Register incident
  ```sh
  curl -d '{"completion_time":"2016-03-21T19:50:24Z", "tracking_id":"ABC123", "voyage":"V100", "location":"SESTO", "event_type":"Receive"}' -X POST http://localhost:8080/handling/v1/incidents
  ```

  - tracking
  ```sh
  curl -X GET http://localhost:8080/tracking/v1/cargos/ABC123
  ```

## apigateway

- prepare
```sh
go get github.com/hashicorp/consul/api
```
