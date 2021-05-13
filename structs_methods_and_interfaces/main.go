package main

import "math"

type Shape interface {
	Area() float64
	// Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) Area() (a float64) {
	a = 0.5 * (t.Base * t.Height)
	return
}

func (c *Circle) Area() (a float64) {
	a = math.Pow(c.Radius, 2) * math.Pi
	return
}

func (r *Rectangle) Area() (a float64) {
	a = r.Width * r.Height
	return
}

func Perimeter(rect Rectangle) (p float64) {
	p = 2 * (rect.Height + rect.Width)
	return
}

func main() {

}
