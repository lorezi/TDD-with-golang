package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type Sleeper interface {
	Sleep()
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
	sleeper := &ConfigurableSleeper{time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
