package mongo

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

type Token struct {
	Id    string  `bson:"_id,omitempty"`
	Name  string  `bson:"name,omitempty"`
	Price float64 `bson:"price,omitempty"`
	Votes uint    `bson:"votes,omitempty"`
}

func ListTokens(limit, page uint) []Token {
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
	oid, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: oid}}
	var result bson.D
	var token Token
	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	doc, _ := bson.Marshal(result)

	bson.Unmarshal(doc, &token)
	return token
}
