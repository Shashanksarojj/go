package main

import (
	"fmt"
)

type lut int

var x lut

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Println(x)

}
