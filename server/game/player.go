package game

import "errors"

type player struct {
	Id      int          `json:"id"`
	Circles [CIRCLES]int `json:"circles"` // available circles, i0 == small, i1 == medium, i2 == large, value == remaining
	IsReady bool         `json:"isReady"`
}

// player index
const (
	MAX_PLAYER = 4
	p_none     = -1
	p1         = 0
	p2         = 1
	p3         = 2
	p4         = 3
)

func InitPlayer(id int) (*player, error) {
	if id < p1 || id > p4 {
		return nil, errors.New("player index out of range")
	}

	res := &player{
		Id:      id,
		Circles: [3]int{3, 3, 3},
	}
	return res, nil
}

func (p *player) canCircleUsed(c int) bool {
	return !(p.Circles[c] <= 0)
}
