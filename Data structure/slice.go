package main

import (
	"fmt"
)

func main() {
	// x := type ( vlue) //composite literal
	x := []int{4, 5, 6, 7, 8}
	fmt.Println(x)

	xi := []int{4, 5, 6, 7, 8}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}
