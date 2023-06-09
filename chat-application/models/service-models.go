package models

import "golang.org/x/net/websocket"

type Service struct {
	connection map[websocket.Conn]bool
}
