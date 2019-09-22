# Hello world

- tour.golang.org

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
