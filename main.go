package main

import (
	"go-with-tests/mocking"
	"os"
)

func main() {
	ds := &mocking.DefaultSleeper{}

	mocking.Countdown(os.Stdout, ds)
}
