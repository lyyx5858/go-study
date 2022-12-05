package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 8}
	nums2 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersection(nums1, nums2))

}

func intersection(nums1 []int, nums2 []int) []int {

	m1 := make(map[int]int)
	m2 := make(map[int]int)
	s := make([]int, 0)

	for _, v := range nums1 {
		m1[v]++ //先根据nums1建表
	}
	for _, v := range nums2 { //遍历时key在前，值在后
		_, ok := m1[v] //根据表2的值查表1查表时，值在前，OK在后
		if ok {
			m2[v]++         //每查到一次加1
			if m2[v] == 1 { //如果大于1，说明已经查到并放进切片了，如果小于1，说明没有
				s = append(s, v)
			}
		}
	}

	return s
}
