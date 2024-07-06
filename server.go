package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
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

type Message struct {

}

func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
  // get query parameter
  conn, _, _, err := ws.UpgradeHTTP(r, w)
  roomId := r.URL.Query().Get("room")
  if roomId == "" {
    roomId = "1234"
    // create a new room
  }
  room := s.rooms[roomId]
  room.clients = append(room.clients, Client{
  	conn:     &conn,
  	username: "something",
  	cursor:   CursorPosition{
  		x: 0,
  		y: 0,
  	},
  })
  log.Printf("%v connected", conn.RemoteAddr().String())
  for {
    msg, _, err := wsutil.ReadClientData(conn)
    log.Printf("%v said: %v", conn.RemoteAddr(), string(msg))
    if err != nil {
      log.Println(err)
      break
    }
  }
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
