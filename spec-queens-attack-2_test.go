package queensAttack

import (
	"testing"
)

func TestQueensAttack2(t *testing.T) {
	attackableSquares := queensAttack(int32(5), int32(3), int32(4), int32(3), [][]int32{[]int32{5, 5}, []int32{4, 2}, []int32{2, 3}})
	if attackableSquares != 10 {
		t.Errorf("got %d attackable squares, need 10", attackableSquares)
	}
}
