package main

import (
	"fmt"
)

func main() {
	a := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Println(a)

	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(b)

	x := [][]string{a, b}
	fmt.Print(x)

	for i, g := range x {
		fmt.Println("record", i)
		for j, val := range g {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

}
