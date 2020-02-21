#!/bin/sh

MAX=3

redis_cluster_run()
{
    for ind in `seq 1 ${MAX}`; do \
        let port=6379+${ind}-1
        name=redis-master-${ind}
        echo $name, $port
        docker run --rm \
            -v /var/run/docker.sock:/var/run/docker.sock \
            -p ${port}:${port} \
            -d --name ${name} redis:latest \
            redis-server --port ${port} --cluster-enabled yes --cluster-config-file node.conf --cluster-node-timeout 5000 --bind 0.0.0.0
    done
}

redis_cluster_stop()
{
    for ind in `seq 1 ${MAX}`; do \
        name=redis-master-${ind}
        docker stop ${name}
        docker rm ${name}
    done
}

redis_cluster_create()
{
    for ind in `seq 1 ${MAX}`; do
        name=redis-master-${ind}
        IP=`docker inspect -f \
            '{{(index .NetworkSettings.Networks "bridge").IPAddress}}' \
            ${name}`
        PORT=`docker port ${name}|cut -d'/' -f1`
        IPS="${IPS} ${IP}:${PORT}"
    done
    echo "IPS=$IPS"
    if [ "X$IPS" != "X" ] ; then
        docker run -i --rm redis:latest redis-cli --cluster create ${IPS}
    else
        echo "Not Cluster"
    fi
}

case $1 in
    start)
        redis_cluster_run ;;
    create)
        redis_cluster_create ;;
    stop)
        redis_cluster_stop ;;
    *)
        echo "$0 start|create|stop" ;;
esac

# end of script.