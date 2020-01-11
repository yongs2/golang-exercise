package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	oczipkin "contrib.go.opencensus.io/exporter/zipkin"
	"github.com/go-kit/kit/log"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/oklog/run"
	zipkin "github.com/openzipkin/zipkin-go"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"go.opencensus.io/trace"

	"08.go-kit/09.kitloc/hello/endpoints"
	svchttp "08.go-kit/09.kitloc/hello/http"
	"08.go-kit/09.kitloc/hello/service"
)

const (
	serviceName = "oc-gokit-example"
	zipkinURL   = "http://172.17.0.3:9411/api/v2/spans"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "svc", serviceName)
	}

	{
		var (
			reporter         = httpreporter.NewReporter(zipkinURL)
			localEndpoint, _ = zipkin.NewEndpoint(serviceName, ":0")
			exporter         = oczipkin.NewExporter(reporter, localEndpoint)
		)
		defer reporter.Close()

		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

		trace.RegisterExporter(exporter)
	}

	var handler http.Handler
	{
		svc := service.Service{}

		endpoints := endpoints.Endpoints{
			Hello: endpoints.MakeHelloEndpoint(svc),
		}

		endpoints.Hello = kitoc.TraceEndpoint("gokit:endpoint hello")(endpoints.Hello)

		var serverOptions []httptransport.ServerOption
		serverOptions = append(serverOptions, httptransport.ServerErrorLogger(logger))
		serverOptions = append(serverOptions, kitoc.HTTPServerTrace())

		handler = svchttp.NewHTTPHandler(endpoints, serverOptions...)
	}

	var g run.Group
	{
		var (
			listener, _ = net.Listen("tcp", ":0") // dynamic port assignment
			addr        = listener.Addr().String()
		)
		g.Add(func() error {
			logger.Log("msg", "service start", "transport", "http", "address", addr)
			return http.Serve(listener, handler)
		}, func(error) {
			listener.Close()
		})
	}
	{
		var (
			cancelInterrupt = make(chan struct{})
			c               = make(chan os.Signal, 2)
		)
		defer close(c)

		g.Add(func() error {
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	logger.Log("exit", g.Run())
}
