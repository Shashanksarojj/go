package main

import (
	"fmt"
)

func main() {

	// n:= bond  *this can also be used
	//instead of direct giving value to switch

	switch "bond" {
	case "wish":
		fmt.Println("this should not print")
	case "bond":
		fmt.Println("it print on given switch value")
	default:
		fmt.Println("this is default")
	}
}
