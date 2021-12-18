package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Server running ...")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	log.Fatalln(server.Serve(listener))
}
