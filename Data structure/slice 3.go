package main

import (
	"fmt"
)

func main() {

	x := []int{4, 5, 6, 7, 8}
	fmt.Println(x[1])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[1:4])

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
