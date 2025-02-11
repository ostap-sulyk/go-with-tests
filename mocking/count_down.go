package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	countDownStart = 3
	finalWord      = "Go!"
	write          = "write"
	sleep          = "sleep"
)

// sleeper interface
type Sleeper interface {
	Sleep()
}

// SpySleeper
type SpySleeper struct{ Calls int }

func (s *SpySleeper) Sleep() { s.Calls++ }

// DefaultSleeper
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() { time.Sleep(1 * time.Second) }

type SpyCountdownOperations struct{ Calls []string }

func (s *SpyCountdownOperations) Sleep() { s.Calls = append(s.Calls, sleep) }

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(out, finalWord)
}
