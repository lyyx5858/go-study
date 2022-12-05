package main

import "fmt"

func main() {

	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}
	nums2 := []int{0, 6}
	r := findMedianSortedArrays(nums1, nums2)
	fmt.Println(r)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := append(nums1, nums2...)
	nums4 := QuickSort(nums3)
	l := len(nums4)
	fmt.Println(nums4, l, l/2)
	if isEven(l) {
		fmt.Println(nums4[l/2-1], nums4[l/2])
		return (float64(nums4[l/2-1]) + float64(nums4[l/2])) / 2
	} else {
		return float64(nums4[l/2])
	}
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitdata := arr[0]          //第一个数据
	low := make([]int, 0, 0)     //比我小的数组
	hight := make([]int, 0, 0)   //比我大的数组
	mid := make([]int, 0, 0)     //与我一样大的数组
	mid = append(mid, splitdata) //加入第一个数字

	for i := 1; i < len(arr); i++ {
		if arr[i] < splitdata {
			low = append(low, arr[i])
		} else if arr[i] > splitdata {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight)
	myarr := append(append(low, mid...), hight...)
	return myarr
}

func isEven(num int) bool { //判断奇偶
	if num%2 == 0 {
		return true
	}
	return false
}
