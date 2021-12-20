package main

import "fmt"

func main() {
	nextnumber := getsequence()
	fmt.Println(nextnumber())
	fmt.Println(nextnumber())
	fmt.Println(nextnumber())
	fmt.Println(nextnumber())
}
func getsequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
