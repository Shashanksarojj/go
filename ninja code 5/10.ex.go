package main

import (
	"fmt"
)

func main() {
	m := map[string][]int{
		"James":    {33, 2, 2, 3, 3, 4, 4},
		"miss mad": {11, 3, 4, 4, 5, 6, 7, 9},
		"hero":     {1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	//fmt.Println(m)
	m["Bond"] = []int{2, 4, 7, 9, 0, 6, 5}

	delete(m, "hero")
	for k, v := range m {
		fmt.Println("this is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
