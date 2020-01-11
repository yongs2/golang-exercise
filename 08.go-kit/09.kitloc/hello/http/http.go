package http

import (
	"context"
	"encoding/json"
	"net/http"

	"08.go-kit/09.kitloc/hello/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(endpoints endpoints.Endpoints, options ...httptransport.ServerOption) http.Handler {
	m := http.NewServeMux()
	m.Handle("/hello", httptransport.NewServer(endpoints.Hello, DecodeHelloRequest, EncodeHelloResponse, options...))
	return m
}
func DecodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.HelloRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeHelloResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
