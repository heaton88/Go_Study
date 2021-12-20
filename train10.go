package main

import (
	"fmt"
)

func main() {
	var s1 []int
	//s1 = make([]int, 3)
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	//s1 = append(s1, 4)
	//printSlice(s1)
	s2 := make([]int, 3, 3)
	s2 = append(s2, 100)
	s2 = append(s2, 101)
	s2 = append(s2, 101)
	s2 = append(s2, 101)
	//s2 = append(s2, 101)
	//printSlice(s2)

	s3 := make([]string, 3, 3)
	//printSlice(s3)
	fmt.Printf(s3)
}

//func printSlice(s1 []int) {
//	fmt.Printf("len = %d, cap = %d, %v\n", len(s1), cap(s1), s1)
//}
