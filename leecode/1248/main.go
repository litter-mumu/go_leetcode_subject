package main

import "fmt"

func main() {
	subarrays := numberOfSubarrays([]int{45627,50891,94884,11286,35337,46414,62029,20247,72789,89158,54203,79628,25920,16832,47469,80909}, 1)
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
