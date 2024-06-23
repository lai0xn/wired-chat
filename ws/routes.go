package ws

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
}

func RoomRoute(w http.ResponseWriter,r *http.Request){

   code := chi.URLParam(r,"code")
   usermame := r.URL.Query().Get("username")
   conn,err := upgrader.Upgrade(w,r,nil)
   defer conn.Close()
   if err != nil {
    http.Error(w,err.Error(),http.StatusBadRequest)
   }
   usr := User{
      Conn: conn,
      Username: usermame,
      Current_Code: code,
   }
   join<-usr
   for {
    _,msg,err := conn.ReadMessage()
    if err != nil {
      http.Error(w,err.Error(),http.StatusBadRequest)
    }
 
   broadcast<-Message{Code: code,Content: string(msg),Date: time.Now(),User: usr}
  }

  
}
