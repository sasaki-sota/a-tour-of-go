package main

import "fmt"

//この場合はi = 1とj = 2となる
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
