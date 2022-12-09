docker run --rm --privileged ^
    -v %CD%:/go/src/ ^
    -v /var/run/docker.sock:/var/run/docker.sock ^
    -e GO111MODULE=on ^
    -it --name golang golang:1.18 /bin/bash
