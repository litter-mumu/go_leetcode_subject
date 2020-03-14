package main

import (
	"fmt"
)

/*
	// TODO 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
	你可以假设数组是非空的，并且给定的数组总是存在多数元素。

	示例 1:
	输入: [3,2,3]
	输出: 3

	示例 2:
	输入: [2,2,1,1,1,2,2]
	输出: 2
*/

func main() {
	var nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

//方法一
/*func majorityElement(nums []int) int {
	var maps map[int]int
	maps = make(map[int]int)
	for _, i := range nums {
		maps[i] +=1
		if maps[i] > (int)(len(nums)/2) {
			return i
		}
	}
	return 0
}*/

//方法二
/*func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[(len(nums)/2)]
}*/

//方法三	摩尔投票法
func majorityElement(nums []int) int {
	start := nums[0]
	count := 1
	for _, i := range nums {
		if i != start {
			count--
			if count == 0 {
				start = i
				count = 1
			}
		} else {
			count++
		}
	}
	return start
}
