package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"fmt"
)

func TestRacer(t *testing.T) {
	t.Run("return fastServer", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
	
		defer slowServer.Close()
		defer fastServer.Close()
	
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		fmt.Println("slowURL=", slowURL, ",fastURL=", fastURL)
	
		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if got != want {
			t.Errorf("got %q, want %q, err %q", got, want, err)
		}
	})

	t.Run("returns an error if a server doesn't response within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Millisecond)
		serverB := makeDelayedServer(12 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
