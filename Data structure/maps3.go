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

	delete(m, "James")
	fmt.Println(m)

	delete(m, "miss horr")
	fmt.Println(m)

	fmt.Println(m["miss mad"])
	fmt.Println(m["miss horr"])

	if v, ok := m["miss mad "]; ok {
		fmt.Println("value:", v)
		delete(m, "miss mad")
	}

	fmt.Println(m)

}
