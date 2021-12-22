package token

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"Token_service/pkg/mongodb"
	pb "Token_service/pkg/proto/token"
)

func eraseData(coll *mongo.Collection) {
	coll.DeleteMany(context.TODO(), bson.D{{}})
}

func getNumberOfDocuments(coll *mongo.Collection) int64 {
	n_doc, _ := coll.EstimatedDocumentCount(context.TODO())
	return n_doc
}

func createOneToken(coll *mongo.Collection) primitive.ObjectID {

	token := bson.D{
		{Key: "name", Value: "some_name"},
		{Key: "price", Value: "hirata"},
	}

	id, _ := coll.InsertOne(context.TODO(), token)

	return id.InsertedID.(primitive.ObjectID)
}

func createOnevote(coll *mongo.Collection) primitive.ObjectID {
	vote := bson.D{
		{Key: "to", Value: "invalid_object_id"},
	}

	id, _ := coll.InsertOne(context.TODO(), vote)

	return id.InsertedID.(primitive.ObjectID)
}
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

	coll := mongodb.DB.Database("example").Collection("tokens")
	defer eraseData(coll)
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
	coll := mongodb.DB.Database("example").Collection("tokens")

	t.Run("it should return error if Id is empty", func(t *testing.T) {
		request := &pb.GetTokenRequest{Id: ""}
		_, err := client.GetToken(ctx, request)

		want := status.Error(codes.InvalidArgument, "Id musn't be empty")
		if err.Error() != want.Error() {

			t.Errorf("Should return a invalidArgument error when id is empty")
		}
	})
	id := createOneToken(coll)
	t.Run("it should return a Token Object if all suceeds", func(t *testing.T) {

		request := &pb.GetTokenRequest{Id: id.Hex()}

		response, err := client.GetToken(ctx, request)

		if err != nil {

			t.Errorf("Should return a nil Error object")
		}
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
	coll := mongodb.DB.Database("example").Collection("tokens")
	defer eraseData(coll)
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

func TestUpdateToken(t *testing.T) {
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
	coll := mongodb.DB.Database("example").Collection("tokens")

	client := pb.NewTokenUpvoteServiceClient(conn)
	defer eraseData(coll)
	t.Run("it should return noData equal true if id is invalid", func(t *testing.T) {
		createOneToken(coll)
		request := &pb.UpdateTokenRequest{Id: "invalid_object_id"}
		response, _ := client.UpdateToken(ctx, request)
		if response.NoData != true {
			t.Errorf("response.NoData must be true")
		}

	})

	t.Run("it should return noData equal true if the database is empty", func(t *testing.T) {
		eraseData(coll)
		request := &pb.UpdateTokenRequest{Id: "invalid_object_id"}
		response, _ := client.UpdateToken(ctx, request)
		if response.NoData != true {
			t.Errorf("response.NoData must be true")
		}
	})
}

func TestDeleteToken(t *testing.T) {
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
	coll := mongodb.DB.Database("example").Collection("tokens")

	client := pb.NewTokenUpvoteServiceClient(conn)
	defer eraseData(coll)
	t.Run("it should return noData equal true if id is invalid", func(t *testing.T) {
		createOneToken(coll)
		request := &pb.DeleteTokenRequest{Id: "invalid_object_id"}
		response, _ := client.DeleteToken(ctx, request)
		if response.NoData != true {
			t.Errorf("response.NoData must be true")
		}

	})
	t.Run("it should return noData equal true if the database is empty", func(t *testing.T) {
		eraseData(coll)
		request := &pb.DeleteTokenRequest{Id: "invalid_object_id"}
		response, _ := client.DeleteToken(ctx, request)
		if response.NoData != true {
			t.Errorf("response.NoData must be true")
		}
	})
}

func TestUpvoteToToken(t *testing.T) {
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
	coll := mongodb.DB.Database("example").Collection("vote")

	client := pb.NewTokenUpvoteServiceClient(conn)
	defer eraseData(coll)

	t.Run("it should have one more document", func(t *testing.T) {
		before := getNumberOfDocuments(coll)
		request := &pb.UpvoteVoteRequest{Id: "invalid_object_id"}
		client.UpvoteToToken(ctx, request)
		after := getNumberOfDocuments(coll)
		if before+1 != after {
			t.Errorf("response.NoData must be true")
		}
	})

	t.Run("it should have return a OK message", func(t *testing.T) {

		request := &pb.UpvoteVoteRequest{Id: "invalid_object_id"}
		response, _ := client.UpvoteToToken(ctx, request)
		if response.Ok != true {
			t.Errorf("response.NoData must be true")
		}
	})
}

func TestDownvoteToToken(t *testing.T) {
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
	coll := mongodb.DB.Database("example").Collection("vote")

	client := pb.NewTokenUpvoteServiceClient(conn)

	defer eraseData(coll)

	t.Run("it should have the same number of document minus 1", func(t *testing.T) {
		createOnevote(coll)
		before := getNumberOfDocuments(coll)
		request := &pb.DownVoteRequest{Id: "invalid_object_id"}
		client.DownvoteToToken(ctx, request)
		after := getNumberOfDocuments(coll)
		if before-1 != after {
			t.Errorf("response.NoData must be true")
		}
	})

	t.Run("it should return a OK message with ok ", func(t *testing.T) {
		createOnevote(coll)
		request := &pb.DownVoteRequest{Id: "invalid_object_id"}
		response, _ := client.DownvoteToToken(ctx, request)
		if response.Ok != true {
			t.Errorf("response.NoData must be true")
		}
	})

}
