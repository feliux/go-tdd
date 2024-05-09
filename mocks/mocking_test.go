package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/mock"
)

/*
// The problem
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	Countdown(buffer, )
	want := `3
2
1
Go!
`
	got := buffer.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
*/

/*
// Solution 1
type SleeperMock struct {}
func (s *SleeperMock) Sleep() {}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	Countdown(buffer, &SleeperMock{})
	want := `3
2
1
Go!
`
	got := buffer.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
*/

// Solution 2: Testify
type SleeperMock struct {
	mock.Mock
}

func (m *SleeperMock) Sleep(sec int) error {
	args := m.Called(sec)
	return args.Error(0)
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeperMock := &SleeperMock{}
	//sleeperMock.On("Sleep", 1).Return(nil) // fails
	sleeperMock.On("Sleep", 3).Return(nil)
	Countdown(buffer, sleeperMock, 3)
	want := `3
2
1
Go!
`
	got := buffer.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	sleeperMock.AssertExpectations(t)
	// number of method calls
	//sleeperMock.AssertNumberOfCalls(t, "Sleep", 1) // fails
	sleeperMock.AssertNumberOfCalls(t, "Sleep", 3)
}
