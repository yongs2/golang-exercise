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
  - [#70](https://go-tour-kr.appspot.com/#70)
- [Go Bootcamp](http://www.golangbootcamp.com/book/)
- [Go by Example](https://gobyexample.com/) : ([한글번역](https://mingrammer.com/gobyexample), [github](https://github.com/mingrammer/gobyexample))
- [Awsome go](https://awesome-go.com/)
  - [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) 
  - [gophercises.com](https://gophercises.com/)
- [gRPC](https://github.com/grpc/grpc-go)
- [go-cshared-example](https://github.com/vladimirvivien/go-cshared-examples)
  - export LD_LIBRARY_PATH=/go/src/07.go-cshared/01.goutil/libgoutil:${LD_LIBRARY_PATH}

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

## 99. [VScode 설정](https://ux.stories.pe.kr/111)

- Remote Containers 에 접속할 때의 설정 정보는 아래 폴더에 기록됨

```sh
/Users/<user>/Library/Application Support/Code/User/globalStorage/ms-vscode-remote.remote-containers/imageConfigs/
```

## Refer
- [A command-line benchmarking tool](https://github.com/sharkdp/hyperfine)

- [NB-IoT NIDD via SCEF](http://www.definitionnetworks.com/3gpp-nidd-via-scef-nb-iot/)
- [Figure 4.6. IoT Architectures with EPC and 5GC.](https://www.5gamericas.org/wp-content/uploads/2019/07/5G_Americas_White_Paper_on_5G_IOT_FINAL_7.16.pdf) : T8 (SCS/AS - SCEF) vs N33 (AF - NEF)

- [3GPP Documentation](https://github.com/emanuelfreitas/3gpp-documentation)

- [free5GC Stage 2](https://bitbucket.org/free5GC/free5gc-stage-2/)
- [Docker-free5gc](https://github.com/abousselmi/docker-free5gc)
- [free5gc-nssf](https://github.com/stevenchiu30801/free5gc-nssf)

- [OpenAPI.Tools](https://openapi.tools/)
- [OpenAPI Generator](https://openapi-generator.tech/)
  - [dockerfile](https://hub.docker.com/r/openapitools/openapi-generator/dockerfile)

```sh
git clone https://github.com/jdegre/5GC_APIs;

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/TS29122_NIDD.yaml -g go -o /local/out/go/t8-client

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/TS29122_NIDD.yaml -g go-server -o /local/out/go/t8-server

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/TS29122_NIDD.yaml -g go-gin-server -o /local/out/go/t8-gin-server

docker run --rm --privileged \
    -v $PWD/out/go/t8-server:/go/src/ \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -it --name t8-server golang-grpc /bin/bash

mkdir -p /go/src/github.com/GIT_USER_ID/GIT_REPO_ID/
ln -s /go/src/go /go/src/github.com/GIT_USER_ID/GIT_REPO_ID/go

CID=`docker ps -f name=t8-server -q`
GEN_IP=`docker inspect --format '{{.NetworkSettings.IPAddress}}' ${CID}`
export GEN_IP=172.17.0.3
curl -i -s -v -X GET --header 'Content-Type: application/json' --header 'Accept: application/json' \
-d '{"scsAsId": "scsAsId001"}' \
http://${GEN_IP}:8080/3gpp-nidd/v1/scsAsId001/configurations


docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/TS29522_NIDDConfigurationTrigger.yaml -g go -o /local/out/go/niddconf-client

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/TS29522_NIDDConfigurationTrigger.yaml -g go-server -o /local/out/go/niddconf-server

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/TS29522_NIDDConfigurationTrigger.yaml -g go-gin-server -o /local/out/go/niddconf-gin-server

docker run --rm -v $PWD:/local \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -it --name openapi-generator-cli openapitools/openapi-generator-cli /bin/bash

cd /opt/openapi-generator/modules/openapi-generator-cli/target/; java -jar openapi-generator-cli.jar version

wget https://github.com/stevenchiu30801/free5gc-nssf/raw/master/openapi/openapi-generator-cli.jar 
mv openapi-generator-cli.jar openapi-generator-cli-4.0.0-SNAPSHOT.jar
java -jar openapi-generator-cli-4.0.0-SNAPSHOT.jar version

git clone https://github.com/OpenAPITools/openapi-generator
docker build -t openapi-generator-cli -f Dockerfile .

git clone https://github.com/openapitools/openapi-generator
cd openapi-generator
./run-in-docker.sh mvn package
```

- Finite State Machine for go
  - [Programming business processes in Golang](https://medium.com/swlh/programming-business-processes-in-golang-f3612108d16b)
  - [Developing a statemachine in golang](https://www.codingdream.com/index.php/developing-a-statemachine-in-golang)
  - [State Design pattern in Go](https://golangbyexample.com/state-design-pattern-go/)
  - [Finite State Machine for Go](https://github.com/looplab/fsm) : 17 Jan 2019, Star:926, Fork:139
  - [Finite state machine for Go](https://github.com/bykof/stateful) : 31 Jul 2019, Star:114, Fork:3
  - [Finite State Machine for Go inspired by Akka FSM](https://github.com/dyrkin/fsm) : 17 days ago, Star:45, Fork:4

### [Go Report Card](https://github.com/gojp/goreportcard)

- install

```sh
go get github.com/gojp/goreportcard
cd $GOPATH/src/github.com/gojp/goreportcard
make install
go get github.com/gojp/goreportcard/cmd/goreportcard-cli
goreportcard-cli -v
```
