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

```sh
go get github.com/streadway/handy/breaker
go get github.com/afex/hystrix-go/hystrix
go get golang.org/x/time/rate
```

```sh
go run 02.stringsvc3/*.go

curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
```

## addsvc

## profilesvc

## shipping

## apigateway