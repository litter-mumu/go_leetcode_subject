package main

import "fmt"

/**
给你一个整数数组 nums 和一个整数 k。
如果某个 连续 子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。

请返回这个数组中「优美子数组」的数目。


示例 1：
输入：nums = [1,1,2,1,1], k = 3
输出：2
解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。

示例 2：
输入：nums = [2,4,6], k = 1
输出：0
解释：数列中不包含任何奇数，所以不存在优美子数组。

示例 3：
输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
输出：16
 

提示：

1 <= nums.length <= 50000
1 <= nums[i] <= 10^5
1 <= k <= nums.length

*/

func main() {
	subarrays := numberOfSubarrays([]int{45627, 50891, 94884, 11286, 35337, 46414, 62029, 20247, 72789, 89158, 54203, 79628, 25920, 16832, 47469, 80909}, 1)
	fmt.Println(subarrays)
}

func numberOfSubarrays(nums []int, k int) int {
	var subArrayCount = 0
	//先获取奇数的个数
	var oddCount = 0
	//再获取奇数对应的游标的切片
	oddIndex := make([]int, 0, len(nums))
	for i, v := range nums {
		if v%2 == 1 {
			oddCount++
			oddIndex = append(oddIndex, i)
		}
	}
	if oddCount < k {
		return 0
	}
	//oddIndex = [0,1,2,3,4]
	for i := 0; i < len(oddIndex); i++ {
		if i == 0 && oddCount == k {
			subArrayCount += (oddIndex[i] + 1) * (len(nums) - oddIndex[len(oddIndex)-1])
			return subArrayCount
		}
		if i == 0 {
			subArrayCount += (oddIndex[i] + 1) * (oddIndex[i+k] - oddIndex[i+k-1])
			continue
		}
		if i+k == oddCount {
			subArrayCount += (oddIndex[i] - oddIndex[i-1]) * (len(nums) - oddIndex[len(oddIndex)-1])
			return subArrayCount
		}
		subArrayCount += (oddIndex[i] - oddIndex[i-1]) * (oddIndex[i+k] - oddIndex[i+k-1])
	}
	return subArrayCount
}
