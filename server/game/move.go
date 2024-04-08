package game

import (
	"errors"
	"fmt"
)

type Move struct {
	R int `json:"r"`
	C int `json:"c"`
	X int `json:"x"` // index 0,1,2 -> 1-4 represent player id
}

func InitMove(r, c, x int) (*Move, error) {
	mv := &Move{
		R: r,
		C: c,
		X: x,
	}
	if err := mv.CheckMove(); err != nil {
		return nil, err
	}

	return mv, nil
}

func (gs *GameState) CheckIsPlayerTurn(pyer *Player) bool {
	return gs.PlayerSeq[gs.NextPlayerSeqIdx] == pyer.Id
}

func (mv *Move) CheckMove() error {
	if mv.R < 0 || mv.C < 0 || mv.X < 0 || mv.R >= ROWS || mv.C >= COLS || mv.X >= CIRCLES {
		return errors.New("Invalid move: move out of bound")
	}
	return nil
}

func (gs *GameState) mutNextPlayer() {
	if gs.NextPlayerSeqIdx == MAX_PLAYER-1 {
		gs.NextPlayerSeqIdx = 0
	} else {
		gs.NextPlayerSeqIdx++
	}
}

func (gs *GameState) mutCell(r, c, x, pyerIdx int) error {
	if x >= CIRCLES || x < 0 {
		return errors.New("circle value must be 0 to 2")
	}

	board := &(gs.Board)

	fmt.Printf("player %v making move in %v %v %v \n", pyerIdx, r, c, x)
	board[r][c][x] = pyerIdx
	return nil
}
