# Hello world

- go 빌드용 docker 실행

```sh
docker run --rm --privileged \
    -v $PWD:/go/ \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -it --name golang golang:latest /bin/bash
```
