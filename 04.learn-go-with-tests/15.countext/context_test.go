package context

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"time"
	"context"
)

type SpyStore struct {
	response string
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done() :
				s.t.Log("spy store got cancelled")
				return
			default :
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data :
		return res, nil
	}
}

func TestHandler(t *testing.T) {
	data := "hello, world"

	t.Run("return data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)
	
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
	
		svr.ServeHTTP(response, request)
	
		if response.Body.String() != data {
			t.Errorf(` got "%s", want "%s"`, response.Body.String(), data)
		}
	})
}
