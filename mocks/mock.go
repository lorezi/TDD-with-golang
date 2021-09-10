package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type RealSleeper struct{}

func (d *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintf(w, "%d\n", i)
	}
	s.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &RealSleeper{}
	Countdown(os.Stdout, sleeper)
}
