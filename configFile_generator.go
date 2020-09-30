package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"context"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Annotation struct {
	Name  string
	Score float64
}

func main() {

	// storeAnnotation()
	retrieveAnnotation()

	// setConfigurationFile()

}

func storeAnnotation() {

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

	ruan := Annotation{"Ownership", rand.Float64()}
	fmt.Print(ruan)

	insertResult, err := collection.InsertOne(ctx, ruan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult)
}

func retrieveAnnotation() {
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
	findOptions.SetLimit(1)

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

type ConfigFile struct {
	SensorName        string
	GatewayName       string
	ServerName        string
	StorageName       string
	SensorType        string
	TangleLocation    string
	AnnotationOwners  []string
	Annotations       []string
	IOTAStreamID      string
	EmissionFrequency int64 `json:"ref"`
	// private string // An unexported field is not encoded.
	Created time.Time
}

func setConfigurationFile() {

	var CF1 = setRandomData()

	var jsonData []byte
	jsonData, err := json.Marshal(CF1)
	if err != nil {
		log.Println(err)
	}
	// fmt.Println(string(jsonData))
	//After the configuration file data are set, it should be exported in a JSON formated string to be used
	writeToFile(string(jsonData))
	//This is our simulator entry point where it reads the configuration file, then parses the required data
	parseData()
	//Then we move on with the flow

}

func setRandomData() ConfigFile {
	cf := ConfigFile{}
	cf.SensorName = "TestSensor2"
	cf.GatewayName = "TestGateWay"
	cf.ServerName = "TestServer"
	cf.StorageName = "TestStorage"
	cf.SensorType = "Binary"
	cf.TangleLocation = "Testttt"
	cf.AnnotationOwners = []string{"apple", "ibm", "dell"}
	cf.Annotations = []string{"policy", "ownership", "date"}
	cf.IOTAStreamID = "s7g37gd"
	cf.EmissionFrequency = 10
	cf.Created = time.Now()

	return cf

}

func writeToFile(s string) {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(s)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func readFromFile(filename string) string {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}

	return string(data)
}

func parseData() {
	var configuartions = readFromFile("test.txt")
	// fmt.Println("Contents of file:", configuartions)
	var Data ConfigFile
	json.Unmarshal([]byte(configuartions), &Data)
	fmt.Print("The Sensor Name is: ", Data.GatewayName)
}
