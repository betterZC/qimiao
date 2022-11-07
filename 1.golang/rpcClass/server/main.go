package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
	
}

type Req struct {
	NumOne int
	NumTwo int
}

type Res struct {
	Num int
}

func (s *Server) Add(req Req, res *Res)error {
	time.Sleep(5*time.Second)
	res.Num = req.NumTwo + req.NumOne
	return nil
	
}

func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	l,err:=net.Listen("tcp",":8888")
	if err!= nil {
		fmt.Println("出错")
	}
	
	http.Serve(l,nil)
}
