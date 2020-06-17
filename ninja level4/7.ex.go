package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("shouldn't print")
	case true:
		fmt.Print("yes")
	}
}
