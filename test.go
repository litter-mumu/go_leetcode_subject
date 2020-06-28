package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Right *TreeNode
	Val   int
	Left  *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var data []int
	traverse(root, &data)
	return data
}

func traverse(node *TreeNode, data *[]int) {
	if node == nil {
		return
	}
	*data = append(*data, node.Val)
	traverse(node.Left, data)
	traverse(node.Right, data)
}

func main() {
	//s := make([]int, 0, 1)
	//s1 := append(s, 1)
	//s2 := append(s, 2)
	//s3 := append(s2, 1)
	//fmt.Printf("value = %v,ptr = %T\n", s, &s)
	//fmt.Printf("value = %v,ptr = %T\n", s1, &s1)
	//fmt.Printf("value = %v,ptr = %T\n", s2, &s2)
	//fmt.Printf("value = %v,ptr = %T\n", s3, &s3)
	//[] [2] [2] [2 1]

	var v []int
	v = append(v, 11,12,12,3,13)
	fmt.Println(v)

	//reverseStr("this is my room")
}

func reverseStr(s string) {
	split := strings.Split(s, " ")
	for i := len(split); i > 0; i-- {
		fmt.Println(split[i-1])
	}
}
