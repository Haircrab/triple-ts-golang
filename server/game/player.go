package game

import "errors"

type Player struct {
	Id      int          `json:"id"`
	Circles [CIRCLES]int `json:"circles"` // available circles, i0 == small, i1 == medium, i2 == large, value == remaining
	IsReady bool         `json:"isReady"`
}

// player index
const (
	MAX_PLAYER = 4
	p_none     = -1
	P1         = 0
	P2         = 1
	P3         = 2
	P4         = 3
)

func InitPlayer(id int) (*Player, error) {
	if id < P1 || id > P4 {
		return nil, errors.New("player index out of range")
	}

	res := &Player{
		Id:      id,
		Circles: [3]int{3, 3, 3},
	}
	return res, nil
}

func (p *Player) canCircleUsed(c int) bool {
	return !(p.Circles[c] <= 0)
}
