package main

import (
	"fmt"
	"time"
)

func main() {

	// 记录程序开始时间
	startTime := time.Now()

	s := solveNQueens(5)
	fmt.Println(s)
	fmt.Println(res)
	fmt.Println(len(res))

	// 记录程序结束时间
	endTime := time.Now()

	// 计算程序执行时间
	elapsedTime := endTime.Sub(startTime)

	// 输出程序执行时间
	fmt.Printf("Total elapsed time: %s\n", elapsedTime)

}

var (
	path []int
	res  [][]int
)

func solveNQueens(n int) [][]string {
	path, res = make([]int, 0), make([][]int, 0)

	dfs(n)
	return resPrint(res, n)
}

func dfs(n int) {
	if len(path) == n { //这是垂直循环的退出条件
		tmp := make([]int, n)
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for j := 1; j <= n; j++ { //水平循环 ：1 to n

		if len(path) > 0 {

			valid := true
			for i := 0; i < len(path); i++ { //当走到第3层时，需要将第1-2层都检查一遍，因此需要一个循环
				if j == path[i] || j == path[i]+i-len(path) || j == path[i]-i+len(path) {
					//第一是判断垂直， 第二个判断135度斜线。 第三个是判断45度斜线
					valid = false
					break //这个是break这个小循环。
				}
			}
			if valid == false { //valid 为false则表明这个j不能被记录到path中。
				continue
			}

		}

		path = append(path, j)    //记录成功后，则继续下一层。
		dfs(n)                    //内循环，也叫深度循环， 这个递归在一层for循环中，叫一维递归。
		path = path[:len(path)-1] //当走到这句，说明第J列结束，开始进入下一个J，这时需要将最近一次加到path上的J弹出来）
	}
}

func resPrint(res [][]int, n int) [][]string { //用来将数字结果转换成字符串结果
	s := make([][]string, len(res))
	for i := 0; i < len(res); i++ {
		s[i] = make([]string, n)
	}
	tmp := make([]byte, n)

	for i := 0; i < len(res); i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				tmp[k] = '.'
			}
			tmp[res[i][j]-1] = 'Q'
			s[i][j] = string(tmp)

		}
	}

	return s
}
