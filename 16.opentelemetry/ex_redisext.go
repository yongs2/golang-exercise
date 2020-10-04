package main

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redisext"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/label"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// export JAEGER_URI=http://172.17.0.4:14268/api/traces
func InitTraceProvider(service string) func() {
	jaegerCollectorUri := os.Getenv("JAEGER_URI")
	log.Printf("JAEGER_URI=[%v]\n", jaegerCollectorUri)
	// Create and install Jaeger export pipeline
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(jaegerCollectorUri),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: service,
			Tags: []label.KeyValue{
				label.String("exporter", "jaeger"),
				label.Float64("float", 312.23),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		log.Fatal(err)
	}

	return func() {
		flush()
	}
}

// docker run -d --rm --name redis -p 6379:6379 redis:latest
// export REDIS_HOST=172.17.0.2; export REDIS_PORT=6379
func main() {
	fn := InitTraceProvider("tracing-demo")
	defer fn()

	redisAddr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	log.Println("REDIS=", redisAddr)

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	defer client.Close()

	ctx := context.Background()
	client.AddHook(redisext.OpenTelemetryHook{})

	tr := global.Tracer("component-main")
	ctx, span := tr.Start(ctx, "foo")
	span.End()

	key := "key1"
	value := "value1"

	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Println("Setting, Err=", err)
	}
	log.Printf("client.Set(%v,%v)=Err[%v]\n", key, value, err)

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Println("Getting, Err=", err)
	}
	log.Printf("client.Get(%v)=[%v].Err[%v]\n", key, val, err)
}
