package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("if this print mean it is wrong")
	case (2 == 4):
		fmt.Println("if this print mean it is wrong")
	case (3 == 3):
		fmt.Println("right")
		fallthrough
	case (4 == 4):
		fmt.Println("also, right")
		fallthrough
	case (7 == 9):
		fmt.Println("this is wronge but still print.")
		fallthrough
	case (11 == 13):
		fmt.Println("print because of fallthrough statement")

	}
}
