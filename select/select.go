package main

import (
	"net/http"
	"time"
)

func makeRequest(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)

	return duration
}

func Racer(a, b string) (winner string) {

	responseA := makeRequest(a)
	responseB := makeRequest(b)

	if responseA < responseB {
		return a
	}

	return b
}
