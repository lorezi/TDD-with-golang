package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(w io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(w, "%d\n", i)
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
