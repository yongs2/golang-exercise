# Hello world

- go 빌드용 docker 실행

```sh
docker run --rm --privileged \
    -v $PWD:/go/ \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -it --name golang golang:latest /bin/bash
```

- go 소스를 빌드하면 실행 파일명이 go로 생성

```sh
go build
./go
```
