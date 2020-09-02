package main

import (
	"fmt"
	"math"
)

type Type struct {
	x, y float64
}

func (v Type) Ads() float64  {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main()  {
	v := Type{3, 4}
	fmt.Println(v.Ads())
}
