package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 7, 8}
	fmt.Println(x)

	x = append(x, 45, 46, 67, 86, 98, 45)
	fmt.Println(x)

	y := []int{234, 453, 451, 235, 768}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
