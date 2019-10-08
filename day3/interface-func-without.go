package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	getPerimeter() string
}

type circle struct {
	radius float64
	name string
}

func display(s shape) {
	fmt.Printf("Area is %0.1f\n", s.area())
}

func (c circle) area() float64 {
	return float64(math.Pi * c.radius * c.radius)
}

func main() {
	c := circle{
		radius: 5,
	}
	display(c)
}

// getDimension() is not implemented
//
// ERROR :
//# _/C_/Users/girish.warrier/go/src/hello
//.\interface-func-without.go:42:9: cannot use c (type circle) as type shape in argument to display:
//circle does not implement shape (missing getDimensions method)
