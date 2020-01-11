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

## [apigateway](https://github.com/go-kit/kit/tree/master/examples/apigateway)

- prepare
```sh
go get github.com/hashicorp/consul/api
```

### [usage of Go-Kit Api Gateway](https://medium.com/@jwenz723/go-kit-api-gateway-4bee401e77a2)

1) run addsvc
```sh
cd ${GOPATH}/src/08.go-kit/04.addsvc;
make pb thrift;
cd ${GOPATH}/src/08.go-kit/04.addsvc/cmd/addsvc;
go run addsvc.go -debug.addr :7080 -http-addr :7081 -grpc-addr :7082 -thrift-addr :7083 -jsonrpc-addr :7084
curl -XPOST -d'{"a":"7","b":"4"}' http://localhost:7081/concat
curl -XPOST -d'{"a":7,"b":4}' http://localhost:7081/sum
```

2) run stringsvc
```sh
cd ${GOPATH}/src/08.go-kit/03.stringsvc3;
go run . -listen :8081
curl -XPOST -d'{"s":"mytest"}' http://localhost:8081/count
curl -XPOST -d'{"s":"mytest"}' http://localhost:8081/uppercase
```

3) run consul
```sh
cd $HOME; wget https://releases.hashicorp.com/consul/1.6.2/consul_1.6.2_linux_amd64.zip;
unzip consul_1.6.2_linux_amd64.zip; cp consul /usr/local/bin/;

cd ${GOPATH}/src/08.go-kit/07.apigateway;
mkdir ./consul.d
echo '{"service": {"name": "addsvc", "tags": [], "port": 7082}}' > ./consul.d/addsvc.json
echo '{"service": {"name": "stringsvc", "tags": [], "port": 8081}}' > ./consul.d/stringsvc.json
consul agent -dev -config-dir=./consul.d
```

4) starting api gateway
```sh
cd ${GOPATH}/src/08.go-kit/07.apigateway;
go run main.go -consul.addr :8500

curl -XPOST -d'{"a":"7","b":"4"}' http://localhost:8000/addsvc/concat
curl -XPOST -d'{"a":7,"b":4}' http://localhost:8000/addsvc/sum
curl -XPOST -d'{"s":"mytest"}' http://localhost:8000/stringsvc/count
curl -XPOST -d'{"s":"mytest"}' http://localhost:8000/stringsvc/uppercase
```

## [Example : napodate](https://dev.to/napolux/how-to-write-a-microservice-in-go-with-go-kit-a66)

- run napodate service
```sh
cd ${GOPATH}/src/08.go-kit/08.napodate/cmd;
go run main.go
```

- test napodate service
```sh
curl http://localhost:8080/get
curl http://localhost:8080/status
curl -XPOST -d '{"date":"32/12/2020"}' http://localhost:8080/validate
curl -XPOST -d '{"date":"12/12/2021"}' http://localhost:8080/validate
```

## [kitlog](https://opencensus.io/integrations/go_kit/)

- install kitgen binaray
```sh
cd ${GOPATH}/src/github.com/go-kit/kit/cmd/kitgen
go get .
go install
kitgen -h
```

- autogenerate endpoints, http, service with kitgen
```sh
cd ${GOPATH}/src/08.go-kit/09.kitloc;
mkdir hello; cd hello
kitgen ../service.go
```
