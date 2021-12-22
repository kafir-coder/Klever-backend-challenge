package main

import (
	"Token_service/internal/token"
	mongodb "Token_service/pkg/mongodb"
	"context"
	"log"
	"net"
	"os"

	pb "Token_service/pkg/proto/token"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	defer mongodb.DB.Disconnect(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("🔥 Server listening on port " + port)

	server := grpc.NewServer()

	pb.RegisterTokenUpvoteServiceServer(server, &token.TokenService{})
	reflection.Register(server)
	log.Fatalln(server.Serve(listener))
}
