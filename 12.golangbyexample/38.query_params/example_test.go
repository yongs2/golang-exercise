package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// go test -v ./... -cover
func TestQueryParams(t *testing.T) {
	t.Run("getProducts1 #1", func(t *testing.T) {
		request := newGetRequest("/products1?filters=red&filters=color&filters=price&filters=brand")
		response := httptest.NewRecorder()

		getProducts1(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "red,color,price,brand")
	})

	// go test -run=TestQueryParams/getProducts1_#2 -v
	t.Run("getProducts1 #2", func(t *testing.T) {
		request := newGetRequest("/products1?filters=color")
		response := httptest.NewRecorder()

		getProducts1(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "color")
	})

	t.Run("getProducts2 #1", func(t *testing.T) {
		request := newGetRequest("/products2?filters=red&filters=color&filters=price&filters=brand")
		response := httptest.NewRecorder()

		getProducts2(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "red")
	})

	t.Run("getProducts3 #1", func(t *testing.T) {
		request := newGetRequest("/products3?filters=red&filters=color&filters=price&filters=brand")
		response := httptest.NewRecorder()

		getProducts3(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "red,color,price,brand")
	})

	t.Run("getProducts3 #2", func(t *testing.T) {
		request := newGetRequest("/products3?filters=red")
		response := httptest.NewRecorder()

		getProducts3(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "red")
	})

	t.Run("getProducts4 #1", func(t *testing.T) {
		request := newGetRequest("/products4?filters=red&filters=color&filters=price&filters=brand")
		response := httptest.NewRecorder()

		getProducts4(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "red")
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got status %d want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

func newGetRequest(query string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, query, nil)
	return req
}

// go test -bench=. -benchmem
func BenchmarkQueryParams(b *testing.B) {
	b.Run("getProducts1 #1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			request := newGetRequest("/products1?filters=red&filters=color&filters=price&filters=brand")
			response := httptest.NewRecorder()
			getProducts1(response, request)
		}
	})

	b.Run("getProducts1 #2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			request := newGetRequest("/products1?filters=color")
			response := httptest.NewRecorder()
			getProducts1(response, request)
		}
	})

	b.Run("getProducts2 #1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			request := newGetRequest("/products2?filters=red&filters=color&filters=price&filters=brand")
			response := httptest.NewRecorder()
			getProducts2(response, request)
		}
	})

	b.Run("getProducts3 #1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			request := newGetRequest("/products3?filters=red&filters=color&filters=price&filters=brand")
			response := httptest.NewRecorder()
			getProducts3(response, request)
		}
	})

	b.Run("getProducts3 #2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			request := newGetRequest("/products3?filters=red")
			response := httptest.NewRecorder()
			getProducts3(response, request)
		}
	})
	b.Run("getProducts4 #1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			request := newGetRequest("/products4?filters=red&filters=color&filters=price&filters=brand")
			response := httptest.NewRecorder()
			getProducts4(response, request)
		}
	})
}
