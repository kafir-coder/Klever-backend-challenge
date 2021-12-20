package token

import (
	"Token_service/pkg/mongodb"
	pb "Token_service/pkg/proto/token"
	"context"
	"fmt"

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

	token := mongodb.Token{
		Name:  req.GetName(),
		Price: req.GetPrice(),
	}

	inserted_document, _ := mongodb.InsertToken(token)

	result := &pb.Token{
		Id:    inserted_document.Id,
		Name:  inserted_document.Name,
		Price: inserted_document.Price,
	}
	return result, nil

}
