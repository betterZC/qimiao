package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var conns []*websocket.Conn //为了让不同客户端间可以通讯

func handler(w http.ResponseWriter, r *http.Request) {
	conn,err:=UP.Upgrade(w,r, nil)
	if err!= nil{
		log.Println(err)
		return
	}
	conns = append(conns, conn)
	for {
		m, p , e := conn.ReadMessage()
		if e!= nil { 
			break
		}
		for i := range conns {
			conns[i].WriteMessage(websocket.TextMessage,[]byte("你说的是："+string(p)+"吗？"))
		}
		//conn.WriteMessage(websocket.TextMessage,[]byte("你说的是："+string(p)+"吗？"))
		fmt.Println(m,string(p),e)
	}
	defer conn.Close()
	log.Println("服务关闭")
	
}

var UP =websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8888",nil)
}