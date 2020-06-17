package main

import (
	"fmt"
)

func main() {
	// x := type ( vlue) //composite literal
	x := []int{4, 5, 6, 7, 8}
	fmt.Println(len(x))

	for i, v := range x {
		fmt.Println(i, v)
	}

}
