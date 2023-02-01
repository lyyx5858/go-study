package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 1, 1, 2, 2, 3, 100}
	fmt.Println(topKFrequent(nums, 2))
}

//方法一：小顶堆

//构建小顶堆
type IHeap [][2]int //构建一个切片，每个切片是一个长度为2的数组

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1] //这里是以Value值进行排序的。Value(也就是出现频次)最小的，排在最上面! 如下面解释：

	/*
		由于push的时候是以key，Value放进来的，因此，h[0][0]：代表第0行的key, h[0][1]代表第0行的Value

		      Key[0]，Value[1]
		i=0    1		5
		i=1	   2		2
		i=2	   3		1
		i=3    ..		..
	*/

}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i] //h[i]表示第i个数组。
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int)) //如果断言成功，将会将这个容量为2的数组push到h
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}
	//记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, value := range map_num {
		heap.Push(h, [2]int{key, value}) // 官方的这个调用，原程序有两行，一个是h.Push，一个是up算法，
		// Push这个方法由调用的用户自己程序实现（本代码30行实现了它）。而up算法由官方程序实现。
		if h.Len() > k { //golang实现的是小堆树，这里是以出现频次也就是value排序的，因此，最后只剩下最大的k个数。
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}
