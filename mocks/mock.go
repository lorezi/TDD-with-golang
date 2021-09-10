package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(w io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(w, "%d\n", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
