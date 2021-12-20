package main

import "fmt"

func main() {
	//	a := 10
	//	var b = &a
	//	fmt.Println(*b)
	//	*b = 100
	//	fmt.Println(a, *b, b)
	//
	var fruit string
	fmt.Scanf("%s", &fruit)

	switch fruit {
	case "apple":
		fmt.Println("u choose apple")
	case "banana":
		fmt.Println("u choose banana")
	case "orange":
		fmt.Println("u choose orange")
	default:
		fmt.Println("are u ill?")
	}
}
