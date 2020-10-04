#!MAKEFILE

IMG_URI=jaegertracing/opentelemetry-all-in-one:latest
IMG_NAME=jaeger

.PHONY: all
all: stop run

.PHONY: run
run:
	docker run -d --rm --name ${IMG_NAME} \
		-p 5775:5775/udp \
		-p 6831:6831/udp \
		-p 6832:6832/udp \
		-p 5778:5778 \
		-p 55680:55680 \
		-p 16686:16686 \
		-p 14268:14268 \
		${IMG_URI}

.PHONY: stop
stop:
	docker stop ${IMG_NAME} || :
	docker rm ${IMG_NAME} || :
