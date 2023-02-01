package main

import (
	"fmt"
)

func main() {
	//a := []int{-7, -8, 7, 5, 7, 1, 6, 0}
	a := []int{1, 2, 3, 4, 7, 1, 6, 0}
	//a := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//a := []int{7, 4, 2}
	fmt.Println(maxSlidingWindow(a, 3))
}

// 封装单调队列的方式解题
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Push(val int) { //单调队列
	for !m.Empty() && val > m.Back() { //如果队列不是空并且新加入的值大于队尾元素的值，则：
		m.queue = m.queue[:len(m.queue)-1] //将队尾元素弹出，因为这是个循环，因此，如果新元素的值是在队列中最大，将最后只剩一个元素
	}
	m.queue = append(m.queue, val) //将新元素的值加入队尾
}

func (m *MyQueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	length := len(nums)
	res := make([]int, 0)
	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	// 记录前k个元素的最大值
	res = append(res, queue.Front())

	for i := k; i < length; i++ {
		// 滑动窗口移除最前面的元素
		queue.Pop(nums[i-k])
		// 滑动窗口添加最后面的元素
		queue.Push(nums[i])
		// 记录最大值
		res = append(res, queue.Front())
	}
	return res
}
