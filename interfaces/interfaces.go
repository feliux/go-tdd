package main

import "math"

/*
// Minimal
func Perimeter(a, b float64) float64 {
	return 2*a + 2*b
}

func Area(a, b float64) float64 {
	return a * b
}
*/

/*
// Improve 1
type Rectangle struct {
	Height float64
	Width  float64
}

func Perimeter(r Rectangle) float64 {
	return 2*r.Height + 2*r.Width
}

func Area(r Rectangle) float64 {
	return r.Height * r.Width
}
*/

/*
// Improve 2
type Figure interface {
	Area() float64
	GetName() string
}

type Rectangle struct {
	Height float64
	Width  float64
	Name   string
}

type Circle struct {
	Radius float64
	Name   string
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) GetName() string {
	return r.Name
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) GetName() string {
	return c.Name
}
*/

// Improve 3
type Figure interface {
	Area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
