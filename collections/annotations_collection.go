package collections

import (
	"context"
	"log"

	"github.com/project-alvarium/go-simulator/configuration"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Annotation struct {
	Iss string  `json:"iss" bson:"iss"`
	Sub string  `json:"sub" bson:"sub"`
	Iat int64   `json:"iat" bson:"iat"`
	Jti string  `json:"jti" bson:"jti"`
	Ann string  `json:"ann" bson:"ann"`
	Avl float64 `json:"avl" bson:"avl"`
}

var annotationCollection string = "annotations"

//InsertAnnotation : inserts new annotation
func InsertAnnotation(annotation Annotation) (string, error) {
	var collection = DatabaseClient.Database(configuration.Config.DatabaseName).Collection(annotationCollection)
	insertResult, err := collection.InsertOne(context.TODO(), annotation)
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
