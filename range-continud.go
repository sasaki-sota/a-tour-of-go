package main

import "fmt"

func main() {
	pow4 := make([]int, 10)
	for i := range pow4 {
		pow4[i] = 1 << uint(i)
	}
	for _, value := range pow4 {
		fmt.Printf("%d\n", value)
	}
}
