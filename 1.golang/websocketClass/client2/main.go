package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

func main() {
	dl := websocket.Dialer{
		NetDial:           nil,
		NetDialContext:    nil,
		NetDialTLSContext: nil,
		Proxy:             nil,
		TLSClientConfig:   nil,
		HandshakeTimeout:  0,
		ReadBufferSize:    0,
		WriteBufferSize:   0,
		WriteBufferPool:   nil,
		Subprotocols:      nil,
		EnableCompression: false,
		Jar:               nil,
	}
	conn, _, err:=dl.Dial("ws://localhost:8888", nil)
	if err != nil {
		log.Println(err)
		return
	}
	//conn.WriteMessage(websocket.TextMessage, []byte("我来了"))
	go send(conn)
	for {
		m,p,e := conn.ReadMessage()
		if e!= nil {
			log.Println(e)
			break
		}
		fmt.Println(m,string(p))
	}
}


func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		l,_,_ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, l)
	}
}