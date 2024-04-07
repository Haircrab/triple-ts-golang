package game

import (
	"errors"
	"log"
	"strings"

	"triple-ts-golang/game"

	socketio "github.com/googollee/go-socket.io"
)

type (
	ErrorEventRes struct {
		Message string `json:"message"`
	}
	CreateRoomOkEventRes struct {
		RoomId   string `json:"roomId"`
		PlayerId int    `json:"playerId"`
	}
	JoinRoomOkEventRes struct {
		RoomId     string       `json:"roomId"`
		PlayerId   int          `json:"playerId"`
		ReadyState [4]bool      `json:"readyState"`
		RoomState  game.RoomCtx `json:"roomState"`
	}
	ReadyEventRes struct {
		PlayerId      int  `json:"playerId"`
		IsReady       bool `json:"isReady"`
		IsGameStarted bool `json:"isGameStarted"`
	}
	MoveEventRes struct {
		PlayerId int       `json:"playerId"`
		Move     game.Move `json:"move"`
	}
	WinnerEventRes struct {
		PlayerId int `json:"playerId"`
	}

	ConnCtx struct {
		RoomId   string `json:"roomId"`
		PlayerId int    `json:"playerId"`
	}
)

const (
	ns         = "/game"
	r          = "/socket.io" + ns
	roomIdKey  = "roomId"
	errorEvent = "error"
	maxConn    = 4

	// client to server events
	sendMsgEvent     = "sendMsg"
	toggleReadyEvent = "toggleReady"

	// server to client events
	receiveMsgEvent   = "receiveMsg"
	createRoomOkEvent = "createRoomOk"
	joinRoomOkEvent   = "joinRoomOk"

	playerToggleReadyEvent = "playerToggleReady"
)

// TODO, clean up roomStates periodically
var roomStates = map[string]*game.RoomCtx{}

func InitGameSocketNS(ser *socketio.Server) {
	ser.OnConnect(ns, func(conn socketio.Conn) error {
		log.Printf("%v onConnected, sid: %v \n", r, conn.ID())

		rawQuery := conn.URL().RawQuery
		querys := strings.Split(rawQuery, "&")

		rid, err := findRoomId(querys)
		if err != nil {
			return err
		}

		// roomId can be undefined -> new room
		if rid == "undefined" {
			//* --- create room ---
			log.Println("create room", conn.ID())
			if err = createRoom(&conn, &roomStates); err != nil {
				log.Println("create room error", err)
				return err
			}
		} else {
			//* --- join room ---
			if err = joinRoom(rid, ser, &conn, &roomStates); err != nil {
				return err
			} else {
				log.Println("join room", conn.ID())
			}
		}

		return nil
	})

	//* testing event
	ser.OnEvent(ns, sendMsgEvent, func(conn socketio.Conn, msg string) {
		log.Printf("%v onEvent - 'msg', sid: %v, rooms: %v \n", r, conn.ID(), conn.Rooms())

		for _, r := range conn.Rooms() {
			ser.BroadcastToRoom(ns, r, receiveMsgEvent, msg)
		}
	})

	ser.OnEvent(ns, toggleReadyEvent, func(conn socketio.Conn) {
		connCtx := conn.Context().(ConnCtx)
		rid, playerId := connCtx.RoomId, connCtx.PlayerId

		if ctx, ok := roomStates[rid]; ok {
			var player *game.Player
			if playerId == game.P1 {
				player = &(ctx).P1
			} else if playerId == game.P2 {
				player = &(ctx).P2
			} else if playerId == game.P3 {
				player = &(ctx).P3
			} else if playerId == game.P4 {
				player = &(ctx).P4
			}

			ctx.ToggleReady(player)

			for _, r := range conn.Rooms() {
				ser.BroadcastToRoom(ns, r, playerToggleReadyEvent, ReadyEventRes{
					PlayerId:      playerId,
					IsReady:       player.IsReady,
					IsGameStarted: ctx.IsGameStarted,
				})
			}
		} else {
			log.Println("room not found")
			err := errors.New("room not found")
			conn.Emit(errorEvent, ErrorEventRes{
				Message: err.Error(),
			})
			conn.Close()
		}

	})

	ser.OnDisconnect(ns, func(conn socketio.Conn, reason string) {
		log.Printf("%v onDisconnected, sid: %v, reason: %v \n", r, conn.ID(), reason)

		// remove player from roomStates
		if ctx, ok := conn.Context().(ConnCtx); ok {
			rid := ctx.RoomId
			roomState := roomStates[rid]

			id := roomState.OnPlayerDisconnect(conn.ID())

			ser.BroadcastToRoom(ns, rid, playerToggleReadyEvent, ReadyEventRes{
				PlayerId:      id,
				IsReady:       false,
				IsGameStarted: false,
			})
		}

	})
	// exception handlers
	ser.OnError(ns, func(conn socketio.Conn, e error) {
		log.Printf("%v onError, sid: %v, error: %v \n", r, conn.ID(), e)
	})

}
