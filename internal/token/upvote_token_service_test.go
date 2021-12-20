package token

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"Token_service/pkg/mongodb"
	pb "Token_service/pkg/proto/token"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterTokenUpvoteServiceServer(server, &TokenService{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestAddToken(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	mongo_url := os.Getenv("MONGO_URL")

	mongodb.DB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_url))
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTokenUpvoteServiceClient(conn)

	t.Run("it should returns a error if name is empty", func(t *testing.T) {
		request := &pb.AddTokenRequest{Name: "", Price: 33}

		_, err := client.AddToken(ctx, request)

		want := status.Error(codes.InvalidArgument, "Name musn't be empty")
		if err.Error() != want.Error() {

			t.Errorf("Should return a invalidArgument error when name is empty")
		}
	})

	t.Run("it should returns a error if price is negative", func(t *testing.T) {
		request := &pb.AddTokenRequest{Name: "some", Price: -2}

		_, err := client.AddToken(ctx, request)

		want := status.Error(codes.InvalidArgument, "Price must be positive")
		if err.Error() != want.Error() {

			t.Errorf("Should return a invalidArgument error when Price is negative")
		}
	})

	t.Run("it should return a nil error if Name is a no empty string and price is a non-negative number", func(t *testing.T) {
		request := &pb.AddTokenRequest{Name: "some", Price: 2}

		_, err := client.AddToken(ctx, request)

		if err != nil {

			t.Errorf("Should return a nil Error object")
		}
	})

	t.Run("it should return a pb.Token object if if Name is a no empty string and price is a non-negative number", func(t *testing.T) {
		request := &pb.AddTokenRequest{Name: "some", Price: 2}

		response, err := client.AddToken(ctx, request)

		if err != nil {

			t.Errorf("Should return a nil Error object")
		}

		isTrue := response.Name != request.Name && response.Price != request.Price && fmt.Sprintf("%T", response.Id) != "string" && response.Id == ""
		if isTrue {
			t.Errorf("Should return a valid token object")
		}
	})
}

func TestGetToken(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	mongo_url := os.Getenv("MONGO_URL")

	mongodb.DB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_url))
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTokenUpvoteServiceClient(conn)
	t.Run("it should return error if Id is empty", func(t *testing.T) {
		request := &pb.GetTokenRequest{Id: ""}
		_, err := client.GetToken(ctx, request)

		want := status.Error(codes.InvalidArgument, "Id musn't be empty")
		if err.Error() != want.Error() {

			t.Errorf("Should return a invalidArgument error when id is empty")
		}
	})
	t.Run("it should return a Token Object if all suceeds", func(t *testing.T) {
		request := &pb.GetTokenRequest{Id: "61c0a02936477d49339ea394"}

		response, err := client.GetToken(ctx, request)

		if err != nil {

			t.Errorf("Should return a nil Error object")
		}

		fmt.Println(response)
		have_the_proper_types :=
			fmt.Sprintf("%T", response.Name) != "string" && fmt.Sprintf("%T", response.Price) != "float64" && fmt.Sprintf("%T", response.Id) != "string"

		arent_empty :=
			response.Id == "" && response.Name == "" && response.Price < 0

		if have_the_proper_types && arent_empty {
			t.Errorf("Should return a valid token object")
		}
	})
}

func TestListToken(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	mongo_url := os.Getenv("MONGO_URL")

	mongodb.DB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_url))
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTokenUpvoteServiceClient(conn)

	t.Run("it should returns error if Limit and Page doesnt exists", func(t *testing.T) {
		request := &pb.ListTokenRequest{}
		response, _ := client.ListToken(ctx, request)

		_, err = response.Recv()

		want := status.Error(codes.InvalidArgument, "You must provide paginate data")
		if err.Error() != want.Error() {

			t.Errorf("Should return a invalidArgument error when Name and Page doenst exists")
		}
	})
	t.Run("returns 3 times if Limit equals 3", func(t *testing.T) {
		request := &pb.ListTokenRequest{Limit: 3, Page: 1}
		response, _ := client.ListToken(ctx, request)

		//result, _ := response.Recv()

		response.Recv()
		response.Recv()
		response.Recv()
		_, err := response.Recv()

		want := "EOF"
		if want != err.Error() {
			t.Errorf("returns 3 times if Limit equals 3")
		}
	})
}
