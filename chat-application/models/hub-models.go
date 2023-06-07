package models

type Hub struct {
	Rooms      map[string]*Room `json:"rooms"`
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}
type Room struct {
	RoomId   int64              `json:"room-id"`
	RoomName string             `json:"room-name"`
	Clients  map[string]*Client `json:"clients"`
}
