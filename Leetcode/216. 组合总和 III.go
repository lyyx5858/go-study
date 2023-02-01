package main

import "fmt"

func main() {

	fmt.Println(combinationSum3(9, 45))
}

//请先理解第77题。
func combinationSum3(k int, n int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs(k, n, 1, 0) //sum初始值为0
	return res
}

var (
	path []int
	res  [][]int
)

func dfs(k int, n int, start int, sum int) { //k是树的深度，n是树的宽度
	if len(path) == k { // 说明path的深度已经到达k了,开始检测和是否等于n
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp) //如果满足条件，请记录。不满足就直接返回。
		}
		return //整个dfs函数就这一个返回点。
	}
	for i := start; i <= 9; i++ { // 从start开始，不往回走，避免出现重复组合
		if i+sum > n || 9+1-i < k-len(path) { // 剪枝, sum到了减枝非常容易理解。但另外一个条件是：
			// 比如k=3,在第0层时，宽度循环里的8和9都不用循环了，因为只有两个数了，组成不成3个数。
			break
		}
		path = append(path, i)
		dfs(k, n, i+1, i+sum)
		path = path[:len(path)-1] //能走到这句，说明上面一句话已经走到了最大深度，必须回到上面节点，走其它树枝了。
	}
}
