package main

import "fmt"

type Vertex4 struct {
	x, y int
}

var (
	v1 = Vertex4{1, 2}
	v2 = Vertex4{x: 1}
	v3 = Vertex4{}
	p  = &Vertex4{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
