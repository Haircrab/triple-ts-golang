package game

import "sync"

type RoomCtx struct {
	GameState     gameState  `json:"gameState"`
	P1            player     `json:"p1"`
	P2            player     `json:"p2"`
	P3            player     `json:"p3"`
	P4            player     `json:"p4"`
	ConnPlayerIdx [4]string  `json:"connPlayerIdx"`
	IsGameStarted bool       `json:"isGameStarted"`
	mu            sync.Mutex `json:"-"`
}

func InitRoomCtx() (*RoomCtx, error) {
	gs := InitGameState()
	p1, err := InitPlayer(0)
	if err != nil {
		return nil, err
	}
	p2, err := InitPlayer(1)
	if err != nil {
		return nil, err
	}
	p3, err := InitPlayer(2)
	if err != nil {
		return nil, err
	}
	p4, err := InitPlayer(3)
	if err != nil {
		return nil, err
	}

	res := &RoomCtx{
		GameState:     *gs,
		P1:            *p1,
		P2:            *p2,
		P3:            *p3,
		P4:            *p4,
		ConnPlayerIdx: [4]string{"", "", "", ""},
	}

	return res, nil
}
