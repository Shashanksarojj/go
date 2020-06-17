package main

import (
	"fmt"
)

func main() {
	sd := []string{"a", "b", "c", "d"}
	fmt.Println(sd)
	ds := []string{"e", "f", "g", "h"}
	fmt.Println(ds)

	sdd := [][]string{sd, ds}
	fmt.Println(sdd)
}
