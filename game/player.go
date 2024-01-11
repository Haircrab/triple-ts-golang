package game

type player struct {
	id int

	// available circles
	circles [3]int // available circles, i0 == small, i1 == medium, i2 == large, value == remaining
}

// player index
const (
	p_none int = iota
	p1
	p2
	p3
	p4
)

func InitPlayer(id int) *player {
	res := &player{
		id:      id,
		circles: [3]int{3, 3, 3},
	}
	return res
}
