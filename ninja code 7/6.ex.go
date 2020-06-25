package main

import (
	"fmt"
)

func main() {

	defer foo()

	func() {
		x := 43
		fmt.Println(x)
		fmt.Println("this ananomous")
	}()

}
func foo() {
	fmt.Println("this is foo")
}
