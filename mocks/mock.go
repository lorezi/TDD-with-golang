package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(w io.Writer) {
	fmt.Fprint(w, `3
	2
	1
	Go!`)
}

func main() {
	Countdown(os.Stdout)
}
