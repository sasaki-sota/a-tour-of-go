package main

import (
	"fmt"
	"math"
)

type Type3 struct {
	x, y float64
}

func Abs(v Type3) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Scale(v Type3, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := Type3{3, 4}
	Scale(v, 10)
	fmt.Println(Abs(v))
}
