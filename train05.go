package main

import "fmt"

func main() {

	fmt.Println(count(1, 2, 3, add))
	fmt.Println(count(1, 2, 3, area))
	fmt.Println(count(1, 2, 3, sub))

}

func count(a, b, c int, math func(a, b, c int) int) int {
	return math(a, b, c)
}
func add(a, b, c int) int {
	return a + b + c
}

func area(a, b, c int) int {
	return a * b * c
}

func sub(a, b, c int) int {
	return a - b - c
}
