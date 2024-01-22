package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	defer fmt.Println("Stopping server")

	ginRouter := gin.Default()
	ginRouter.Use(gin.Recovery())
	err := ginRouter.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	socketRoute := ginRouter.Group("/socket")
	socketRoute.GET("", SocketHandler)

	if err := ginRouter.Run(":8080"); err != nil {
		panic(err)
	}
}

func SocketHandler(c *gin.Context) {
	upGrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	// Close socket impl
	defer func() {
		closeSocketErr := ws.Close()
		if closeSocketErr != nil {
			panic(err)
		}
	}()

	// Socket message receiving
	for {
		// receiving
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Message Type: %d, Message: %s\n", msgType, string(msg))

		// sending
		err = ws.WriteJSON(struct {
			Reply string `json:"reply"`
		}{
			Reply: "Echo...",
		})

		if err != nil {
			panic(err)
		}
	}
}
