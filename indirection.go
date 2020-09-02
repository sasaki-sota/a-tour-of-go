package main

import "fmt"

type Top11 struct {
	X, Y float64
}

func (v *Top11) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Top11, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Top11{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Top11{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
