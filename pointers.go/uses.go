package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)

}
func foo(y *int) {
	fmt.Println("before", y)
	fmt.Println("before", *y)
	*y = 43
	fmt.Println("after", y)
	fmt.Println("after", *y)
}
