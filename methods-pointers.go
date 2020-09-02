package main

import (
	"fmt"
	"math"
)

type Type2 struct {
	x, y float64
}

func (v Type2) Abs() float64  {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Type2) Scale(f float64)  {
	v.x = v.x * f
	v.y = v.y * f
}

func main()  {
	v := Type2{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
