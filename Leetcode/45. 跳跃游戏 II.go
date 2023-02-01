package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	//nums := []int{3, 2, 1, 0, 4}
	//nums := []int{2, 0, 1, 4}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}
	nextMaxLenth := 0
	curMaxLenth := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		nextMaxLenth = max(nextMaxLenth, nums[i]+i) // 记录下一次可以跳跃的最远距离；
		if i == curMaxLenth {                       // 当i走到上次的最远距离的时候
			count++                         //因为已经到了最远距离，因此跳数+1
			curMaxLenth = nextMaxLenth      // 再向前走记录的下一次可以走的最远距离
			if curMaxLenth >= len(nums)-1 { // 超过则退出返回
				break
			}
			nextMaxLenth = 0 // 当前具体走到记录的下一次能走的最远距离后。对next清空重新计数
		}

	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
