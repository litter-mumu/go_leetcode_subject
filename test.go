package main

import "fmt"

func hello(i int) {
	fmt.Println(i)
}
func main() {
	var m = make(map[string]int) //A
	m["a"] = 1
	if v := m["b"]; v != 0 { //B
		fmt.Println(v)
	}
}
