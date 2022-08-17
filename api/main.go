package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rAndrade360/go-link-shortener/api/handler"
	"github.com/rAndrade360/go-link-shortener/api/handler/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.HandleUrl)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Println("Erro to create grpc listener ", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterShortUrlServer(s, &handler.Server{})
	go func() {
		reflection.Register(s)
		err = s.Serve(listener)
		if err != nil {
			log.Println("Erro to create grpc server ", err.Error())
		}
	}()

	log.Fatal(srv.ListenAndServe())
}
