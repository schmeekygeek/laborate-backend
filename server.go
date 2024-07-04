package main

import (
	"net"
  "log"
	"net/http"
  "github.com/gobwas/ws"
)

type Server struct {
  rooms map[string]Room
}

type Room struct {
  clients[] Client
  textState []byte
}

type Client struct {
  conn      *net.Conn
  username  string
  cursor    CursorPosition
}

type CursorPosition struct {
  x, y int
}

func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
  // get query parameter
  conn, _, _, err := ws.UpgradeHTTP(r, w)
  log.Printf("%v connected", conn.RemoteAddr().String())
  if err != nil {
    log.Println(err.Error())
  }
}

func Init() *Server {
  s := Server{
  	rooms: map[string]Room{},
  }
  s.rooms["1234"] = Room{
  	clients:   []Client{},
  	textState: []byte{},
  }
  return &s
}
