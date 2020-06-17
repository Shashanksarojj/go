package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":    32,
		"miss mad": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["barnabus"])

	v, ok := m["barnabus"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["miss mad"]; ok {
		fmt.Println("this print", v)
	}
}
