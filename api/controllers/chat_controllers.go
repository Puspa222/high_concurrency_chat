package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/puspa222/high_concurrency_chat/services"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type ChatController struct {
	Hub *services.Hub
}

func NewChatController(hub *services.Hub) ChatController {
	return ChatController{Hub: hub}
}

func (cc ChatController) HandleWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := &services.Client{Hub: cc.Hub, Conn: conn, Send: make(chan []byte, 256)}
	cc.Hub.Register <- client

	// Start the pumps in new goroutines
	go client.WritePump()
	go client.ReadPump()
}
