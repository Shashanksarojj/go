package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("the outer loop: %d\t the inner loop : %d\n ", i, j)
		}
	}
}
