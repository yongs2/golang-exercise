docker run --rm --privileged \
    -v $PWD:/go/src/ \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -it --name golang-grpc golang-grpc /bin/bash
