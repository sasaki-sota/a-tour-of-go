package main

import "fmt"

type Vertex2 struct {
	x int
	y int
}

func main() {
	v := Vertex2{1, 2}
	v.x = 4
	fmt.Println(v.x)
}
