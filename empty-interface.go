package main

import "fmt"

func descride(i interface{})  {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main()  {
	var i interface{}
	descride(i)

	i = 42
	descride(i)

	i = "hello"
	descride(i)
}
