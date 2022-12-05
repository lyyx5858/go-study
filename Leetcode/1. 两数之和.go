package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 13}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		if p, ok := m[target-v]; ok {
			return []int{i, p}
		} else {
			m[v] = i
		}
	}
	return []int{0, 0}
}
