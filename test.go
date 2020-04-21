package main

import "fmt"

func main() {
	day06_1()
}

type MyInt1 int
type MyInt2 = int

func day06_1() {
	var i int = 0
	var i1 MyInt1 = MyInt1(11)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}
