package main

import "fmt"

func hello(num ...int) {
	num[0] = 18
}
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	t = append(t,10,10)
	fmt.Printf("%T,%v\n", t, t)
	fmt.Printf("len()=%v,cap()=%v", len(t), cap(t))
}
