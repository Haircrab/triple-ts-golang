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
	HOR    = [2][2]int{{0, 1}, {0, -1}}
	VER    = [2][2]int{{1, 0}, {-1, 0}}
	NEG_D  = [2][2]int{{1, 1}, {-1, -1}}
	POS_D  = [2][2]int{{-1, 1}, {1, -1}}
	PLANES = [4][2][2]int{HOR, VER, NEG_D, POS_D}
)

type gameState struct {
	board            boardState
	playerSeq        [4]int
	nextPlayerSeqIdx int

	winnerIdx int
}
type (
	boardState     [ROWS][COLS]boardCellState
	boardCellState [CIRCLES]int // 0-5, 0 == unoccupied, 1-4 == player id
)

func InitGameState() *gameState {
	res := &gameState{
		playerSeq:        [4]int{p1, p2, p3, p4},
		nextPlayerSeqIdx: 0,
		winnerIdx:        -1,
	}
	fmt.Println(res)
	return res
}

// win if true
func (gs *gameState) MakeMove(pyer player, mv move) (bool, error) {
	if gs.winnerIdx != -1 {
		return false, errors.New("Invalid move: the game is over")
	}
	if err := checkMove(mv); err != nil {
		return false, err
	}
	// check if move is valid
	// 1. if player is next player
	if !gs.checkIsPlayerTurn(pyer) {
		return false, errors.New("Invalid move: it is not your turn dude")
	}
	// 2. if player have valid count of the circle using
	if !pyer.canCircleUsed(mv.x) {
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
		gs.winnerIdx = pyer.id

		return true, nil
	}

	return false, nil
}

// player's move should be already checked and made to the baord
// tranverse board and check winning conditions
func (gs *gameState) checkWin(pyer player, mv move) bool {
	// board states
	currCell := gs.board[mv.r][mv.c]

	// check if player occupied all slots in same cell
	if checkSameCell(pyer, currCell) {
		return true
	}

	// check all direction
	// TODO: seperate checking condition in parallel
	for _, plane := range PLANES {
		if checkCrossCells(pyer.id, &(gs.board), plane, mv.r, mv.c, mv.x, N) {
			return true
		}
	}

	return false
}

func checkSameCell(pyer player, cell boardCellState) bool {
	for _, v := range cell {
		if v != pyer.id {
			return false
		}
	}
	return true
}

func checkCrossCells(pyerId int, boardState *boardState, plane [2][2]int, r, c, x int, targetAcc int) bool {
	dir1, dir2 := plane[0], plane[1]

	p1a := dfs(pyerId, boardState, dir1, r+dir1[0], c+dir1[1], x+1, 1, 1)
	if p1a == targetAcc {
		return true
	}
	p1b := dfs(pyerId, boardState, dir2, r+dir2[0], c+dir2[1], x-1, -1, 1)
	if p1b == targetAcc {
		return true
	}
	if p1a+p1b-1 == targetAcc {
		return true
	}

	p2a := dfs(pyerId, boardState, dir1, r+dir1[0], c+dir1[1], x-1, -1, 1)
	if p2a == targetAcc {
		return true
	}
	p2b := dfs(pyerId, boardState, dir2, r+dir2[0], c+dir2[1], x+1, 1, 1)
	if p2b == targetAcc {
		return true
	}
	if p2a+p2b-1 == targetAcc {
		return true
	}

	p3a := dfs(pyerId, boardState, dir1, r+dir1[0], c+dir1[1], x, 0, 1)
	if p3a == targetAcc {
		return true
	}
	p3b := dfs(pyerId, boardState, dir2, r+dir2[0], c+dir2[1], x, 0, 1)
	if p3b == targetAcc {
		return true
	}
	if p3a+p3b-1 == targetAcc {
		return true
	}

	return false
}

func dfs(pyerId int, boardState *boardState, dir [2]int, r, c, x int, modifier int, acc int) int {
	if r >= ROWS || c >= COLS || x >= CIRCLES || r < 0 || c < 0 || x < 0 || boardState[r][c][x] != pyerId {
		return acc
	}

	return dfs(pyerId, boardState, dir, r+dir[0], c+dir[1], x+modifier, modifier, acc+1)
}