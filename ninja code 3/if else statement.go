package main

import (
	"fmt"
)

func main() {
	x := 42
	if x == 40 {
		fmt.Println("val 40")
	} else if x == 41 {
		fmt.Println("val 41")
	} else {
		fmt.Println("val 42")
	}

}
