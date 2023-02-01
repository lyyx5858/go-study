package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

var (
	path []int
	res  [][]int
)

func combine(n int, k int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs(n, k, 1)
	return res
}

func dfs(n int, k int, start int) { //k是树的深度，n是树的宽度， 这个程序中，for循环负责广度，递归负责深度！
	if len(path) == k { // 说明path的深度已经到达k了。,记录结果到res. 然后返回。这句话就是递归这个特殊循环的退出条件。
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
		return //整个dfs函数就这一个返回点。
	}
	for i := start; i <= n-(k-len(path))+1; i++ { // 从start开始，不往回走，避免出现重复组合

		/* 其中，n-(k-len(path))+1: 是剪枝操作，用于避免无效重复。
		可以理解为：
		path=0，意味着从最上层(第0层）长出树枝出来，举例，当n=4,k=2,则在第0层，最后一个数4，不需要长出枝了。
		path=1，从第一层长出树枝，当n=4,k=2,path=1,则需要从2，3，4都长出枝出来。
		path=2，path=k，则不会进到for，也就是不需要长出枝来。

		*/

		path = append(path, i)
		dfs(n, k, i+1)
		/*
			从本质上来讲，递归就是一 种特殊的循环，只不过，for循环是有显式的退出条件的。
			而递归的退出条件只能自己定义。
		*/

		path = path[:len(path)-1] //能走到这句，说明上面一句话已经走到了最大深度，必须回到上面节点，走其它树枝了。
	}
}
