package annotator

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/project-alvarium/go-simulator/libs"

	"github.com/project-alvarium/go-simulator/collections"
	"github.com/project-alvarium/go-simulator/configuration"

	"github.com/dgrijalva/jwt-go"
)

type Annotation struct {
	Name  string
	Score float64
}

func (an Annotation) StoreAnnotation() {

	insertResult, err := collections.InsertAnnotation("Policy", rand.Float64())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult)
}

func (an Annotation) RetrieveAnnotation(sensorId string) {

	result, err := collections.FindAnnotation("Ownership")

	fmt.Printf("Found a single document: %+v\n", result, err)

	var token2 = setJWT(result, sensorId)
	fmt.Print("\nThe JWT is:\n", token2)

	tokenString := token2
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(configuration.Config.Secret), nil
	})
	fmt.Print("\nVerified would be\n:", token.Claims)
	for key, val := range claims {
		fmt.Printf("\nKey: %v, value: %v\n", key, val)
	}

}

func setJWT(ann collections.Annotation, sensorId string) string {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", configuration.Config.Secret) //this should be in an env file
	rl := libs.RandLib{Charset: "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"}
	atClaims := jwt.MapClaims{}
	atClaims["iss"], err = os.Hostname()
	atClaims["sub"] = sensorId
	atClaims["iat"] = time.Now()
	atClaims["jti"] = rl.StringWithCharset(10)
	atClaims["ann"] = ann.Name
	atClaims["avl"] = ann.Score
	// atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return ""
	}
	return token
}
