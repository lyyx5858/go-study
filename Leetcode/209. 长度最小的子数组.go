package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))

}

func minSubArrayLen(target int, nums []int) int {
	i := 0                   // i:窗口左界
	l := len(nums)           // 数组长度
	sum := 0                 // 子数组之和
	result := l + 1          // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ { //j: 窗口右界
		sum += nums[j]
		for sum >= target { //当满足条件时，记录数组长度，然后开始向右滑动左窗口，直至不满足target的条件。
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++ //向右滑动左窗口
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}
