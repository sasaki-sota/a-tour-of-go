package main

import (
	"fmt"
	"math"
)

type Type5 struct {
	x, y float64
}

func (v *Type5) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *Type5) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := &Type5{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
