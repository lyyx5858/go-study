package main

import "fmt"

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	for i := 0; i < 9; i++ {
		fmt.Println(string(board[i]))
	}
}

func solveSudoku(board [][]byte) {
	var backtracking func(board [][]byte) bool
	backtracking = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				//判断此位置是否适合填数字
				if board[i][j] != '.' { //如果已有数字，不用填，直接跳过。
					continue
				}
				//尝试填1-9
				for k := '1'; k <= '9'; k++ {
					if isvalid(i, j, byte(k), board) == true { //如果满足要求就填
						board[i][j] = byte(k)
						if backtracking(board) == true { //每填一次，就检查是否成功，成功就返回！
							//这句递归函数在两层for循环中，所以叫“二维递归”
							return true
						}
						board[i][j] = '.' //到这里说明不成功，回退，继续在本格内试别的数字。
					}
				}
				return false // 到这里，说明本格1-9都试了，都不行，则回退到上格。然后在上格的循环中继续试另外一个数字。如果上格还不行，则回退到上上格。
			}
		}
		return true
	}
	backtracking(board)
}

//判断填入数字是否满足要求
func isvalid(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ { //行
		if board[row][i] == k {
			return false
		}
	}
	for i := 0; i < 9; i++ { //列
		if board[i][col] == k {
			return false
		}
	}
	//方格
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
