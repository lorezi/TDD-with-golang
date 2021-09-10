package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(w io.Writer) {
	fmt.Fprint(w, "3")
}

func main() {
	Countdown(os.Stdout)
}
