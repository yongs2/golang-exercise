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

### 05.Multipart

- Multipart/Urlencoded Form
- run example
```sh
go get github.com/gin-gonic/gin
cd 05.Multipart
go run example.go

curl --form message="hello world" --form nick="john" http://localhost:8080/form_post
```

### 06.PostForm

- query + post form
- run example
```sh
go get github.com/gin-gonic/gin
cd 06.PostForm
go run example.go

curl --form message="hello world" --form name="john" "http://localhost:8080/post?id=1001&page=1"
```

### 07.Map

- Map as querystring or postform parameters
- run example
```sh
go get github.com/gin-gonic/gin
cd 07.Map
go run example.go

curl -d "names[first]=thinkerou&names[second]=tianou" -X POST "http://localhost:8080/post?ids[a]=1234&ids[b]=hello"
```

### 08.UploadFiles

- Single file
```sh
go get github.com/gin-gonic/gin
cd 08.UploadFiles
go run singlefile.go

curl -X POST http://localhost:8080/upload \
  -F "file=@/go/src/README.md" \
  -H "Content-Type: multipart/form-data"
```
