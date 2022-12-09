@REM GOLANGCI_LINT_VER=v1.50.1
@REM curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${GOLANGCI_LINT_VER}
@REM go install -v github.com/gojp/goreportcard@latest
@REM curl -L https://git.io/vp6lP | sh (for gometalinter)
@REM apt-get -y update; apt-get -y install graphviz

docker run --rm --privileged ^
    -v %CD%:/go/src/ ^
    -v /var/run/docker.sock:/var/run/docker.sock ^
    -e GO111MODULE=on ^
    -it --name golang golang:1.18 /bin/bash
