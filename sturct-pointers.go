package main

import "fmt"

type Vertex3 struct {
	x int
	y int
}

func main() {
	v := Vertex3{1, 2}
	p := &v
	p.x = 1e9
	fmt.Println(v)
	fmt.Println(p.x)
}
