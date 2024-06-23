package ws

import (
	"time"

	"github.com/gorilla/websocket"
)



type User struct {
  Conn *websocket.Conn
  Username string
  Current_Code string
}

type Room struct {
  Code string
  Users []User
  Messages []Message
}

type Message struct {
  Content string
  Code string
  User User
  Date time.Time
}


var join = make(chan User)

var broadcast = make(chan Message)


var rooms = make(map[string]Room)

func joinHandler(){
  for {
    output := <-join
    room, ok := rooms[output.Current_Code]
    if ok {
      room.Users = append(room.Users, output)
    }else {
      new_room := Room{
        Code: output.Current_Code,
        Users: []User{output},
        Messages: []Message{},

      }
      rooms[output.Current_Code] = new_room
    }
  } 
}

func broadcastHandler(){
  for {
    msg := <- broadcast
    room := rooms[msg.Code]
    room.Messages = append(room.Messages, msg)
    rooms[msg.Code] = room
    for _,user := range room.Users {
      user.Conn.WriteJSON(msg)
    }
  }
}
