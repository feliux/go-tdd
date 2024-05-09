package main

import (
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper() // if error show the case lane, not this lane
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	t.Run("Test para Juan", func(t *testing.T) {
		want := "Hello, Juan"
		got := Hello("Juan")
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test para Pepe", func(t *testing.T) {
		want := "Hello, Pepe"
		got := Hello("Pepe")
		assertCorrectMessage(t, got, want)
	})
}

func TestSum(t *testing.T) {
	t.Run("Test A", func(t *testing.T) {
		given := []int{1, 2, 3, 4, 5}
		want := 15
		got := Sum(given)
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, given)
		}
	})
	t.Run("Test B", func(t *testing.T) {
		given := []int{1, 1, 1}
		want := 3
		got := Sum(given)
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, given)
		}
	})
	t.Run("Test C", func(t *testing.T) {
		want := []int{3, 9}
		got := SumAll([]int{1, 2}, []int{0, 9})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
