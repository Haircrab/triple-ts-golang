package game

import "errors"

type player struct {
	id int

	circles [CIRCLES]int // available circles, i0 == small, i1 == medium, i2 == large, value == remaining
}

// player index
const (
	MAX_PLAYER = 4
	p_none     = 0
	p1         = 1
	p2         = 2
	p3         = 3
	p4         = 4
)

func InitPlayer(id int) (*player, error) {
	if id < p1 || id > p4 {
		return nil, errors.New("player index out of range")
	}

	res := &player{
		id:      id,
		circles: [3]int{3, 3, 3},
	}
	return res, nil
}

func (p *player) canCircleUsed(c int) bool {
	return !(p.circles[c] <= 0)
}
