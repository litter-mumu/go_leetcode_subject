package main

import "fmt"

func main() {
	s := make([]int, 0, 1)
	s1 := append(s, 1)
	s2 := append(s, 2)
	s3 := append(s2, 1)
	fmt.Printf("value = %v,ptr = %T\n", s, &s)
	fmt.Printf("value = %v,ptr = %T\n", s1, &s1)
	fmt.Printf("value = %v,ptr = %T\n", s2, &s2)
	fmt.Printf("value = %v,ptr = %T\n", s3, &s3)
	//[] [2] [2] [2 1]
}
