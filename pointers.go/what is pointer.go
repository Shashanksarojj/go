package main

import (
	"fmt"
)

func main() {
	x := 4
	fmt.Println("x")
	fmt.Println(&x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x)

	b := &x
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*&x)

}
