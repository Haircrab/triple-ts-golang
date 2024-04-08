package game

import "sync"

type RoomCtx struct {
	GameState     *GameState `json:"gameState"`
	P1            *Player    `json:"p1"`
	P2            *Player    `json:"p2"`
	P3            *Player    `json:"p3"`
	P4            *Player    `json:"p4"`
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
		GameState:     gs,
		P1:            p1,
		P2:            p2,
		P3:            p3,
		P4:            p4,
		ConnPlayerIdx: [4]string{"", "", "", ""},
	}

	return res, nil
}

func (ctx *RoomCtx) SetPlayerConnId(playerIdx int, connId string) [4]bool {
	ctx.mu.Lock()
	ctx.ConnPlayerIdx[playerIdx] = connId
	readyState := [4]bool{ctx.P1.IsReady, ctx.P2.IsReady, ctx.P3.IsReady, ctx.P4.IsReady}
	ctx.mu.Unlock()

	return readyState
}

func (ctx *RoomCtx) ToggleReady(player *Player) {
	ctx.mu.Lock()
	player.IsReady = !player.IsReady
	if ctx.P1.IsReady && ctx.P2.IsReady && ctx.P3.IsReady && ctx.P4.IsReady {
		ctx.IsGameStarted = true
	} else {
		ctx.IsGameStarted = false
	}
	ctx.mu.Unlock()
}

func (ctx *RoomCtx) PlayerMakeMove(player *Player, mv *Move) (bool, error) {
	return ctx.GameState.MakeMove(player, *mv)
}

func (ctx *RoomCtx) OnPlayerDisconnect(playerConnId string) int {
	var res int
	for i, id := range ctx.ConnPlayerIdx {
		if id == playerConnId {
			res = i
			ctx.ConnPlayerIdx[i] = ""

			player := ctx.FindPlayerByIdx(i)
			player.IsReady = false
			ctx.IsGameStarted = false
		}
	}

	return res
}

func (ctx *RoomCtx) Copy() RoomCtx {
	return RoomCtx{
		GameState:     ctx.GameState,
		ConnPlayerIdx: ctx.ConnPlayerIdx,
		P1:            ctx.P1,
		P2:            ctx.P2,
		P3:            ctx.P3,
		P4:            ctx.P4,
		IsGameStarted: ctx.IsGameStarted,
	}
}

func (ctx *RoomCtx) FindPlayerByIdx(idx int) *Player {
	var player *Player
	switch idx {
	case P1:
		player = (ctx).P1
	case P2:
		player = (ctx).P2
	case P3:
		player = (ctx).P3
	case P4:
		player = (ctx).P4
	}

	return player
}
