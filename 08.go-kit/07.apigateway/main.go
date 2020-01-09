package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"

	"08.go-kit/04.addsvc/pkg/addendpoint"
	"08.go-kit/04.addsvc/pkg/addservice"
	"08.go-kit/04.addsvc/pkg/addtransport"
)

func main() {
	var (
		httpAddr     = flag.String("http.addr", ":8500", "Address for HTTP (JSON) server")
		consulAddr   = flag.String("consul.addr", "", "Consul agent address")
		retryMax     = flag.Int("retry.max", 3, "per-request retries to different instances")
		retryTimeout = flag.Duration("retry.timeout", 500*time.Millisecond, "per-request timeout, including retries")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		if len(*consulAddr) > 0 {
			consulConfig.Address = *consulAddr
		}
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = consulsd.NewClient(consulClient)
	}

	tracer := stdopentracing.GlobalTracer() // no-op
	zipkinTracer, _ := stdzipkin.NewTracer(nil, stdzipkin.WithNoopTracer(true))
	ctx := context.Background()
	r := mux.NewRouter()

	// addsvc
	{
		var (
			tags        = []string{}
			passingOnly = true
			endpoints   = addendpoint.Set{}
			instancer   = consulsd.NewInstancer(client, logger, "addsvc", tags, passingOnly)
		)
		{
			factory := addsvcFactory(addendpoint.MakeSumEndpoint, tracer, zipkinTracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(*retryMax, *retryTimeout, balancer)
			endpoints.SumEndpoint = retry
		}
		{
			factory := addsvcFactory(addendpoint.MakeConcatEndpoint, tracer, zipkinTracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(*retryMax, *retryTimeout, balancer)
			endpoints.ConcatEndpoint = retry
		}

		r.PathPrefix("/addsvc").Handler(http.StripPrefix("/addsvc", addtransport.NewHTTPHandler(endpoints, tracer, zipkinTracer, logger)))
	}

	// stringsvc
	{
		var (
			tags        = []string{}
			passingOnly = true
			uppercase   endpoint.Endpoint
			count       endpoint.Endpoint
			instancer   = consulsd.NewInstancer(client, logger, "stringsvc", tags, passingOnly)
		)
		{
			factory := stringsvcFactory(ctx, "GET", "/uppercase")
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(*retryMax, *retryTimeout, balancer)
			uppercase = retry
		}
		{
			factory := stringsvcFactory(ctx, "GET", "/count")
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(*retryMax, *retryTimeout, balancer)
			count = retry
		}

		r.Handle("/stringsvc/uppercase", httptransport.NewServer(uppercase, decodeUppercaseRequest, encodeJSONResponse))
		r.Handle("/stringsvc/count", httptransport.NewServer(count, decodeCountRequest, encodeJSONResponse))
	}

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errc <- http.ListenAndServe(*httpAddr, r)
	}()

	logger.Log("exit", <-errc)
}

func addsvcFactory(makeEndpoint func(addservice.Service) endpoint.Endpoint, tracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := addtransport.NewGRPCClient(conn, tracer, zipkinTracer, logger)
		endpoint := makeEndpoint(service)

		return endpoint, conn, nil
	}
}

func stringsvcFactory(ctx context.Context, method, path string) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}
		tgt, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		tgt.Path = path

		var (
			enc httptransport.EncodeRequestFunc
			dec httptransport.DecodeResponseFunc
		)
		switch path {
		case "/uppercase":
			enc, dec = encodeJSONRequest, decodeUppercaseResponse
		case "/count":
			enc, dec = encodeJSONRequest, decodeCountResponse
		default:
			return nil, nil, fmt.Errorf("unknown stringsvc path %q", path)
		}

		return httptransport.NewClient(method, tgt, enc, dec).Endpoint(), nil, nil
	}
}

func encodeJSONRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeUppercaseResponse(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response struct {
		V   string `json:"v"`
		Err string `json:"err,omitempty"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeCountResponse(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response struct {
		V int `json:"v"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeUppercaseRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var request struct {
		S string `json:"s"`
	}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var request struct {
		S string `json:"s"`
	}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
