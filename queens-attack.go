package queensAttack

// queensAttack has the following parameters:
// - n: an integer, the number of rows and columns in the board
// - k: an integer, the number of obstacles on the board
// - r_q: integer, the row number of the queen's position
// - c_q: integer, the column number of the queen's position
// - obstacles: a two dimensional array of integers where each element is an array of  integers, the row and column of an obstacle

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	var compass = [][]int32{[]int32{1, 0}, []int32{1, 1}, []int32{0, 1}, []int32{-1, 1}, []int32{-1, 0}, []int32{-1, -1}, []int32{0, -1}, []int32{1, -1}}
	var attackableSquares int32

	for _, direction := range compass {
		row := r_q
		column := c_q

		for {
			row, column = calcNextSquare(row, column, direction)
			if onBoard(row, column, n) {
				if pathIsClear(row, column, obstacles) {
					attackableSquares++
				} else {
					break
				}
			} else {
				break
			}
		}
	}

	return attackableSquares
}

func calcNextSquare(currentRow int32, currentColumn int32, direction []int32) (row int32, column int32) {
	return currentRow + direction[0], currentColumn + direction[1]
}

func onBoard(row int32, column int32, n int32) bool {
	if row <= n && row > 0 && column <= n && column > 0 {
		return true
	}

	return false
}

func pathIsClear(row int32, column int32, obstacles [][]int32) bool {
	if len(obstacles) > 0 {
		for _, arr := range obstacles {
			if arr[0] == row && arr[1] == column {
				return false
			}
		}
	}

	return true
}
