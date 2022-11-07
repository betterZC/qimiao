package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr,_ := net.ResolveTCPAddr("tcp", ":8088")
	conn,_ := net.DialTCP("tcp", nil, tcpAddr)
	reader := bufio.NewReader(os.Stdin)
	for {
		bytes,_,_ := reader.ReadLine()
		conn.Write(bytes)
		rb := make([]byte,1024)
		rn,_:=conn.Read(rb)
		fmt.Println(string(rb[0:rn]))
	}
	
}
