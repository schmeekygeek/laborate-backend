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

const (
  TEXTUPDATE = "text/update"
  USERUPDATE = "user/update"
  USERDSCNT = "user/disconnect"
  USERCNT = "user/connect"
)

type Message struct {
  messageType   string
  payload       []byte
}

func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
  // get query parameter
  conn, _, _, err := ws.UpgradeHTTP(r, w)
  if err != nil {
    log.Println(err.Error())
  }
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
  s.rooms[roomId] = room
  log.Printf("%v connected", conn.RemoteAddr().String())
  for {
    msg, _, err := wsutil.ReadClientData(conn)
    parseServerMessage(msg)
    if string(msg) == "hi" {
      log.Println(room)
    }
    log.Printf("Received message %v from %v", conn.RemoteAddr(), string(msg))
    if err != nil {
      log.Println(err)
      return
    }
  }
}

func parseServerMessage(msg []byte) {
  switch string(msg) {
    case TEXTUPDATE:
      log.Println("hi")
  }
}

func (s *Server) broadcastMessage(roomId string, msg []byte) {
  room, ok := s.rooms[roomId]
  if ok {
    return
  }
  for _, v := range room.clients {
    wsutil.WriteServerBinary(*v.conn, msg)
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
