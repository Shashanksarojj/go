package main

import "fmt"

func main() {
	s := foo()
	x, s1 := bar()

	fmt.Println(s, x, s1)
}

func foo() int {

	return 42

}

func bar() (int, string) {

	return 1998, "string"

}
