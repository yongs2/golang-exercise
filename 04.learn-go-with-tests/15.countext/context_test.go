package context

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"time"
	"context"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	fmt.Println("Fetch:", s.response)
	return s.response
}

func (s *StubStore) Cancel() {
}

type SpyStore struct {
	response string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	data := "hello, world"

	t.Run("return hello, world", func(t *testing.T) {	
		svr := Server(&StubStore{data})
	
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
	
		svr.ServeHTTP(response, request)
	
		if response.Body.String() != data {
			t.Errorf(` got "%s", want "%s"`, response.Body.String(), data)
		}
	})
	
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})
}
