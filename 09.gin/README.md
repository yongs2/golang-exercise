# [Gin Web framework](https://gin-gonic.com/)

## [github](https://github.com/gin-gonic/gin)

### 01.QuickStart

- run example
```sh
go get github.com/gin-gonic/gin
cd 01.QuickStart;
go run example.go
curl http://localhost:8080/ping
```

### 02.Router

- Using GET, POST, PUT, PATCH, DELETE and OPTIONS
- run example
```sh
go get github.com/gin-gonic/gin
cd 02.Router;
go run example.go

curl -X GET localhost:8080/someGet
curl -X POST localhost:8080/somePost
curl -X PUT localhost:8080/somePut
curl -X DELETE localhost:8080/someDelete
curl -X PATCH localhost:8080/somePatch
curl -X HEAD localhost:8080/someHead
curl -X OPTIONS localhost:8080/someOptions
```

### 03.Parameters

- Parameters in path
- run example
```sh
go get github.com/gin-gonic/gin
cd 03.Parameters
go run example.go

curl -X GET localhost:8080/user/john
curl -X GET localhost:8080/user/john/run
curl -X POST localhost:8080/user/john/run
```

### 04.QueryString

- QueryString parameters
- run example
```sh
go get github.com/gin-gonic/gin
cd 04.QueryString;
go run example.go

curl -X GET "http://localhost:8080/welcome?firstname=Jane&lastname=Doe"
```
