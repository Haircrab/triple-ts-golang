package game

import (
	"testing"
)

func TestInitGameState(t *testing.T) {
	gs := InitGameState()

	if len(&gs.board) != 3 {
		t.Errorf("len of board ROWS is not equal to 3: %v", len(&gs.board))
	}
	if len(&gs.board[0]) != 3 {
		t.Errorf("len of board ROWS is not equal to 3: %v", len(&gs.board[0]))
	}
	if len(&gs.board[0][0]) != 3 {
		t.Errorf("len of board cell is not equal to 3: %v", len(&gs.board[0][0]))
	}
}

func TestCheckWin(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer")
	}
	t.Log("p1: ", p1)

	mv, err := InitMove(0, 0, 0)
	if err != nil {
		t.Error("test fail in InitMove")
	}
	t.Log("mv: ", mv)

	res := gs.checkWin(*p1, *mv)
	t.Log("res: ", res)
}
