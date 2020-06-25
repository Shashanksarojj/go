package main

import (
	"fmt"
)

var g = func() { fmt.Println("from outside") }

func main() {
	s1 := func() {
		for i := 0; i < 4; i++ {
			fmt.Println(i)
		}
	}
	g()
	s1()

}
