package main

import (
	"fmt"
)

func main() {
	x := "SDD"

	if x == "SSD" {
		fmt.Println(x)
	} else if x == " JAMES BOND" {
		fmt.Println("007")
	} else if x == "SDD" {
		fmt.Println("task is complete")
	} else {
		fmt.Println("wasted")
	}
}
