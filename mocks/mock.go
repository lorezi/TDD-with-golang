package main

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(w, "%d\n", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(w, finalWord)
}

func main() {
	// Countdown(os.Stdout, )
}
