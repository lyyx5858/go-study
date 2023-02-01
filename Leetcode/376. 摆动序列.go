package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 7, 4, 9, 2, 5}
	nums2 := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	nums3 := []int{1, 2, 2, 2, 2, 2, 2, 2, 2}
	nums4 := []int{0, 0}
	nums5 := []int{2, 2, 2}
	fmt.Println(wiggleMaxLength(nums1))
	fmt.Println(wiggleMaxLength(nums2))
	fmt.Println(wiggleMaxLength(nums3))
	fmt.Println(wiggleMaxLength(nums4))
	fmt.Println(wiggleMaxLength(nums5))

}

func wiggleMaxLength(nums []int) int {

	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
		}
	}
	return ans
}
