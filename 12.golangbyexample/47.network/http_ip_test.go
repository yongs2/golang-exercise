package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

// go test -v http_ip_test.go http_ip.go
func TestHttpIp(t *testing.T) {
	handlerFunc := http.HandlerFunc(handleRequest)
	t.Run("Fail to getIp", func(t *testing.T) {
		apitest.New().
			//Report(apitest.SequenceDiagram()).
			Handler(handlerFunc).
			Get("/getIp").
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})

	t.Run("Success to getIp using X-REAL-IP", func(t *testing.T) {
		apitest.New().
			//Report(apitest.SequenceDiagram()).
			Handler(handlerFunc).
			Get("/getIp").
			Header("X-REAL-IP", "172.17.0.2").
			Expect(t).
			Body(`172.17.0.2`).
			Status(http.StatusOK).
			End()
	})

	t.Run("Success to getIp using X-FORWARDED-FOR", func(t *testing.T) {
		apitest.New().
			//Report(apitest.SequenceDiagram()).
			Handler(handlerFunc).
			Get("/getIp").
			Header("X-FORWARDED-FOR", "172.17.0.2").
			Expect(t).
			Body(`172.17.0.2`).
			Status(http.StatusOK).
			End()
	})
}
