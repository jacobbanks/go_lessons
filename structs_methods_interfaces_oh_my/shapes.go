package stucts_methods_interfaces

import "math"



type Rectangle struct {
	Width float64
	Length float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	base float64
	height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}
