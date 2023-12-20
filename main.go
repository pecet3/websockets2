package main

import (
	"fmt"

	"nhooyr.io/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func main() {
	fmt.Println("hello world")
}
