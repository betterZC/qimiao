package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Req struct {
	NumOne int
	NumTwo int
}

type Res struct {
	Num int
}

func main() {
	req:=Req{NumOne: 1,NumTwo: 2}
	var res Res
	client,err := rpc.DialHTTP("tcp", "localhost:8888")
	if err!= nil {
		log.Fatal("dialing:",err)
		
	}
	//client.Call("Server.Add",req,&res)
	ca:=client.Go("Server.Add",req,&res,nil)
	//fmt.Println("此期间可做很多事情 当server sleep")
	//<-ca.Done
	for {
		select {
		case <-ca.Done:
			fmt.Println(res)
			return
		default:
			time.Sleep(1*time.Second)
			println("等待中////")
		}
	}
	//fmt.Println(res)
	 
	
}
