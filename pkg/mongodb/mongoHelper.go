package mongodb

import (
	"context"
	"fmt"
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

func ListDocuments(limit, page uint32, collectionName string) []bson.D {

	godotenv.Load()
	dbname := os.Getenv("DBNAME")
	coll := DB.Database(dbname).Collection(collectionName)
	findOptions := options.Find()
	if limit > 0 && page > 0 {

		findOptions.SetLimit(int64(limit))
		findOptions.SetSkip(int64(page - 1))
	}
	var results []bson.D
	cur, err := coll.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem bson.D
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

func ListDocumentsByX(limit, page uint32, collectionName string, field, filter string) []bson.D {

	godotenv.Load()
	dbname := os.Getenv("DBNAME")
	coll := DB.Database(dbname).Collection(collectionName)
	findOptions := options.Find()
	if limit > 0 && page > 0 {

		findOptions.SetLimit(int64(limit))
		findOptions.SetSkip(int64(page - 1))
	}
	var results []bson.D
	cur, err := coll.Find(context.TODO(), bson.D{{
		Key:   field,
		Value: filter,
	}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem bson.D
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

func GetDocumentById(id, collectionName string) bson.D {
	godotenv.Load()
	dbname := os.Getenv("DBNAME")
	coll := DB.Database(dbname).Collection(collectionName)
	object_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: object_id}}
	var result bson.D
	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		log.Fatalln(err)
	}
	return result
}

func InsertOne(data primitive.D, collectionName string) (bson.D, error) {
	godotenv.Load()
	dbname := os.Getenv("DBNAME")
	coll := DB.Database(dbname).Collection(collectionName)

	inserted_doc, err := coll.InsertOne(context.TODO(), data)

	if err != nil {
		return data, status.Error(codes.Unknown, "err")
	}

	filter := bson.D{{Key: "_id", Value: inserted_doc.InsertedID}}
	var result bson.D
	coll.FindOne(context.TODO(), filter).Decode(&result)

	return result, nil
}

func UpdateById(id, collectionName string, data primitive.D) (*mongo.UpdateResult, error) {
	godotenv.Load()
	dbname := os.Getenv("DBNAME")
	coll := DB.Database(dbname).Collection(collectionName)
	object_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: object_id}}
	update := bson.D{{
		Key:   "$set",
		Value: data,
	}}

	res, err := coll.UpdateOne(
		context.TODO(),
		filter,
		update,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteById(id, collectionName string) (*mongo.DeleteResult, error) {
	godotenv.Load()
	dbname := os.Getenv("DBNAME")
	coll := DB.Database(dbname).Collection(collectionName)
	object_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: object_id}}
	res, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteByX(field, value, collectionName string) (*mongo.DeleteResult, error) {
	godotenv.Load()
	dbname := os.Getenv("DBNAME")
	coll := DB.Database(dbname).Collection(collectionName)

	filter := bson.D{{Key: field, Value: value}}
	res, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func Unmarshal(document bson.D, docType string) interface{} {
	doc, _ := bson.Marshal(document)

	fmt.Println(doc)
	switch {
	case docType == "token":
		var unmarshalled_document Token
		bson.Unmarshal(doc, &unmarshalled_document)
		return unmarshalled_document
	case docType == "vote":
		fmt.Println("sdfs")
		var unmarshalled_document Votes
		bson.Unmarshal(doc, &unmarshalled_document)
		fmt.Println(unmarshalled_document)
		return unmarshalled_document
	}
	return nil
}
