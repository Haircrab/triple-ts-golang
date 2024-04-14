package game

import (
	"errors"
	"log"
)

const (
	N                   = 3
	ROWS, COLS, CIRCLES = N, N, N
	BOARD_DEFAULT_VAL   = -1
)

var (
	HOR    = [2][2]int{{0, 1}, {0, -1}}
	VER    = [2][2]int{{1, 0}, {-1, 0}}
	NEG_D  = [2][2]int{{1, 1}, {-1, -1}}
	POS_D  = [2][2]int{{-1, 1}, {1, -1}}
	PLANES = [4][2][2]int{HOR, VER, NEG_D, POS_D}
)

type GameState struct {
	Board            *boardState `json:"board"`
	PlayerSeq        [4]int      `json:"playerSeq"`
	NextPlayerSeqIdx int         `json:"nextPlayerSeqIdx"`

	WinnerIdx int `json:"winnerIdx"`
}
type (
	boardState     [ROWS][COLS]boardCellState
	boardCellState [CIRCLES]int // 0-5, 0 == unoccupied, 1-4 == player id
)

func InitBoardState() *boardState {
	bs := &boardState{{
		{-1, -1, -1},
		{-1, -1, -1},
		{-1, -1, -1},
	}, {
		{-1, -1, -1},
		{-1, -1, -1},
		{-1, -1, -1},
	}, {
		{-1, -1, -1},
		{-1, -1, -1},
		{-1, -1, -1},
	}}
	return bs
}

func InitGameState() *GameState {
	return &GameState{
		Board:            InitBoardState(),
		PlayerSeq:        [4]int{P1, P2, P3, P4},
		NextPlayerSeqIdx: 0,
		WinnerIdx:        -1,
	}
}

// win if true
func (gs *GameState) MakeMove(pyer *Player, mv Move) (bool, error) {
	log.Printf("make move in MakeMove %v", mv)
	if gs.WinnerIdx != -1 {
		return false, errors.New("Invalid move: the game is over")
	}
	if err := mv.CheckMove(); err != nil {
		return false, err
	}
	// check if move is valid
	// 1. if player is next player
	if !gs.CheckIsPlayerTurn(pyer) {
		return false, errors.New("Invalid move: it is not your turn dude")
	}
	// 2. if player have valid count of the circle using
	if !pyer.CanCircleUsed(mv.X) {
		return false, errors.New("Invalid move: no remaining selected circle")
	}
	// 3. if cell is occupied
	if gs.Board[mv.R][mv.C][mv.X] != BOARD_DEFAULT_VAL {
		return false, errors.New("Invalid move: cell is occupied")
	}

	// mutate board
	log.Printf("make move before mutCell%v", mv)
	gs.mutCell(mv.R, mv.C, mv.X, pyer.Id)
	gs.mutNextPlayer()

	// check winning condition
	if gs.checkWin(pyer, mv) {
		gs.WinnerIdx = pyer.Id

		return true, nil
	}

	return false, nil
}

// player's move should be already checked and made to the baord
// tranverse board and check winning conditions
func (gs *GameState) checkWin(pyer *Player, mv Move) bool {
	// board states
	currCell := gs.Board[mv.R][mv.C]

	// check if player occupied all slots in same cell
	if checkSameCell(pyer, currCell) {
		return true
	}

	// check all direction
	// TODO: seperate checking condition in parallel
	for _, plane := range PLANES {
		if checkCrossCells(pyer.Id, (gs.Board), plane, mv.R, mv.C, mv.X, N) {
			return true
		}
	}

	return false
}

func checkSameCell(pyer *Player, cell boardCellState) bool {
	for _, v := range cell {
		if v != pyer.Id {
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
