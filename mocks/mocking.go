package main

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWrod      = "Go!"
	countdownStart = 3
)

/*
// The problem: we have to wait 3*3 seconds to complete tests
const sleep = 3
func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
		time.Sleep(sleep * time.Second)
	}
	_, _ = fmt.Fprintln(out, finalWrod)
}
*/

/*
// Solution 1
const sleep = 3
type Sleeper interface {
	Sleep()
}

type DefaultSleepr struct {}

func (s *DefaultSleepr) Sleep() {
	time.Sleep(sleep * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	_, _ = fmt.Fprintln(out, finalWrod)
}
*/

// Solution 2: Testify
type Sleeper interface {
	Sleep(sec int) error
}

type DefaultSleepr struct{}

func (s *DefaultSleepr) Sleep(sec int) error {
	time.Sleep(time.Duration(sec) * time.Second)
	return nil
}

func Countdown(out io.Writer, sleeper Sleeper, sec int) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
		_ = sleeper.Sleep(sec)
	}
	_, _ = fmt.Fprintln(out, finalWrod)
}
