package game

import (
	"log"

	"triple-ts-golang/game"

	socketio "github.com/googollee/go-socket.io"
)

func createRoom(conn *socketio.Conn, rs *map[string]*game.RoomCtx) error {
	defer func() {
		for k, v := range *rs {
			log.Printf("roomStates[%v]: %v \n", k, v)
		}
	}()

	if ctx, err := game.InitRoomCtx(); err != nil {
		return err
	} else {
		connId := (*conn).ID()
		(*rs)[connId] = ctx

		ctx.SetPlayerConnId(0, connId)

		(*conn).Emit(createRoomOkEvent, CreateRoomOkEventRes{
			RoomId:   connId,
			PlayerId: 0,
		})
		(*conn).SetContext(ConnCtx{
			RoomId:   connId,
			PlayerId: 0,
		})
		return nil
	}
}
