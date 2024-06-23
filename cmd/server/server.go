package main

import "github.com/lai0xn/websocket.go/ws"

func main(){
  s := ws.NewServer()
  
  s.Run()
}
