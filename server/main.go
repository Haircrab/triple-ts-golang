package main

import (
	"log"
	"net/http"

	app_socket "triple-ts-golang/app_socket"
	game_socket "triple-ts-golang/app_socket/game"

	"github.com/gin-gonic/gin"
)

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func main() {
	r := gin.New()

	ser := app_socket.InitSocketioSer()
	defer ser.Close()

	app_socket.InitSocketNS(ser)
	game_socket.InitGameSocketNS(ser)

	go func() {
		if err := ser.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	r.Use(GinMiddleware("http://localhost:3000"))
	r.GET("/socket.io/*any", gin.WrapH(ser))
	r.POST("/socket.io/*any", gin.WrapH(ser))

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
