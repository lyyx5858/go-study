package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))

}

func removeElement(nums []int, val int) int {
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...) //将i所在元素剔除，将i左边和i右边的元素拼接起来
			i--
			fmt.Printf("%p\n", nums)
		}
	}
	fmt.Println(nums)
	return len(nums)
}

func removeElement01(nums []int, val int) int {
	l := len(nums)
	res := 0
	for i := 0; i < l; i++ {
		if nums[i] != val {
			nums[res] = nums[i] //如果当前元素不等于val，就向前填充
			res++               //上面填充完成后，位置加1
		}
	}
	nums = nums[:res] //res是最后的位置，也是长度值
	return res

}
