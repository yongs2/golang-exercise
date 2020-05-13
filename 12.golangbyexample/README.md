# Refer [golangbyexample.com](https://golangbyexample.com//)

## Design pattern

### 01.abstrct

- [abstract pattern](https://golangbyexample.com/abstract-factory-design-pattern-go/)

```sh
cd 01.abstract && go run .
```

### 02.builder

- [Builder pattern](https://golangbyexample.com/builder-pattern-golang/)

```sh
cd 02.builder && go run .
```

### 03.factory

- [Factory pattern](https://golangbyexample.com/golang-factory-design-pattern/)

```sh
cd 03.factory && go run .
```

### 04.object

- [Object Pool Pattern](https://golangbyexample.com/golang-object-pool/)

```sh
cd 04.object && go run .
```

### 05.prototype

- [Prototype Pattern](https://golangbyexample.com/prototype-pattern-go/)

```sh
cd 05.prototype && go run .
```

### 06.singleton

- [Silgleton pattern](https://golangbyexample.com/singleton-design-pattern-go/)

```sh
cd 06.singleton && go run .
```

### 07.chain

- [Chain of Responsibility Design Pattern](https://golangbyexample.com/chain-of-responsibility-design-pattern-in-golang/)

```sh
cd 07.chain && go run .
```

### 08.command

- [Command design pattern](https://golangbyexample.com/command-design-pattern-in-golang/)

```sh
cd 08.command && go run .
```

### 09.iterator

- [Iterator Design Pattern](https://golangbyexample.com/go-iterator-design-pattern/)

```sh
cd 09.iterator && go run .
```

### 10.mediator

- [Mediator design pattern](https://golangbyexample.com/mediator-design-pattern-golang/)

```sh
cd 10.mediator && go run .
```

### 11.memento

- [Memento design pattern](https://golangbyexample.com/memento-design-pattern-go/)

```sh
cd 11.memento && go run .
```

### 12.null

- [Null object design pattern](https://golangbyexample.com/null-object-design-pattern-golang/)

```sh
cd 12.null && go run .
```

### 13.observer

- [Observer design pattern](https://golangbyexample.com/observer-design-pattern-golang/)

```sh
cd 13.observer && go run .
```

### 14.state

- [State design pattern](https://golangbyexample.com/state-design-pattern-go/)

```sh
cd 14.state && go run .
```

### 15.strategy

- [Strategy design pattern](https://golangbyexample.com/strategy-design-pattern-golang/)

```sh
cd 15.strategy && go run .
```

### 16.template

- [Template Method Design Pattern](https://golangbyexample.com/template-method-design-pattern-golang/)

```sh
cd 16.template && go run .
```

### 17.visitor

- [Visitor design pattern](https://golangbyexample.com/visitor-design-pattern-go)

```sh
cd 17.visitor && go run .
```

### 18.adapter

- [Adapter design pattern](https://golangbyexample.com/adapter-design-pattern-go/)

```sh
cd 18.adapter && go run .
```

### 19.bridge

- [Bridge design pattern](https://golangbyexample.com/bridge-design-pattern-in-go/)

```sh
cd 19.bridge && go run .
```

### 20.composite

- [Composite design pattern](https://golangbyexample.com/composite-design-pattern-golang/)

```sh
cd 20.composite && go run .
```

### 21.facade

- [Facade design pattern](https://golangbyexample.com/facade-design-pattern-in-golang/)

```sh
cd 21.facade && go run .
```

### 22.flyweight

- [flyweight design pattern](https://golangbyexample.com/flyweight-design-pattern-golang/)

```sh
cd 22.flyweight && go run .
```

### 23.proxy

- [proxy design pattern](https://golangbyexample.com/proxy-design-pattern-in-golang/)

```sh
cd 23.proxy && go run .
```

### 24.inherit

- [OOP: Inheritance](https://golangbyexample.com/oop-inheritance-golang-complete/)

```sh
cd 24.inherit && go run .
```

### 25.polymorphism

- [OOP: Polymorphism](https://golangbyexample.com/oop-polymorphism-in-go-complete-guide/)

```sh
cd 25.polymorphism && go run exaxmple.go
```

- [Runtime Polymorphism](https://golangbyexample.com/runtime-polymorphism-go/)

```sh
cd 25.polymorphism && go run exaxmple_tax.go
```

- [Function/Method overloading](https://golangbyexample.com/function-method-overloading-golang/)

```sh
cd 25.polymorphism && go run exaxmple_overloading.go
```

### 26.abstract

- [Abstract class](https://golangbyexample.com/go-abstract-class/)

```sh
cd 26.abstract && go run example.go
```

### 27.Encapsulation

- [Encapsulation](https://golangbyexample.com/encapsulation-in-go/)

```sh
cd 27.encapsulation && go run example.go
```

### 28.ProtocolBuffer

- [Protocol Buffers](https://golangbyexample.com/protocol-buffers-go/)

```sh
cd 28.ProtocolBuffer
protoc -I ./ --go_out=./ ./person/person.proto
go run main.go
```

## Files
### 29.readfiles

- [Read a large file word by word](https://golangbyexample.com/read-large-file-word-by-word-go)
- [Read a large file line by line](https://golangbyexample.com/read-large-file-line-by-line-go/)

```sh
cd 29.readfiles
go run word.go
go run line.go
```

### 30.writefiles

- [write to a file](https://golangbyexample.com/write-to-a-file-go/)
- [append to a file](https://golangbyexample.com/append-file-golang)

```sh
cd 30.writefiles && go run write.go
cd 30.writefiles && go run writeFile.go
```

### 31.deletefile

- [Delete a file](https://golangbyexample.com/delete-file-go/)
- [Delete a folder](https://golangbyexample.com/delete-folder-go/)

```sh
cd 31.deletefile && touch sample.txt; go run del_file.go
cd 31.deletefile && mkdir sample && go run del_folder.go
```

### 32.updatefile

- [Change the modified/updated time and access time of a file](https://golangbyexample.com/change-updated-time-file-go/)

```sh
cd 32.updatefile && touch sample.txt; go run update_file.go
```

### 33.genericfile

- [Check if a file is a directory in Go](33.genericfile/checkfile.go)
- [Create an empty file in Go](33.genericfile/create_empty_file.go)
- [Iterate over a directory tree ](33.genericfile/iterate.go)
- [Touch a file in Go](33.genericfile/touch_file.go)
- [Move file from one location to another]()

### 34.time

- [time and date in Go](https://golangbyexample.com/all-about-time-and-date-golang/)

### 35.os

- [Execute shell file](https://golangbyexample.com/execute-shell-file-golang/)

### 36.array

- [Check if an item exists in a slice](https://golangbyexample.com/item-exists-slice-golang/)

### 37.map

- [Different ways of iterating over a map](https://golangbyexample.com/different-ways-iterating-over-map-go/)

### 38.query_params

- [net/http package get Query Params](https://golangbyexample.com/net-http-package-get-query-params-golang/)

```sh
curl -v "http://localhost:8080/products1?filters=red&filters=color&filters=price&filters=brand"
curl -v "http://localhost:8080/products1?filters=color"

curl -v "http://localhost:8080/products2?filters=red&filters=color&filters=price&filters=brand"

curl -v "http://localhost:8080/products3?filters=red&filters=color&filters=price&filters=brand"

curl -v "http://localhost:8080/products4?filters=red&filters=color&filters=price&filters=brand"
```

### 39.queue

- [queue](https://golangbyexample.com/queue-in-golang/)

### 40.stack

- [Stack](https://golangbyexample.com/stack-in-golang/)

### 41.set

- [Set implementation](https://golangbyexample.com/set-implementation-in-golang/)

### 42.linkedlist

- [Linked List](https://golangbyexample.com/singly-linked-list-in-golang/)

### 43.binarysearch

- [Binary Search tree](https://golangbyexample.com/binary-search-tree-in-go/)
- [Iterative Binary Search Tree](https://golangbyexample.com/iterative-binary-search-tree-go/)

### 44.heap

- [Heap](https://golangbyexample.com/heap-in-golang/)

### 45.trie

- [Trie Implementation](https://golangbyexample.com/trie-implementation-in-go/)

### 46.sort

- [Heap sort](https://golangbyexample.com/heapsort-in-golang/)

### 47.network

- [Validate an IP address](https://golangbyexample.com/validate-an-ip-address-in-go)
- [Check if an IP address is IPV4 or IPV6](https://golangbyexample.com/check-ip-address-is-ipv4-or-ipv6-go/)
- [Get the IP address from an incoming HTTP request](https://golangbyexample.com/golang-ip-address-http-request/)

```sh
curl -v -X GET http://172.17.0.2:8080/getIp
curl -v -H "X-REAL-IP: 172.17.0.2" -X GET http://172.17.0.2:8080/getIp
curl -v -H "X-FORWARDED-FOR: 172.17.0.2" -X GET http://172.17.0.2:8080/getIp
```

### 48.logger

- [Go Logger Rotate](https://golangbyexample.com/go-logger-rotation/)

### 49.channel

- [channel as function argument](https://golangbyexample.com/channel-function-argument-go/)

### 50.goroutine

- [Get number of currently running/active goroutines](https://golangbyexample.com/number-currently-running-active-goroutines/)
- [Wait for all Go routines to finish execution](https://golangbyexample.com/wait-all-goroutines-go/)
