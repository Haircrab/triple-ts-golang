package game

import (
	"errors"
	"fmt"
)

const (
	// count
	N                   = 3
	ROWS, COLS, CIRCLES = N, N, N
)

var (
	HOR   = [2][2]int{{0, 1}, {0, -1}}
	VER   = [2][2]int{{1, 0}, {-1, 0}}
	NEG_D = [2][2]int{{1, 1}, {-1, -1}}
	POS_D = [2][2]int{{-1, 1}, {1, -1}}
	DIRS  = [4][2][2]int{HOR, VER, NEG_D, POS_D}
)

type gameState struct {
	board            boardState
	playerSeq        [4]int
	nextPlayerSeqIdx int
}
type (
	boardState     [ROWS][COLS]boardCellState
	boardCellState [CIRCLES]int // 0-5, 0 == unoccupied, 1-4 == player id
)

func InitGameState() *gameState {
	res := &gameState{
		playerSeq:        [4]int{p1, p2, p3, p4},
		nextPlayerSeqIdx: 0,
	}
	fmt.Println(res)
	return res
}

// win if true
func (gs *gameState) MakeMove(pyer player, mv move) (bool, error) {
	if err := checkMove(mv); err != nil {
		return false, err
	}
	// check if move is valid
	// 1. if player is next player
	if !gs.checkIsPlayerTurn(pyer) {
		return false, errors.New("Invalid move: it is not your turn dude")
	}
	// 2. if player have valid count of the circle using
	if pyer.canCircleUsed(mv.x) {
		return false, errors.New("Invalid move: no remaining selected circle")
	}
	// 3. if cell is occupied
	if gs.board[mv.r][mv.c][mv.x] != 0 {
		return false, errors.New("Invalid move: cell is occupied")
	}

	// mutate board
	gs.mutCell(mv.r, mv.c, mv.x, pyer.id)
	gs.mutNextPlayer()

	// check winning condition
	if gs.checkWin(pyer, mv) {
		return true, nil
	}

	return false, nil
}

// player's move should be already checked and made to the baord
// tranverse board and check winning conditions
// let move = r,c and player = Red
// C1: board[r][c][0] == board[r][c][1] == board[r][c][2] == Red
// C2: Diagonally, 2->1->0 == Red or 0->1->2 == Red
// C3: Vertically or Horizontally, 2->1->0 == Red or 0->1->2 == Red
func (gs *gameState) checkWin(pyer player, mv move) bool {
	mr, mc, mx := mv.r, mv.c, mv.x

	// board states
	currCell := gs.board[mr][mc]

	// check if player occupied all slots in same cell
	if checkSameCell(pyer, currCell) {
		return true
	}

	// check all direction
	if mr == mx || mc == mx || mr == N-1-mx || mc == N-1-mx {
		for _, dir := range DIRS {
			go func(dir [2][2]int) {
				fmt.Println("dir", dir)

				// dfs func
			}(dir)
		}
	}

	return false
}

func checkSameCell(pyer player, cell boardCellState) bool {
	// equals to N in NxN matrix
	for _, v := range cell {
		if v != pyer.id {
			return false
		}
	}
	return true
}

func checkCrossCells(pyer player, boardState boardState) bool {
	// if move is middle circle, need to check winning condition twice in all dirs
	// e.g. [1][1] is middle circle, then [0][0] and [2][2] can be (small, big) or (big, small)
	return false
}
