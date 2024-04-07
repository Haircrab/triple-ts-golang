package game

import (
	"errors"
	"log"

	"triple-ts-golang/game"

	socketio "github.com/googollee/go-socket.io"
)

func joinRoom(rid string, ser *socketio.Server, conn *socketio.Conn, rs *map[string]*game.RoomCtx) error {
	defer func() {
		for k, v := range *rs {
			log.Printf("roomStates[%v]: %v \n", k, v)
		}
	}()

	var err error
	connID := (*conn).ID()

	if err = findRoomExist(ser.Rooms(ns), rid); err != nil {
		(*conn).Emit(errorEvent, ErrorEventRes{
			Message: err.Error(),
		})
		(*conn).Close()
		return err
	}

	// check if the room is full
	if ser.RoomLen(ns, rid) >= maxConn {
		(*conn).Emit(errorEvent, ErrorEventRes{
			Message: "The room is full already",
		})
		(*conn).Close()
		return err
	}

	(*conn).LeaveAll()
	(*conn).Join(rid)

	var playerId int
	ctx, ok := (*rs)[rid]
	log.Printf("ctx: %v, ok: %v \n", ctx, ok)

	if ok {
		for i, v := range ctx.ConnPlayerIdx {
			if v == "" {
				playerId = i
				break
			}
		}

		(*conn).SetContext(ConnCtx{
			RoomId:   rid,
			PlayerId: playerId,
		})

		readyState := ctx.SetPlayerConnId(playerId, connID)

		//* Join room
		(*conn).Emit(joinRoomOkEvent, JoinRoomOkEventRes{
			RoomId:     rid,
			PlayerId:   playerId,
			ReadyState: readyState,
			RoomState:  ctx.Copy(),
		})

		return err
	} else {
		log.Println("room not found")
		err = errors.New("room not found")
		(*conn).Emit(errorEvent, ErrorEventRes{
			Message: err.Error(),
		})
		(*conn).Close()
		return err
	}
}
