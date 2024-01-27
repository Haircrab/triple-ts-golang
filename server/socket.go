package main

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

const (
	ns         = "/"
	r          = "/socket.io" + ns
	roomIDKey  = "roomID"
	errorEvent = "error"
)

func InitSocketNS(server *socketio.Server) {
	server.OnConnect("/", func(conn socketio.Conn) error {
		log.Printf("%v onConnected, sid: %v \n", r, conn.ID())
		return nil
	})

	server.OnDisconnect(ns, func(conn socketio.Conn, reason string) {
		log.Printf("%v onDisconnect, sid: %v, reason: %v \n", r, conn.ID(), reason)
	})
	server.OnError(ns, func(conn socketio.Conn, e error) {
		log.Printf("%v onError, error: %v \n", r, e)
	})
}
