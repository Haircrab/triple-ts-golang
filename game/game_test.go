package game

import (
	"testing"
)

func TestInitGameState(t *testing.T) {
	d := InitGameState()

	if len(&d.board) != 3 {
		t.Errorf("len of board ROWS is not equal to 3: %v", len(&d.board))
	}
	if len(&d.board[0]) != 3 {
		t.Errorf("len of board ROWS is not equal to 3: %v", len(&d.board[0]))
	}
	if len(&d.board[0][0]) != 3 {
		t.Errorf("len of board cell is not equal to 3: %v", len(&d.board[0][0]))
	}
}
