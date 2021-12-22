package token

import (
	"Token_service/pkg/mongodb"
	pb "Token_service/pkg/proto/token"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TokenService struct {
	pb.UnimplementedTokenUpvoteServiceServer
}

func (*TokenService) AddToken(ctx context.Context, req *pb.AddTokenRequest) (*pb.Token, error) {

	typechecks := []bool{
		fmt.Sprintf("%T", req.GetName()) == "string",
		fmt.Sprintf("%T", req.GetPrice()) == "float64",
	}

	for _, v := range typechecks {
		if !v {
			return nil, status.Errorf(codes.InvalidArgument, "Data types don't match")
		}
	}
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "Name musn't be empty")
	}

	if req.GetPrice() < 0 {
		return nil, status.Error(codes.InvalidArgument, "Price must be positive")
	}

	token := bson.D{
		{Key: "name", Value: req.GetName()},
		{Key: "price", Value: req.GetPrice()},
	}

	inserted_document, _ := mongodb.InsertOne(token, "tokens")
	doc, _ := bson.Marshal(inserted_document)

	var unmarshalled_document mongodb.Token
	bson.Unmarshal(doc, &unmarshalled_document)

	result := &pb.Token{
		Id:    unmarshalled_document.Id,
		Name:  unmarshalled_document.Name,
		Price: unmarshalled_document.Price,
	}
	return result, nil
}

func (*TokenService) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.Token, error) {
	typechecks := []bool{
		fmt.Sprintf("%T", req.GetId()) == "string",
	}
	for _, v := range typechecks {
		if !v {
			return nil, status.Errorf(codes.InvalidArgument, "Data types don't match")
		}
	}

	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "Id musn't be empty")
	}

	id := req.GetId()
	var token bson.D = mongodb.GetDocumentById(id, "tokens")

	doc, _ := bson.Marshal(token)

	var unmarshalled_document mongodb.Token
	bson.Unmarshal(doc, &unmarshalled_document)

	result := &pb.Token{
		Id:    unmarshalled_document.Id,
		Name:  unmarshalled_document.Name,
		Price: unmarshalled_document.Price,
	}
	return result, nil
}

func (*TokenService) ListToken(in *pb.ListTokenRequest, stream pb.TokenUpvoteService_ListTokenServer) error {

	if in.Limit == 0 && in.Page == 0 {
		return status.Error(codes.InvalidArgument, "You must provide paginate data")
	}
	token_list := mongodb.ListDocuments(in.Limit, in.Page, "tokens")
	for _, token := range token_list {

		unmarshalled := mongodb.Unmarshal(token, "token")

		token := unmarshalled.(mongodb.Token)
		if err := stream.Send(&pb.Token{
			Id:    token.Id,
			Name:  token.Name,
			Price: token.Price,
		}); err != nil {
			log.Printf("send error %v", err)
		}
	}
	return nil
}

func (*TokenService) UpdateToken(ctx context.Context, req *pb.UpdateTokenRequest) (*pb.Ok, error) {

	var update bson.D = makeUpdate(req)

	result, err := mongodb.UpdateById(req.GetId(), "tokens", update)
	if err != nil {
		return nil, err
	}
	ok := result.ModifiedCount > 0
	return &pb.Ok{
		Ok:     ok,
		NoData: !ok,
	}, nil
}

func makeUpdate(req *pb.UpdateTokenRequest) bson.D {
	var update bson.D
	if req.GetName() != "NULL" {
		update = bson.D{
			{Key: "name", Value: req.GetName()},
		}
	}
	if req.GetPrice() != -1E4 {
		update = append(update, bson.E{Key: "price", Value: req.GetPrice()})
	}
	return update
}

func (*TokenService) DeleteToken(ctx context.Context, req *pb.DeleteTokenRequest) (*pb.Ok, error) {

	result, err := mongodb.DeleteById(req.GetId(), "tokens")
	if err != nil {
		return nil, err
	}
	ok := result.DeletedCount > 0
	return &pb.Ok{
		Ok:     ok,
		NoData: !ok,
	}, nil
}

func (*TokenService) UpvoteToToken(ctx context.Context, req *pb.UpvoteVoteRequest) (*pb.Ok, error) {

	typechecks := []bool{
		fmt.Sprintf("%T", req.GetId()) == "string",
	}
	for _, v := range typechecks {
		if !v {
			return nil, status.Errorf(codes.InvalidArgument, "Data types don't match")
		}
	}

	vote := bson.D{
		{Key: "to", Value: req.GetId()},
	}

	inserted_document, _ := mongodb.InsertOne(vote, "vote")

	ok := len(inserted_document.Map()) == 2

	return &pb.Ok{Ok: ok}, nil

}

func (*TokenService) DownvoteToToken(ctx context.Context, req *pb.DownVoteRequest) (*pb.Ok, error) {
	result, err := mongodb.DeleteByX("to", req.GetId(), "vote")
	if err != nil {
		return nil, err
	}
	ok := result.DeletedCount > 0
	return &pb.Ok{
		Ok:     ok,
		NoData: !ok,
	}, nil
}

func (*TokenService) WatchTokenVotes(in *pb.WatchTokenRequest, stream pb.TokenUpvoteService_WatchTokenVotesServer) error {

	godotenv.Load()
	dbname := os.Getenv("DBNAME")
	coll := mongodb.DB.Database(dbname).Collection("vote")
	pipeline := mongo.Pipeline{}

	vote_stream, err := coll.Watch(context.TODO(), pipeline)
	if err != nil {
		panic(err)
	}
	type NS struct {
	}
	type ChangeStream struct {
		Id            primitive.ObjectID
		OperationType string
		fullDocument  bson.D
	}
	defer vote_stream.Close(context.TODO())
	for vote_stream.Next(context.TODO()) {
		var data bson.M
		if err := vote_stream.Decode(&data); err != nil {
			panic(err)
		}

		result := mongodb.ListDocumentsByX(0, 0, "vote", "to", in.GetId())

		n_votes := &pb.NumberOfVotes{
			Owner:  in.GetId(),
			NVotes: uint32(len(result)),
		}
		stream.Send(n_votes)
	}
	return nil
}
