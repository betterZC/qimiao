package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler (res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res.Write([]byte("我收到了，给你返回GET"))
		break
	case "POST":
		//res.Write([]byte("我收到了，给你返回POST"))
		b,_ :=io.ReadAll(req.Body) 
		fmt.Println(req.Header["Test"])
		//header := res.Header()
		//header["test"] = []string{"test1,test2"}
		//res.WriteHeader(http.StatusBadRequest)
		res.Write(b)
		break
	
	}
	
	//res.Header()
}

func main() {
	mux := http.NewServeMux()
	//http.HandleFunc("/test", handler)
	//http.ListenAndServe(":8080",nil)
	mux.HandleFunc("/test", handler)
	http.ListenAndServe(":8080",mux)
}
