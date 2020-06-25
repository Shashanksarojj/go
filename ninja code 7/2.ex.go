package main

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7}
	n := foo(ii...)
	println(n)

	ii2 := []int{1, 2, 3, 4, 5, 6, 7}
	n1 := foo(ii2...)
	println(n1)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v

	}
	return total
}
func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v

	}
	return total

}
