package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(len(x))
}
