package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestExample(t *testing.T) {
	t.Run("Success to ping", func(t *testing.T) {
		apitest.New().
			Report(apitest.SequenceDiagram()).
			Handler(newApp().Router).
			Get("/ping").
			Expect(t).
			Body(`{"message": "pong"}`).
			Status(http.StatusOK).
			End()
	})

	t.Run("Fail to ping", func(t *testing.T) {
		apitest.New().
			Report(apitest.SequenceDiagram()).
			Handler(newApp().Router).
			Get("/pick").
			Expect(t).
			Status(http.StatusNotFound).
			End()
	})
}
