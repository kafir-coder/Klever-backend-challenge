package main

import (
	mongodb "Token_service/pkg/mongo"
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongo_url := os.Getenv("MONGO_URL")
	port := os.Getenv("PORT")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}

	mongodb.DB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_url))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("🔥 Server listening on port " + port)

	server := grpc.NewServer()

	log.Fatalln(server.Serve(listener))
}
