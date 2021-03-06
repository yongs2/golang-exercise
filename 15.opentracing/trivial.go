package main

import (
	"fmt"
	"os"
	"time"

	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"

	"github.com/opentracing/opentracing-go"
)

func main() {
	// 1) Create a opentracing.Tracer that sends data to Zipkin
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:	"const",
			Param:	1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:		true,
			BufferFlushInterval:	1 * time.Second,
			CollectorEndpoint: "http://172.17.0.2:14268/api/traces",
		},
	}
	// cfg, err := config.FromEnv()
	// if err != nil {
	// 	// parsing errors might happen here, such as when we get a string where we expect a number
	// 	fmt.Printf("Could not parse Jaeger env vars: %s\n", err.Error())
	// 	return
	// }

	fmt.Printf("cfg.Sampler=[%v]\n", cfg.Sampler)
	if cfg.Sampler != nil {
		fmt.Printf("cfg.Sampler.Type=[%v]\n", cfg.Sampler.Type)
		fmt.Printf("cfg.Sampler.Param=[%v]\n", cfg.Sampler.Param)
	}

	fmt.Printf("cfg.Reporter=[%v]\n", cfg.Reporter)
	if cfg.Reporter != nil {
		fmt.Printf("cfg.Reporter.LocalAgentHostPort=[%v]\n", cfg.Reporter.LocalAgentHostPort)
		fmt.Printf("cfg.Reporter.CollectorEndpoint=[%v]\n", cfg.Reporter.CollectorEndpoint)
	}
	

	tracer, closer, err := cfg.New(
		"your_service_name",
		config.Logger(jaeger.StdLogger),
	)
	if (err != nil){
		fmt.Println("Cannot create tracer: %v\n", err.Error())
		os.Exit(1)
	}
	defer closer.Close()

	// 2) Demonstrate simple OpenTracing instrumentation
	parent := tracer.StartSpan("Parent")
	for i := 0; i < 20; i++ {
		parent.LogEvent(fmt.Sprintf("Starting child #%d", i))
		child := tracer.StartSpan("Child", opentracing.ChildOf(parent.Context()))
		time.Sleep(50 * time.Millisecond)
		child.Finish()
	}
	parent.LogEvent("A Log")
	parent.Finish()
}
