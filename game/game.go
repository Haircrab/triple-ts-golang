package game

import "fmt"

const (
	ROWS    = 3
	COLS    = 3
	CIRCLES = 3
)

type gameState struct {
	board boardState
}
type (
	boardState     [ROWS][COLS]boardCellState
	boardCellState [CIRCLES]int // 0-5, 0 == unoccupied
)

type move struct {
	r, c   int
	circle int
	player player
}

func InitGameState() *gameState {
	res := &gameState{}
	fmt.Println(res)
	return res
}

func (gs *gameState) Move(move *move, player *player) {
	// check if move is valid
	// 1. if player have valid count of the circle using
	// 2. if cell is occupied

	// mutate board

	// check winning condition
}

// player's move should be already checked and made to the baord
// tranverse board and check winning conditions
// let move = r,c and player = Red
// C2: Diagonally, 2->1->0 == Red or 0->1->2 == Red
// C3: Vertically or Horizontally, 2->1->0 == Red or 0->1->2 == Red
func CheckWin(gs *gameState, move *move, player *player) bool {
	mr, mc := move.r, move.c

	// board states
	b := &gs.board
	currCell := b[mr][mc]

	// check if player occupied all slots in same cell
	// C1: board[r][c][0] == board[r][c][1] == board[r][c][2] == Red
	if checkSameCell(&currCell, player) {
		return true
	}

	return false
}

func checkSameCell(cell *boardCellState, player *player) bool {
	for _, v := range cell {
		if v != player.id {
			return false
		}
	}
	return true
}
