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

	slowServer := server(20)
	fastServer := server(0)

	defer slowServer.Close()
	defer fastServer.Close()

	want := fastServer.URL
	got := Racer(slowServer.URL, fastServer.URL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
