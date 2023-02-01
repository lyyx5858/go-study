package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	//nums := []int{3, 2, 1, 0, 4}
	//nums := []int{2, 0, 1, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	l := len(nums)
	if l == 1 {
		return true
	}
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] >= l-i { //如果倒数第二个元素大于等于数组长度与这个元素本身所在数组位置的差，则说明能够跳到。
			l = i //当能跳到，则将楼梯的长度更新为这个元素的位置。如果不行，则试倒数第三个元素，l不更新
		}
		if l == 1 { //如果楼梯的长度最终为1，说明第nums[i-1]可以满足条件，因此为true。
			return true
		}
	}
	return false
}
