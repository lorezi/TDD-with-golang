package main

import "net/http"

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

/*
Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory perspective.
*/
func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
