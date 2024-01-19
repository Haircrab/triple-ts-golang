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

// check same cell win
func TestCheckWinSameCell(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(4)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	mv, _ := InitMove(0, 0, 0)
	res, _ := gs.MakeMove(*p1, *mv)
	// t.Log(res)
	// t.Log(gs)

	mv, _ = InitMove(1, 1, 1)
	res, _ = gs.MakeMove(*p2, *mv)
	// t.Log(gs)

	mv, _ = InitMove(2, 2, 2)
	res, _ = gs.MakeMove(*p3, *mv)
	// t.Log(gs)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p4, *mv)
	t.Log(gs)

	//
	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p1, *mv)
	// t.Log(gs)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	// t.Log(gs)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p3, *mv)
	// t.Log(gs)

	mv, _ = InitMove(0, 1, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	t.Log(gs)

	//
	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p1, *mv)
	t.Log(res)
	t.Log(gs.winnerIdx)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

func TestCheckWinCrossCell(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(4)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	mv, _ := InitMove(0, 0, 0)
	res, _ := gs.MakeMove(*p1, *mv)
	// t.Log(res)
	// t.Log(gs)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)
	// t.Log(gs)

	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p3, *mv)
	// t.Log(gs)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	t.Log(gs)

	//
	mv, _ = InitMove(1, 1, 1)
	res, _ = gs.MakeMove(*p1, *mv)
	// t.Log(gs)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	// t.Log(gs)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)
	// t.Log(gs)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)

	//
	mv, _ = InitMove(2, 2, 2)
	res, _ = gs.MakeMove(*p1, *mv)
	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}
