package main

import (
	"fmt"
)

func main() {
	favsport := "cricket"
	switch favsport {
	case "football":
		fmt.Println("this is football")
	case "badminton":
		fmt.Println("this is badminton")
	case "cricket":
		fmt.Println("this is IPL")
	}

}
