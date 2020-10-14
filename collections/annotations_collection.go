package collections

import (
	"context"
	"log"

	"github.com/project-alvarium/go-simulator/configuration"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Annotation struct {
	Name  string  `json:"name" bson:"name"`
	Score float64 `json:"score" bson:"score"`
}

var annotationCollection string = "annotations"

//InsertAnnotation : inserts new annotation
func InsertAnnotation(name string, score float64) (string, error) {
	var collection = DatabaseClient.Database(configuration.Config.DatabaseName).Collection(annotationCollection)
	insertResult, err := collection.InsertOne(context.TODO(), Annotation{Name: name, Score: score})
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	oid := insertResult.InsertedID.(primitive.ObjectID)
	return oid.String(), nil
}

// FindOne : finds and returns one annotation with the name given
func FindAnnotation(name string) (Annotation, error) {
	var collection = DatabaseClient.Database(configuration.Config.DatabaseName).Collection(annotationCollection)

	filter := bson.D{}

	var annotation Annotation
	err := collection.FindOne(context.Background(), filter).Decode(&annotation)
	if err != nil {
		log.Fatal(err)
	}
	return annotation, nil
}
