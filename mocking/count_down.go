package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface{ Sleep() }
type SpySleeper struct{ Calls int }
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() { time.Sleep(1 * time.Second) }

func (s *SpySleeper) Sleep() { s.Calls++ }

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
