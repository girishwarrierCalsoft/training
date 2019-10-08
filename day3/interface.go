package main

import (
	"fmt"
)

// shape interface
type shape interface {
	area() float64
}

// square struct
type square struct {
	side float64
}

// rectangle struct
type rectangle struct {
	l float64
	b float64
}

func display(s shape) {
	fmt.Printf("Area is %0.1f\n", s.area())
}

func (r rectangle) area() float64 {
	return float64(r.l * r.b)
}
func (s square) area() float64 {
	return float64(s.side * s.side)
}

func main() {
	r := rectangle{
		l:    10,
		b:    8,
	}
	s := square{
		side: 5,
	}
	display(r)
	display(s)
}
