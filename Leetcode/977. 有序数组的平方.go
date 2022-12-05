package main

import "fmt"

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))

}

func sortedSquares(nums []int) []int {
	l := len(nums)
	ans := make([]int, l)
	i, j := 0, l-1                      //i是左边指针，j是右边指针
	for pos := l - 1; pos >= 0; pos-- { //切片是从最右-->左填充, 从大往小填充，因为最大的不是在最左就是在最右
		if v, w := nums[i]*nums[i], nums[j]*nums[j]; v > w {
			ans[pos] = v
			i++ //左边指针移动一位
		} else {
			ans[pos] = w
			j-- //右边指针移动
		}
	}
	return ans
}
