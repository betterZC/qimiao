package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr,_ := net.ResolveTCPAddr("tcp", ":8088")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, err := listener.AcceptTCP()
		if err!= nil {
			fmt.Println(err)
			return
		}
		go handleConnection(conn)	
	}
	
}


func handleConnection(conn *net.TCPConn) {
	for{
		buf := make([]byte, 1024)
		n,err :=conn.Read(buf)
		if err !=nil{
			fmt.Println(err)
			break
		}
		fmt.Println(conn.RemoteAddr().String() +string(buf[:n]))
		str := "收到了"+string(buf[0:n])
		conn.Write([]byte(str))
	}
	
}