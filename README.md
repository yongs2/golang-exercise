# golang-exercise


## 00. 참고 자료

- [A Tour of go](https://tour.golang.org/list)
  - [Welcome](https://tour.golang.org/welcome/1) 은 1 ~ 5
  - Basics
    - [Packages, variables, and functions.](https://tour.golang.org/basics/1) 은  1~17
    - [Flow control statements: for, if, else, switch and defer](https://tour.golang.org/flowcontrol/1) 은 1 ~ 14
    - [More types: structs, slices, and maps.](https://tour.golang.org/moretypes/1) 은 1~27
  - [Methods and interfaces](https://tour.golang.org/methods/1) 은 1~26
  - [Concurrency](https://tour.golang.org/concurrency/1) 은 1~11
- [한글 버전](https://go-tour-kr.appspot.com)
  - [#30](https://go-tour-kr.appspot.com/#30)

## 01.Hello : Hello world

- go 빌드용 docker 실행

```sh
docker run --rm --privileged \
    -v $PWD:/go/ \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -it --name golang golang:latest /bin/bash
```

- go 소스를 하나씩 빌드하면 파일명으로 실행 파일 생성

```sh
go build hello.go
./hello
```
