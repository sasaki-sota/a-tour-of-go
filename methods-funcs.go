package main

import (
	"fmt"
	"math"
)

type Type1 struct {
	x, y float64
}

func Ads(v Type1) float64  {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main()  {
	v := Type1{3 ,4}
	fmt.Println(Ads(v))
}
