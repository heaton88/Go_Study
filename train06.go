package main

import "fmt"

func main() {
	fmt.Println(count(10, 20, 30, add))
}

func count(a, b, c int, math func(a, b, c int) int) int {
	return math(a, b, c)
}

func add(a, b, c int) int {
	return a + b + c
}

func sub(a, b, c int) int {
	return a - b - c
}

func area(a, b, c int) int {
	return a * b * c
}
