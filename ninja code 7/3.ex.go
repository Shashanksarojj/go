package main

import "fmt"

func main() {

	defer foo()
	fmt.Println("hello")
}
func foo() {
	x := 11
	fmt.Println(x)
	defer func() {
		x := 14
		fmt.Println(x)

	}()

}
