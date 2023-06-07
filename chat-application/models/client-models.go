package models

import "github.com/gorilla/websocket"

type Client struct {
	ClientID int64  `json:"client-id"`
	RoomID   int64  `json:"room-id"`
	UserName string `json:"user-name"`
	Conn     *websocket.Conn
	Message  chan *Message
}
type Message struct {
	Content  string `json:"content"`
	RoomID   int64  `json:"room-id"`
	UserName string `json:"user-name"`
}
