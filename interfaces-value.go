package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	descride(i)
	i.M()
}

func descride(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
