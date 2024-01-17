package game

import "errors"

type move struct {
	r, c int

	x int // index 0,1,2 -> 1-4 represent player id
}

func InitMove(r, c, x int) (*move, error) {
	mv := &move{
		r: r,
		c: c,
		x: c,
	}
	if err := checkMove(*mv); err != nil {
		return nil, err
	}

	return mv, nil
}

func (gs *gameState) checkIsPlayerTurn(pyer player) bool {
	return gs.playerSeq[gs.nextPlayerSeqIdx] == pyer.id
}

func checkMove(mv move) error {
	if mv.r < 0 || mv.c < 0 || mv.x < 0 || mv.r >= ROWS || mv.c >= COLS || mv.x >= CIRCLES {
		return errors.New("Invalid move: move out of bound")
	}
	return nil
}

func (gs *gameState) mutNextPlayer() {
	if gs.nextPlayerSeqIdx == MAX_PLAYER-1 {
		gs.nextPlayerSeqIdx = 0
	} else {
		gs.nextPlayerSeqIdx++
	}
}

func (gs *gameState) mutCell(r, c, x, pyerIdx int) error {
	if x >= CIRCLES || x < 0 {
		return errors.New("circle value must be 0 to 2")
	}

	gs.board[r][c][x] = pyerIdx
	return nil
}
