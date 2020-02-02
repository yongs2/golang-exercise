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

#### 1) Single file
```sh
go get github.com/gin-gonic/gin
cd 08.UploadFiles
go run singlefile.go

curl -X POST http://localhost:8080/upload \
  -F "file=@/go/src/README.md" \
  -H "Content-Type: multipart/form-data"
```

#### 2) Multiple files
```sh
go get github.com/gin-gonic/gin
cd 08.UploadFiles
go run multiplefiles.go

curl -X POST http://localhost:8080/upload \
  -F "upload[]=@/go/src/README.md" \
  -F "upload[]=@/go/src/LICENSE" \
  -H "Content-Type: multipart/form-data"
```

### 09.GroupingRoutes

- Grouping Routes
- run example
```sh
go get github.com/gin-gonic/gin
cd 09.GroupingRoutes
go run example.go

curl -d "body" -X POST "http://localhost:8080/v1/login"
curl -d "body" -X POST "http://localhost:8080/v2/login"
```

### 10.UsingMiddleware

- Using middleware
- run example
```sh
go get github.com/gin-gonic/gin
cd 10.UsingMiddleware
go run example.go

curl -X GET "http://localhost:8080/benchmark"
curl -d "body" -X POST "http://localhost:8080/login"
curl -d "body" -X POST "http://localhost:8080/submit"
curl -d "body" -X POST "http://localhost:8080/read"
curl -X GET "http://localhost:8080/testing/analytics"
```

### 11.log

- write log
- run example
```sh
go get github.com/gin-gonic/gin
cd 11.log
go run example.go

curl -X GET "http://localhost:8080/ping"
```

### 12.customlog

- custom log format
- run example
```sh
go get github.com/gin-gonic/gin
cd 12.customlog
go run example.go

curl -X GET "http://localhost:8080/ping"
```

### 13.ModelBinding

- Model binding and validation
- run example
```sh
go get github.com/gin-gonic/gin
cd 13.ModelBinding
go run example.go

curl -v -d '{"user" : "manu", "password" : "123"}' -H 'content-type: application/json' -X POST "http://localhost:8080/loginJSON"
curl -v -d '{"user" : "manu"}' -H 'content-type: application/json' -X POST "http://localhost:8080/loginJSON"
curl -v -d '{"user" : "manu", "password" : "012"}' -H 'content-type: application/json' -X POST "http://localhost:8080/loginJSON"

curl -v -d '<?xml version="1.0" encoding="UTF-8"?><root><user>manu</user><password>123</password></root>' -H 'content-type: application/xml' -X POST "http://localhost:8080/loginXML"

curl -v -d 'user=manu&password=123' -X POST "http://localhost:8080/loginForm"
curl -v -X POST "http://localhost:8080/loginForm?user=manu&password=123"
```

### 14.CustomValidator

- register custom validators
- run example
```sh
go get github.com/gin-gonic/gin
go get github.com/go-playground/validator
cd 14.CustomValidator
go run example.go

curl -v -X GET "http://localhost:8080/bookable?check_in=2018-04-16&check_out=2018-04-17"
curl -v -X GET "http://localhost:8080/bookable?check_in=2018-04-16&check_out=2018-03-17"
```

### 15.BindQuery
- run example for Only Bind Query String
```sh
go get github.com/gin-gonic/gin
cd 15.BindQuery
go run OnlyBindQuery.go

curl -X GET "http://localhost:8080/testing?name=eason&address=xyz"
curl -X POST "http://localhost:8080/testing?name=eason&address=xyz" --data 'name=ignore&address=ignore' -H "Content-Type:application/x-www-form-urlencoded"
```

- run example to Bind Query String or Post Data
```sh
go run BindQueryString.go

curl -X GET "http://localhost:8080/testing?name=appleboy&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"
curl -X POST "http://localhost:8080/testing?name=appleboy&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"
curl -v -d '{"name" : "appleboy", "address" : "xyz", "birthday" : "1992-03-15", "createTime" : 1562400033000000123, "unixTime" : 1562400033 }' -H 'content-type: application/json' -X POST "http://localhost:8080/testing"
```

- run example to Bind Uri
```sh
go run BindUri.go

curl -v -X GET "http://localhost:8080/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3"
curl -v -X GET "http://localhost:8080/thinkerou/not-uuid"
```

- run example to Bind Header
```sh
go run BindHeader.go

curl -v -H "rate:300" -H "domain:music" -X GET "http://localhost:8080"
```

- run example to bind Multipart/Urlencode
```sh
go run BindMultipart

curl -X POST -v --form name=user --form "avatar=@./README.md" http://localhost:8080/profile
```
