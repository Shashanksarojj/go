package main

import (
	"fmt"
)

func main() {

	x := 2
	fmt.Printf("%d\t\t%b", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)
}
