package main

import "fmt"

var a = 123.389
var b = 246.5
var c float64

func main() {
	Arithmetic()
	Relational()
}

func Arithmetic() {
	c = a + b
	fmt.Printf("%0.2f\n", c)
	a++
	fmt.Printf("%0.2f\n", a)
}

func Relational() {
	if a == b {
		fmt.Printf("a等于b")
	} else {
		fmt.Printf("a不等于b")
	}
}
