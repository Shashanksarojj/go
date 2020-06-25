package main

import (
	"fmt"
)

var g = func() { fmt.Println("from outside") }

func main() {
	f := foo()
	fmt.Println(f())

}
func foo() func() int {
	return func() int {
		return 45
	}
}
