package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main()  {
	var a Abser
	f := Myfloat(-math.Sqrt2)
	v := Vertex9{3, 4}

	a = f
	a = &v

	fmt.Println(a.Abs())
}

type Myfloat float64

func (f Myfloat) Abs() float64  {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex9 struct {
	x, y float64
}

func (v *Vertex9) Abs() float64  {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
