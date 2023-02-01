package main

import (
	"fmt"
	"sort"
)

func main() {

	//nums := []int{2, -3, -1, 5, -4}
	//nums := []int{3, -1, 0, 2}
	nums := []int{-8, 3, -5, -3, -5, -2}
	//nums := []int{-2, 3}
	//nums := []int{-2, 9, 9, 8, 4}
	fmt.Println(largestSumAfterKNegations(nums, 6))

}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums) // 先排序
	var sum int
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 { //将所有能变成正的变成正
			nums[i] = -nums[i]
			k--
		}

		sum = sum + nums[i]
	}

	sort.Ints(nums) //再排一次序

	if !isEven(k) {
		//如果剩下的K是奇数，则说明K一定大于0，则说明数组中所有的数已经变正了。
		// 这时需要将最小的一个正数变负，以消耗掉K，而这个正数已经加过一次了，因此需要减掉两次
		sum = sum - 2*nums[0]
	}

	return sum
}

func isEven(num int) bool { //判断奇偶
	if num%2 == 0 {
		return true
	}
	return false
}
