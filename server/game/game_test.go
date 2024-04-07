package game

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestInitGameState(t *testing.T) {
	gs := InitGameState()

	if len(&gs.Board) != 3 {
		t.Errorf("len of board ROWS is not equal to 3: %v", len(&gs.Board))
	}
	if len(&gs.Board[0]) != 3 {
		t.Errorf("len of board ROWS is not equal to 3: %v", len(&gs.Board[0]))
	}
	if len(&gs.Board[0][0]) != 3 {
		t.Errorf("len of board cell is not equal to 3: %v", len(&gs.Board[0][0]))
	}
}

// check same cell winning condition
func TestCheckWinSameCell(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	mv, _ := InitMove(0, 0, 0)
	res, _ := gs.MakeMove(*p1, *mv)

	mv, _ = InitMove(1, 1, 1)
	res, _ = gs.MakeMove(*p2, *mv)

	mv, _ = InitMove(2, 2, 2)
	res, _ = gs.MakeMove(*p3, *mv)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p4, *mv)
	t.Log(gs)

	//
	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p1, *mv)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p3, *mv)

	mv, _ = InitMove(0, 1, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	t.Log(gs)

	//
	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p1, *mv)
	t.Log(res)
	t.Log(gs.WinnerIdx)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

// check cross cell winning condition - ascending/ decending order
// negative diagonal
func TestCheckWinCrossCellStartNEG_D(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
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

func TestCheckWinCrossCellMidNEG_D(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	// first round
	mv, _ := InitMove(1, 1, 1)
	res, _ := gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)
	t.Log(gs)

	// second round
	mv, _ = InitMove(0, 0, 0)
	res, _ = gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)

	// third round
	mv, _ = InitMove(2, 2, 2)
	res, _ = gs.MakeMove(*p1, *mv)

	assert.Equal(t, true, res)

	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

func TestCheckWinCrossCellEndNEG_D(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	// first round
	mv, _ := InitMove(2, 2, 2)
	res, _ := gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)
	t.Log(gs)

	// second round
	mv, _ = InitMove(0, 0, 0)
	res, _ = gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)

	// third round
	mv, _ = InitMove(1, 1, 1)
	res, _ = gs.MakeMove(*p1, *mv)

	assert.Equal(t, true, res)

	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

// positive diagonal
func TestCheckWinCrossCellStartPOS_D(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	mv, _ := InitMove(0, 2, 0)
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
	mv, _ = InitMove(2, 0, 2)
	res, _ = gs.MakeMove(*p1, *mv)
	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

func TestCheckWinCrossCellMidPOS_D(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	// first round
	mv, _ := InitMove(1, 1, 1)
	res, _ := gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)
	t.Log(gs)

	// second round
	mv, _ = InitMove(0, 2, 2)
	res, _ = gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)

	// third round
	mv, _ = InitMove(2, 0, 0)
	res, _ = gs.MakeMove(*p1, *mv)

	assert.Equal(t, true, res)

	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

func TestCheckWinCrossCellEndPOS_D(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	// first round
	mv, _ := InitMove(2, 0, 0)
	res, _ := gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)
	t.Log(gs)

	// second round
	mv, _ = InitMove(0, 2, 2)
	res, _ = gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)

	// third round
	mv, _ = InitMove(1, 1, 1)
	res, _ = gs.MakeMove(*p1, *mv)

	assert.Equal(t, true, res)

	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

// horizontal
func TestCheckWinCrossCellStartHOR(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	mv, _ := InitMove(0, 0, 0)
	res, _ := gs.MakeMove(*p1, *mv)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)

	mv, _ = InitMove(0, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)

	//
	mv, _ = InitMove(0, 2, 2)
	res, _ = gs.MakeMove(*p1, *mv)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)

	//
	mv, _ = InitMove(0, 1, 1)
	res, _ = gs.MakeMove(*p1, *mv)
	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

func TestCheckWinCrossCellMidHOR(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	mv, _ := InitMove(0, 1, 1)
	res, _ := gs.MakeMove(*p1, *mv)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)

	mv, _ = InitMove(0, 2, 2)
	res, _ = gs.MakeMove(*p3, *mv)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)

	//
	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p1, *mv)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)

	//
	mv, _ = InitMove(0, 2, 0)
	res, _ = gs.MakeMove(*p1, *mv)
	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

func TestCheckWinCrossCellEndHOR(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	// first round
	mv, _ := InitMove(2, 2, 0)
	res, _ := gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)
	t.Log(gs)

	// second round
	mv, _ = InitMove(2, 0, 2)
	res, _ = gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)

	// third round
	mv, _ = InitMove(2, 1, 1)
	res, _ = gs.MakeMove(*p1, *mv)

	assert.Equal(t, true, res)

	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

// vertical
func TestCheckWinCrossCellEndVER(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	// first round
	mv, _ := InitMove(2, 2, 0)
	res, _ := gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)
	t.Log(gs)

	// second round
	mv, _ = InitMove(0, 2, 2)
	res, _ = gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)

	// third round
	mv, _ = InitMove(1, 2, 1)
	res, _ = gs.MakeMove(*p1, *mv)

	assert.Equal(t, true, res)

	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

// check cross cell winning condition - circle with same size
func TestCheckWinCrossCellStartNEG_D2(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	mv, _ := InitMove(0, 0, 0)
	res, _ := gs.MakeMove(*p1, *mv)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)

	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p3, *mv)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)

	//
	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p1, *mv)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)

	//
	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p1, *mv)
	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

func TestCheckWinCrossCellMidNEG_D2(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	// first round
	mv, _ := InitMove(1, 1, 1)
	res, _ := gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)

	// second round
	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 2)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)

	// third round
	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p1, *mv)
	t.Logf("res: %v", res)
	assert.Equal(t, true, res)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

func TestCheckWinCrossCellEndNEG_D2(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	// first round
	mv, _ := InitMove(2, 2, 2)
	res, _ := gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)
	t.Log(gs)

	// second round
	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 1, 0)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)

	// third round
	mv, _ = InitMove(1, 1, 2)
	res, _ = gs.MakeMove(*p1, *mv)

	assert.Equal(t, true, res)

	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

// positive diagonal
func TestCheckWinCrossCellStartPOS_D2(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	mv, _ := InitMove(0, 2, 0)
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
	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p1, *mv)
	// t.Log(gs)

	mv, _ = InitMove(1, 1, 1)
	res, _ = gs.MakeMove(*p2, *mv)
	// t.Log(gs)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)
	// t.Log(gs)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)

	//
	mv, _ = InitMove(2, 0, 0)
	res, _ = gs.MakeMove(*p1, *mv)
	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

func TestCheckWinCrossCellMidPOS_D2(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	// first round
	mv, _ := InitMove(1, 1, 1)
	res, _ := gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)
	t.Log(gs)

	// second round
	mv, _ = InitMove(0, 2, 1)
	res, _ = gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)

	// third round
	mv, _ = InitMove(2, 0, 1)
	res, _ = gs.MakeMove(*p1, *mv)

	assert.Equal(t, true, res)

	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

// horizontal
func TestCheckWinCrossCellStartHOR2(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	mv, _ := InitMove(0, 0, 0)
	res, _ := gs.MakeMove(*p1, *mv)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)

	mv, _ = InitMove(0, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)

	//
	mv, _ = InitMove(0, 2, 0)
	res, _ = gs.MakeMove(*p1, *mv)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)

	//
	mv, _ = InitMove(0, 1, 0)
	res, _ = gs.MakeMove(*p1, *mv)
	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

func TestCheckWinCrossCellMidHOR2(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	mv, _ := InitMove(0, 1, 1)
	res, _ := gs.MakeMove(*p1, *mv)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)

	mv, _ = InitMove(0, 2, 2)
	res, _ = gs.MakeMove(*p3, *mv)

	mv, _ = InitMove(2, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)

	//
	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p1, *mv)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)

	//
	mv, _ = InitMove(0, 2, 1)
	res, _ = gs.MakeMove(*p1, *mv)
	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}

func TestCheckWinCrossCellEndHOR2(t *testing.T) {
	gs := InitGameState()
	t.Log("gs: ", gs)

	p1, err := InitPlayer(0)
	if err != nil {
		t.Error("test fail in InitPlayer 1")
	}
	p2, err := InitPlayer(1)
	if err != nil {
		t.Error("test fail in InitPlayer 2")
	}
	p3, err := InitPlayer(2)
	if err != nil {
		t.Error("test fail in InitPlayer 3")
	}
	p4, err := InitPlayer(3)
	if err != nil {
		t.Error("test fail in InitPlayer 4")
	}

	// first round
	mv, _ := InitMove(2, 2, 0)
	res, _ := gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 1)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(0, 0, 2)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 2, 0)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)
	t.Log(gs)

	// second round
	mv, _ = InitMove(2, 0, 0)
	res, _ = gs.MakeMove(*p1, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(1, 1, 0)
	res, _ = gs.MakeMove(*p2, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 2, 1)
	res, _ = gs.MakeMove(*p3, *mv)
	assert.Equal(t, false, res)

	mv, _ = InitMove(2, 1, 2)
	res, _ = gs.MakeMove(*p4, *mv)
	assert.Equal(t, false, res)

	// third round
	mv, _ = InitMove(2, 1, 0)
	res, _ = gs.MakeMove(*p1, *mv)

	assert.Equal(t, true, res)

	t.Log(res)
	t.Log(gs)

	mv, _ = InitMove(1, 0, 2)
	res, err = gs.MakeMove(*p2, *mv)
	if err != nil {
		t.Log(err)
	}
}
