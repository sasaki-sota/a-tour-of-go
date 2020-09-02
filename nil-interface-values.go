package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	descride(i)
	i.M()
}

func descride(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
