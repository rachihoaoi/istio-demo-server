package main

import (
	"fmt"
	"log"
	"net/http"
)

func main () {
	ch := make(chan int, 1)
	//go StartGRPC()
	go StartHttp()
	<- ch
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got req from : %v \n", r.Host)
	w.Write([]byte("bye bye ,this is v1 httpServer"))
}

func StartHttp() {
	log.Println("Starting http server ...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/test/restful", sayBye)
	log.Println("Http Server Started")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
