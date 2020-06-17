package main

import (
	"fmt"
)

func main() {
	for i := 10; i < 101; i++ {
		r := i % 4
		fmt.Println(i, r)
	}
}
