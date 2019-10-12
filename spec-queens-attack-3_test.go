package queensAttack

import (
	"testing"
)

func TestQueensAttack3(t *testing.T) {
	attackableSquares := queensAttack(int32(1), int32(0), int32(1), int32(1), nil)
	if attackableSquares != 0 {
		t.Errorf("got %d attackable squares, need 0", attackableSquares)
	}
}
