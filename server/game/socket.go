package game

import (
	"errors"
	"log"
	"strings"

	socketio "github.com/googollee/go-socket.io"
)

type SocketRes struct {
	Message string `json:"message"`
}

const (
	ns          = "/game"
	r           = "/socket.io" + ns
	roomIDKey   = "roomID"
	errorEvent  = "error"
	roomIdEvent = "roomID"
)

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
			conn.Emit(roomIdEvent, conn.ID())
			return nil
		}

		// roomId not found
		err = findRoomExist(ser.Rooms(ns), rid)
		if err != nil {
			conn.Emit(errorEvent, SocketRes{
				Message: err.Error(),
			})
			conn.Close()
			return err
		}

		conn.LeaveAll()
		conn.Join(rid)
		conn.Emit(roomIdEvent, rid)

		return nil
	})

	ser.OnEvent(ns, "msg", func(conn socketio.Conn, msg string) string {
		log.Printf("%v onEvent - 'msg', sid: %v, rooms: %v \n", r, conn.ID(), conn.Rooms())

		for _, v := range conn.Rooms() {
			ser.BroadcastToRoom(ns, v, "msg", msg)
		}

		return ""
	})

	// exception handlers
	ser.OnDisconnect(ns, func(conn socketio.Conn, reason string) {
		log.Printf("%v onDisconnected, sid: %v, reason: %v \n", r, conn.ID(), reason)
	})
	ser.OnError(ns, func(conn socketio.Conn, e error) {
		log.Printf("%v onError, sid: %v, error: %v \n", r, conn.ID(), e)
	})
}

func findRoomId(querys []string) (string, error) {
	for _, q := range querys {
		tmp := strings.Split(q, "=")
		k, v := tmp[0], tmp[1]

		if k == roomIDKey {
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
