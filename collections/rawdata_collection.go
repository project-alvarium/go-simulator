package collections

import (
	"context"
	"log"
	"sim/go-simulator/configuration"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rawdata struct {
	Data int64 `json:"data" bson:"data"`
}

var rawdataCollection string = "rawdata"

//InsertAnnotation : inserts new annotation
func InsertRawData(data int64) (string, error) {
	var collection = DatabaseClient.Database(configuration.Config.DatabaseName).Collection(rawdataCollection)
	insertResult, err := collection.InsertOne(context.TODO(), Rawdata{data})
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	oid := insertResult.InsertedID.(primitive.ObjectID)
	return oid.String(), nil
}

// FindOne : finds and returns one annotation with the name given
func FindRawData(name string) (Rawdata, error) {
	var collection = DatabaseClient.Database(configuration.Config.DatabaseName).Collection(rawdataCollection)

	filter := bson.D{}

	var data Rawdata
	err := collection.FindOne(context.Background(), filter).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data, nil
}
