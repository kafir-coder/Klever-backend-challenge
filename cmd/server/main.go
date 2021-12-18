package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("🔥 Server listening on port " + port)

	server := grpc.NewServer()

	log.Fatalln(server.Serve(listener))
}
