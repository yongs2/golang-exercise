#!/bin/sh

IMG_NAME=mariadb:latest
CONTAINER_NAME=mariadb

container_start() {
docker run --rm \
    -v `pwd`/initdb.d:/docker-entrypoint-initdb.d \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=root \
    -d --name ${CONTAINER_NAME} ${IMG_NAME}
}

container_exec() {
    docker exec -it ${CONTAINER_NAME} /bin/bash
}
container_stop() {
    docker stop ${CONTAINER_NAME}
}

case $1 in
    start)
        container_start ;;
    exec)
        container_exec ;;
    stop)
        container_stop ;;
    restart)
        container_stop 
        container_start ;;
    *)
        echo "$0 start|exec|restart|stop" ;;
esac

# end of script.
