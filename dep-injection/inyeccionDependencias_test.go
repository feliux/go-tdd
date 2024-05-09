package main

import (
	"bytes"
	"testing"
)

// Solution 1
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Pepe")
	want := "Hello Pepe"
	got := buffer.String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
