package mongodb

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var DB *mongo.Client

func ListTokens(limit, page uint32) []Token {

	godotenv.Load()
	dbname := os.Getenv("DBNAME")
	coll := DB.Database(dbname).Collection("tokens")
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(page - 1))

	var results []Token
	cur, err := coll.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem Token
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	return results

}

func GetTokenById(id string) Token {
	godotenv.Load()
	dbname := os.Getenv("DBNAME")
	coll := DB.Database(dbname).Collection("tokens")
	object_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: object_id}}
	var result bson.D
	var token Token
	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	doc, _ := bson.Marshal(result)

	bson.Unmarshal(doc, &token)
	return token
}

func InsertToken(data Token) (Token, error) {
	godotenv.Load()
	dbname := os.Getenv("DBNAME")
	coll := DB.Database(dbname).Collection("tokens")

	inserted_doc, err := coll.InsertOne(context.TODO(), data)

	if err != nil {
		return data, status.Error(codes.Unknown, "err")
	}

	filter := bson.D{{Key: "_id", Value: inserted_doc.InsertedID}}
	var result bson.D
	coll.FindOne(context.TODO(), filter).Decode(&result)

	doc, _ := bson.Marshal(result)

	var token Token
	bson.Unmarshal(doc, &token)

	return token, nil
}
