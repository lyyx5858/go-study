package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))

}

func threeSum(nums []int) [][]int {
	sort.Ints(nums) //排序
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ { //-2的原因是最少有3个数相加
		n1 := nums[i] //为了方便显示
		if n1 > 0 {   //因为已经排序了，因此n1>0,肯定无解。
			break
		}
		if i > 0 && n1 == nums[i-1] { //按题目要求，n1去重，注意是比较前一个数字，比如：{-1，-1，-1，0，1}，当走到第2个-1时，需要pass
			continue
		}
		l, r := i+1, len(nums)-1 //定义左、右指针
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 { //这个for是n2去重
					l++
				}
				for l < r && nums[r] == n3 { //这个for是n3去重
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++ //因为右指针已经是最大数了，所以移动左指针
			} else {
				r-- //因为左指针已经是最小数了（当然它是大于n1的），所以移动右指针
			}
		}
	}
	return res
}
