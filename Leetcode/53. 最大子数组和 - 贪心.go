package main

import "fmt"

func main() {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {

	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 { //这里的nums[i-1]有两种情况：一种情况就是真实的nums[i-1],一种情况是前几个元素的和！
			//这里的逻辑是：如果前几个元素的和或者前一个元素是负，就抛弃掉，从头开始
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > maxSum { //这里的nums[i]也有两种情况：一种情况就是真实的nums[i],一种情况是前几个元素的和！
			maxSum = nums[i] //如果前几个元素+本次的元素的和或者只是本次元素的值大于历史时，就记录下来
		}
	}
	return maxSum

}
