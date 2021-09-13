package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func server(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	return server
}

func TestRacer(t *testing.T) {

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {

		serverA := server(11)
		serverB := server(12)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})

}
