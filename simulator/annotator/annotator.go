package annotator

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Annotation struct {
	Name  string
	Score float64
}

func (an Annotation) StoreAnnotation() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("alvarium").Collection("annotationsRaw")

	fmt.Print("Annotation is:", an)

	insertResult, err := collection.InsertOne(ctx, an)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult)
}

func (an Annotation) RetrieveAnnotation() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("alvarium").Collection("annotationsRaw")
	var result Annotation

	err = collection.FindOne(ctx, bson.D{}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	findOptions := options.Find()
	findOptions.SetLimit(5)

	var token2 = setJWT(result)
	fmt.Print(token2)

	tokenString := token2
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("techdev"), nil
	})
	// ... error handling
	fmt.Print(token)

	// do something with decoded claims
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

}

func setJWT(ann Annotation) string {
	fmt.Print("Initial Output", ann)
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "techdev") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["annotation"] = ann
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return ""
	}
	return token
}
