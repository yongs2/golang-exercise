package addtransport

import (
	"context"
	"time"

	"golang.org/x/time/rate"

	"github.com/sony/gobreaker"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"

	"08.go-kit/04.addsvc/pkg/addendpoint"
	"08.go-kit/04.addsvc/pkg/addservice"

	addthrift "08.go-kit/04.addsvc/cmd/addsvc/thrift/gen-go/addsvc"
)

type thriftServer struct {
	ctx       context.Context
	endpoints addendpoint.Set
}

func NewThriftServer(endpoints addendpoint.Set) addthrift.AddService {
	return &thriftServer{
		endpoints: endpoints,
	}
}

func (s *thriftServer) Sum(ctx context.Context, a int64, b int64) (*addthrift.SumReply, error) {
	request := addendpoint.SumRequest{A: int(a), B: int(b)}
	response, err := s.endpoints.SumEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	resp := response.(addendpoint.SumResponse)
	return &addthrift.SumReply{Value: int64(resp.V), Err: err2str(resp.Err)}, nil
}

func (s *thriftServer) Concat(ctx context.Context, a string, b string) (*addthrift.ConcatReply, error) {
	request := addendpoint.ConcatRequest{A: a, B: b}
	response, err := s.endpoints.ConcatEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	resp := response.(addendpoint.ConcatResponse)
	return &addthrift.ConcatReply{Value: resp.V, Err: err2str(resp.Err)}, nil
}

func NewThriftClient(client *addthrift.AddServiceClient) addservice.Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))

	var sumEndpoint endpoint.Endpoint
	{
		sumEndpoint = MakeThriftSumEndpoint(client)
		sumEndpoint = limiter(sumEndpoint)
		sumEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Sum",
			Timeout: 30 * time.Second,
		}))(sumEndpoint)
	}

	var concatEndpoint endpoint.Endpoint
	{
		concatEndpoint = MakeThriftConcatEndpoint(client)
		concatEndpoint = limiter(concatEndpoint)
		concatEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Concat",
			Timeout: 30 * time.Second,
		}))(concatEndpoint)
	}

	return addendpoint.Set{
		SumEndpoint:    sumEndpoint,
		ConcatEndpoint: concatEndpoint,
	}
}

func MakeThriftSumEndpoint(client *addthrift.AddServiceClient) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addendpoint.SumRequest)
		reply, err := client.Sum(ctx, int64(req.A), int64(req.B))
		if err == addservice.ErrIntOverflow {
			return nil, err // special case; see comment on ErrIntOverflow
		}
		return addendpoint.SumResponse{V: int(reply.Value), Err: err}, nil
	}
}

func MakeThriftConcatEndpoint(client *addthrift.AddServiceClient) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addendpoint.ConcatRequest)
		reply, err := client.Concat(ctx, req.A, req.B)
		return addendpoint.ConcatResponse{V: reply.Value, Err: err}, nil
	}
}
