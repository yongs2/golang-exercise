#!/bin/sh

docker run --rm \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -p 6379:6379 \
    -d --name myredis redis
