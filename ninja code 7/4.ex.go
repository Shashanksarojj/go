package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("i am", p.first, p.last, "and I am ", p.age, "years old.")
}

func main() {
	p1 := person{
		first: "shashank",
		last:  "dwivedi",
		age:   21,
	}
	p2 := person{
		first: "prashant",
		last:  "dwivedi",
		age:   24,
	}
	p3 := person{
		first: "saket",
		last:  "dwivedi",
		age:   29,
	}

	p1.speak()
	p2.speak()
	p3.speak()

}
