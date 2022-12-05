package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(nums, target))

}

func search(nums []int, target int) int {
	l := len(nums)
	high := l - 1
	low := 0
	for low <= high { //注意是小于等于[low,high]
		i := (low + high) >> 1 //等同于 i=low+(high-low)/2
		if nums[i] < target {
			low++
		} else if nums[i] > target {
			high--
		} else {
			return i
		}
	}
	return -1
}
