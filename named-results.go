package main

import "fmt"

//naked" return = returnの後に何も描かないこと
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
