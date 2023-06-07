package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

type Service struct {
	connections map[*websocket.Conn]bool
}

func NewService() *Service {
	return &Service{
		connections: make(map[*websocket.Conn]bool),
	}
}
func (service *Service) WebSocketHandler(ws *websocket.Conn) {
	//establishing the connection
	fmt.Println("New incoming connection from the client: ", ws.RemoteAddr())
	service.connections[ws] = true
	service.ReadLoop(ws)
}
func (service *Service) ReadLoop(ws *websocket.Conn) {
	buffer := make([]byte, 1024)
	for {
		//reading the message from the browser
		noOfBytes, readErr := ws.Read(buffer)
		if readErr != nil {
			if readErr == io.EOF {
				break
			}
			fmt.Println("Error reading message: ", readErr)
			continue
		}
		message := buffer[:noOfBytes]
		fmt.Println("Message received is: ", string(message))
		service.Broadcast(message)
		//writeBytes, writeErr := ws.Write([]byte("Writing back to browser-->Thank you for your message"))
		//if writeErr != nil {
		//	fmt.Println("Error writing to server: ", writeErr)
		//	break
		//}
		//fmt.Println(writeBytes)
	}
}
func (service *Service) Broadcast(message []byte) {
	for ws := range service.connections {
		go func(ws *websocket.Conn) {
			if _, writeErr := ws.Write(message); writeErr != nil {
				fmt.Println("Error writing back to the browser", writeErr)
			}
		}(ws)
	}
}
func main() {
	service := NewService()
	http.Handle("/websocket", websocket.Handler(service.WebSocketHandler))
	err := http.ListenAndServe(":3300", nil)
	if err != nil {
		return
	}
}
