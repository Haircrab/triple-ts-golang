package game

import (
	"errors"
	"log"
	"strings"

	"triple-ts-golang/game"

	socketio "github.com/googollee/go-socket.io"
	"golang.org/x/tools/go/analysis/passes/printf"
)

type (
	ConnCtx struct {
		RoomId   string `json:"roomId"`
		PlayerId int    `json:"playerId"`
	}
)

// client to server
type (
	MakeMoveEventReq struct {
		PlayerId int       `json:"playerId"`
		Move     game.Move `json:"move"`
	}
)

// server to client
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
		PlayerId      int          `json:"playerId"`
		IsReady       bool         `json:"isReady"`
		IsGameStarted bool         `json:"isGameStarted"`
		RoomState     game.RoomCtx `json:"roomState"`
	}
	OtherPlayerMakeMoveEventRes struct {
		PlayerId  int          `json:"playerId"`
		Move      game.Move    `json:"move"`
		RoomState game.RoomCtx `json:"roomState"`
	}
	WinnerEventRes struct {
		PlayerId int `json:"playerId"`
	}
)

const (
	ns         = "/game"
	r          = "/socket.io" + ns
	roomIdKey  = "roomId"
	errorEvent = "error"
	maxConn    = 4

	// client to server events
	sendMsgEvent        = "sendMsg"
	toggleReadyEvent    = "toggleReady"
	playerMakeMoveEvent = "playerMakeMove"

	// server to client events
	receiveMsgEvent          = "receiveMsg"
	createRoomOkEvent        = "createRoomOk"
	joinRoomOkEvent          = "joinRoomOk"
	playerToggleReadyEvent   = "playerToggleReady"
	otherPlayerMakeMoveEvent = "otherPlayerMakeMove"
)

// TODO, clean up roomStates periodically
var roomStates = map[string]*game.RoomCtx{}

func InitGameSocketNS(ser *socketio.Server) {
	//* testing event
	ser.OnEvent(ns, sendMsgEvent, func(conn socketio.Conn, msg string) {
		log.Printf("%v onEvent - %v, sid: %v, rooms: %v \n", r, conn.ID(), conn.Rooms())

		for _, r := range conn.Rooms() {
			ser.BroadcastToRoom(ns, r, receiveMsgEvent, msg)
		}
	})

	// create or join room
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

	ser.OnEvent(ns, playerMakeMoveEvent, func(conn socketio.Conn, mv game.Move) {
		log.Printf("make move %v", mv)

		if err := mv.CheckMove(); err != nil {
			conn.Emit(errorEvent, ErrorEventRes{
				Message: err.Error(),
			})
		}

		connCtx := conn.Context().(ConnCtx)
		rid, playerId := connCtx.RoomId, connCtx.PlayerId

		if ctx, ok := roomStates[rid]; ok {
			player := ctx.FindPlayerByIdx(playerId)
			win, err := ctx.GameState.MakeMove(player, mv)
			if err != nil {
				conn.Emit(errorEvent, ErrorEventRes{
					Message: err.Error(),
				})
			}

			// TODO
			if win {

			}

			log.Printf("room State after move: %v", ctx.Copy())

			for _, r := range conn.Rooms() {
				ser.BroadcastToRoom(ns, r, otherPlayerMakeMoveEvent, OtherPlayerMakeMoveEventRes{
					PlayerId:  playerId,
					Move:      mv,
					RoomState: ctx.Copy(),
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

	ser.OnEvent(ns, toggleReadyEvent, func(conn socketio.Conn) {
		connCtx := conn.Context().(ConnCtx)
		rid, playerId := connCtx.RoomId, connCtx.PlayerId

		if ctx, ok := roomStates[rid]; ok {
			player := ctx.FindPlayerByIdx(playerId)
			ctx.ToggleReady(player)

			for _, r := range conn.Rooms() {
				ser.BroadcastToRoom(ns, r, playerToggleReadyEvent, ReadyEventRes{
					PlayerId:      playerId,
					IsReady:       player.IsReady,
					IsGameStarted: ctx.IsGameStarted,
					RoomState:     ctx.Copy(),
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
				RoomState:     (*roomState).Copy(),
			})
		}

	})
	// exception handlers
	ser.OnError(ns, func(conn socketio.Conn, e error) {
		log.Printf("%v onError, sid: %v, error: %v \n", r, conn.ID(), e)
	})

}
