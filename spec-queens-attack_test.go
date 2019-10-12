package queensAttack

import (
	"testing"
)

func TestQueensAttack(t *testing.T) {
	attackableSquares := queensAttack(int32(4), int32(0), int32(4), int32(4), nil)
	if attackableSquares != 9 {
		t.Errorf("got %d attackable squares, need 9", attackableSquares)
	}
}
