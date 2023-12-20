package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	log.Println("New connection: ", ws.RemoteAddr())

	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("read error: ", err)
			continue
		}

		msg := buf[:n]
	}
}

func main() {
	fmt.Println("hello world")
}
