package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	//ここでfloatからuintへ型の変換
	var z = uint(f)
	fmt.Println(x, y, z)
	fmt.Printf("%T", z)
}
