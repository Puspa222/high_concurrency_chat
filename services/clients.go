package services

import "github.com/gorilla/websocket"

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

// ReadPump pumps messages from the websocket connection to the hub.
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		c.Hub.Broadcast <- message
	}
}

// WritePump pumps messages from the hub to the websocket connection.
func (c *Client) WritePump() {
	for message := range c.Send {
		c.Conn.WriteMessage(websocket.TextMessage, message)
	}
}
