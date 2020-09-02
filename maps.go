package main

import "fmt"

type Vertex7 struct {
	Lat, Long float64
}

var m map[string]Vertex7

func main() {
	m = make(map[string]Vertex7)
	m["Bell Labs"] = Vertex7{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
