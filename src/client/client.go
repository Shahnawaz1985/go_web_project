package main

import (
	"fmt"
	"net/http"
	"requesthandlers"
	"typedef"
)

func main(){
	fmt.Println("client for running web application!")
	http.HandleFunc("/guestbook", requesthandlers.ViewHandler)
	http.HandleFunc("/guestbook/new", requesthandlers.NewHandler)
	http.HandleFunc("/guestbook/create", requesthandlers.CreateHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	typedef.Check(err)
	//fmt.Println(typedef.GetStrings("signatures.txt"))
}
