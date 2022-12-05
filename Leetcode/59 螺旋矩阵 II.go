package main

import "fmt"

func main() {
	v := generateMatrix(3)
	fmt.Println(v)

}

func generateMatrix(n int) [][]int {
	var l, r, t, b = 0, n - 1, 0, n - 1
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}

	num, tar := 1, n*n
	for num <= tar {
		for i := l; i <= r; i++ {
			mat[t][i] = num // left to right.
			num++
		}
		t++

		for i := t; i <= b; i++ {
			mat[i][r] = num // top to bottom.
			num++
		}
		r--

		for i := r; i >= l; i-- {
			mat[b][i] = num // right to left.
			num++
		}
		b--

		for i := b; i >= t; i-- {
			mat[i][l] = num // bottom to top.
			num++
		}
		l++
	}
	return mat
}

/*
作者：Krahets
链接：https://leetcode.cn/problems/spiral-matrix-ii/solutions/12594/spiral-matrix-ii-mo-ni-fa-she-ding-bian-jie-qing-x/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
