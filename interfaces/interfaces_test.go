package main

import (
	"testing"
)

/*
// Minimal
func TestPerimeter(t *testing.T) {
	want := 40.0
	got := Perimeter(10.0, 10.0)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestArea(t *testing.T) {
	want := 100.0
	got := Area(10.0, 10.0)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
*/

/*
// Improve 1
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Height: 10.0,
		Width:  10.0,
	}
	want := 40.0
	got := Perimeter(rectangle)
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{
		Height: 10.0,
		Width:  10.0,
	}
	want := 100.0
	got := Area(rectangle)
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
*/

/*
// Improve 2
func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, want float64, figure Figure) {
		t.Helper() // if error show the case lane, not this lane
		got := figure.Area()
		if got != want {
			t.Errorf("got %.2f, want %.2f, given %v", got, want, figure.GetName())
		}
	}
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			Height: 10.0,
			Width:  10.0,
			Name:   "rectangle",
		}
		want := 100.0
		checkArea(t, want, &rectangle)
	})
	t.Run("Circle", func(t *testing.T) {
		circle := Circle{
			Radius: 10.0,
			Name:   "circle",
		}
		want := 314.1592653589793
		checkArea(t, want, &circle)
	})
}
*/

// Improve 3. Table driven test
func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, name string, want float64, figure Figure) {
		t.Helper() // if error show the case lane, not this lane
		got := figure.Area()
		if got != want {
			t.Errorf("got %.2f, want %.2f, given %v", got, want, name)
		}
	}
	// Tale Driven Tests
	areaTest := []struct {
		name   string
		figure Figure  // input
		want   float64 // output
	}{
		{
			"rectangle",
			&Rectangle{
				Height: 10.0,
				Width:  10.0,
			},
			100.0,
		},
		{
			"circle",
			&Circle{
				Radius: 10.00,
			},
			314.1592653589793,
		},
	}

	for _, tdt := range areaTest {
		checkArea(t, tdt.name, tdt.want, tdt.figure)
	}
}
