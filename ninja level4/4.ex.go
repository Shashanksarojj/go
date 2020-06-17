package main

import (
	"fmt"
)

func main() {
	i := 1999
	for {

		if i > 2020 {
			break
		}
		fmt.Println(i)
		i++
	}
}
