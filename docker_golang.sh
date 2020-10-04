docker run --rm --privileged \
    -v $PWD:/go/src/ \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -e GO111MODULE=auto \
    -it --name golang-grpc golang-grpc /bin/bash
