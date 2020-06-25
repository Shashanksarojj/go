package main

import (
	"fmt"
)

var x int

func main() {

	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	g := foo()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

}
func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
