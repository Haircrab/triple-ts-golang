package game

import (
	"errors"
	"log"
	"strings"

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
		RoomId     string  `json:"roomId"`
		PlayerId   int     `json:"playerId"`
		ReadyState [4]bool `json:"readyState"`
		RoomState  RoomCtx `json:"roomState"`
	}
	ReadyEventRes struct {
		PlayerId      int  `json:"playerId"`
		IsReady       bool `json:"isReady"`
		IsGameStarted bool `json:"isGameStarted"`
	}
	MoveEventRes struct {
		PlayerId int  `json:"playerId"`
		Move     move `json:"move"`
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
var roomStates = map[string]*RoomCtx{}

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
			if err = createRoom(rid, &conn, &roomStates); err != nil {
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

	ser.OnEvent(ns, sendMsgEvent, func(conn socketio.Conn, msg string) {
		log.Printf("%v onEvent - 'msg', sid: %v, rooms: %v \n", r, conn.ID(), conn.Rooms())

		for _, r := range conn.Rooms() {
			ser.BroadcastToRoom(ns, r, receiveMsgEvent, msg)
		}
	})

	ser.OnEvent(ns, toggleReadyEvent, func(conn socketio.Conn) {
		connCtx := conn.Context().(ConnCtx)
		rid, playerId := connCtx.RoomId, connCtx.PlayerId

		if rs, ok := roomStates[rid]; ok {
			var ptr *player
			if playerId == p1 {
				ptr = &(rs).P1
			} else if playerId == p2 {
				ptr = &(rs).P2
			} else if playerId == p3 {
				ptr = &(rs).P3
			} else if playerId == p4 {
				ptr = &(rs).P4
			}

			rs.mu.Lock()
			ptr.IsReady = !ptr.IsReady
			if rs.P1.IsReady && rs.P2.IsReady && rs.P3.IsReady && rs.P4.IsReady {
				rs.IsGameStarted = true
			} else {
				rs.IsGameStarted = false
			}
			rs.mu.Unlock()

			for _, r := range conn.Rooms() {
				ser.BroadcastToRoom(ns, r, playerToggleReadyEvent, ReadyEventRes{
					PlayerId:      playerId,
					IsReady:       ptr.IsReady,
					IsGameStarted: rs.IsGameStarted,
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
			for i, v := range roomStates[rid].ConnPlayerIdx {
				if v == conn.ID() {
					rs := roomStates[rid]
					rs.ConnPlayerIdx[i] = ""
					if i == 0 {
						rs.P1.IsReady = false
					} else if i == 1 {
						rs.P2.IsReady = false
					} else if i == 2 {
						rs.P3.IsReady = false
					} else if i == 3 {
						rs.P4.IsReady = false
					}
					rs.IsGameStarted = false

					ser.BroadcastToRoom(ns, rid, playerToggleReadyEvent, ReadyEventRes{
						PlayerId:      i,
						IsReady:       false,
						IsGameStarted: false,
					})
				}
			}
		}

	})
	// exception handlers
	ser.OnError(ns, func(conn socketio.Conn, e error) {
		log.Printf("%v onError, sid: %v, error: %v \n", r, conn.ID(), e)
	})

}

func findRoomId(querys []string) (string, error) {
	for _, q := range querys {
		tmp := strings.Split(q, "=")
		k, v := tmp[0], tmp[1]

		if k == roomIdKey {
			return v, nil
		}

	}
	return "", errors.New("roomId key not found")
}

func findRoomExist(rooms []string, rid string) error {
	for _, v := range rooms {
		if v == rid {
			return nil
		}
	}

	return errors.New("roomId not found")
}

func createRoom(rid string, conn *socketio.Conn, rs *map[string]*RoomCtx) error {
	defer func() {
		for k, v := range *rs {
			log.Printf("roomStates[%v]: %v \n", k, v)
		}
	}()
	if ctx, err := InitRoomCtx(); err != nil {
		return err
	} else {
		connID := (*conn).ID()
		(*rs)[connID] = ctx

		ctx.mu.Lock()
		ctx.ConnPlayerIdx[0] = connID
		ctx.mu.Unlock()

		(*conn).Emit(createRoomOkEvent, CreateRoomOkEventRes{
			RoomId:   connID,
			PlayerId: 0,
		})
		(*conn).SetContext(ConnCtx{
			RoomId:   connID,
			PlayerId: 0,
		})
		return nil
	}
}

func joinRoom(rid string, ser *socketio.Server, conn *socketio.Conn, rs *map[string]*RoomCtx) error {
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

		ctx.mu.Lock()
		ctx.ConnPlayerIdx[playerId] = connID
		ctx.mu.Unlock()

		(*conn).SetContext(ConnCtx{
			RoomId:   rid,
			PlayerId: playerId,
		})

		ctx.mu.Lock()
		readyState := [4]bool{ctx.P1.IsReady, ctx.P2.IsReady, ctx.P3.IsReady, ctx.P4.IsReady}
		ctx.mu.Unlock()

		//* Join room
		(*conn).Emit(joinRoomOkEvent, JoinRoomOkEventRes{
			RoomId:     rid,
			PlayerId:   playerId,
			ReadyState: readyState,
			RoomState:  *ctx,
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
