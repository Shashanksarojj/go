package main

import (
	"fmt"
)

func main() {
	fmt.Println("hellow world ! , how are you guys")
	loop()
}

func loop() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
