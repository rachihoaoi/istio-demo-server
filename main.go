package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"

	"github.com/rachihoaoi/istio-demo-server/pb"
)

func main () {
	ch := make(chan int, 1)
	go StartGRPC()
	go StartHttp()
	<- ch
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got req from : %v \n", r.Host)
	w.Write([]byte("bye bye ,this is v2 httpServer"))
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


type server struct {}

func (s *server) SayHello(ctx context.Context, in *pbs.HelloReq) (*pbs.HelloResp, error) {
	return &pbs.HelloResp{Result: "bye bye ,this is v2 grpcServer"}, nil
}

func StartGRPC() {
	log.Println("Starting grpc server ...")
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbs.RegisterHelloServer(s, &server{})
	log.Println("Grpc Server Started")
	_ = s.Serve(lis)
}