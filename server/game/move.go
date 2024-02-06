package game

import (
	"errors"
	"fmt"
)

type move struct {
	r, c int

	x int // index 0,1,2 -> 1-4 represent player id
}

func InitMove(r, c, x int) (*move, error) {
	mv := &move{
		r: r,
		c: c,
		x: x,
	}
	if err := checkMove(*mv); err != nil {
		return nil, err
	}

	return mv, nil
}

func (gs *gameState) checkIsPlayerTurn(pyer player) bool {
	return gs.PlayerSeq[gs.NextPlayerSeqIdx] == pyer.Id
}

func checkMove(mv move) error {
	if mv.r < 0 || mv.c < 0 || mv.x < 0 || mv.r >= ROWS || mv.c >= COLS || mv.x >= CIRCLES {
		return errors.New("Invalid move: move out of bound")
	}
	return nil
}

func (gs *gameState) mutNextPlayer() {
	if gs.NextPlayerSeqIdx == MAX_PLAYER-1 {
		gs.NextPlayerSeqIdx = 0
	} else {
		gs.NextPlayerSeqIdx++
	}
}

func (gs *gameState) mutCell(r, c, x, pyerIdx int) error {
	if x >= CIRCLES || x < 0 {
		return errors.New("circle value must be 0 to 2")
	}

	board := &(gs.Board)

	fmt.Printf("player %v making move in %v %v %v \n", pyerIdx, r, c, x)
	board[r][c][x] = pyerIdx
	return nil
}
