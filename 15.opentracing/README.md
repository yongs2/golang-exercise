# opentracing

## run jaeger UI

```sh
docker run --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:latest
```

## fix udp send size

```sh
sysctl -a | grep udp

echo "net.ipv4.udp_rmem_min = 4096" >> /etc/sysctl.conf
echo "net.ipv4.udp_wmem_min = 4096" >> /etc/sysctl.conf
sysctl -p

sysctl -a | grep udp
```

## example

- https://github.com/opentracing-contrib/examples/go
- https://github.com/opentracing/opentracing-go

## test

```sh
export JAEGER_SERVICE_NAME=test
export JAEGER_REPORTER_LOG_SPANS=true
export JAEGER_DISABLED=true
```

- udp로 연동하는 경우는 다음과 같이 설정하나, udp 전송시 1500 bytes 이상인 경우 packet drop 이 발생될 수 있음, 그래서 sidecar 패턴에서만 사용

```sh
docker inspect jaeger
export JAEGER_AGENT_HOST=172.17.0.2
export JAEGER_AGENT_PORT=6831
```

- http 연동인 경우가 일반적으로 사용함

```sh
export JAEGER_SAMPLER_TYPE=const
export JAEGER_SAMPLER_PARAM=1
export JAEGER_REPORTER_FLUSH_INTERVAL=1s
export JAEGER_ENDPOINT=http://172.17.0.2:14268/api/traces
```
