package annotator

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"github.com/project-alvarium/go-simulator/collections"
	"github.com/project-alvarium/go-simulator/configuration"
	"time"

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

func (an Annotation) RetrieveAnnotation() {

	result, err := collections.FindAnnotation("Ownership")

	fmt.Printf("Found a single document: %+v\n", result, err)

	var token2 = setJWT(result)
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

func setJWT(ann collections.Annotation) string {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", configuration.Config.Secret) //this should be in an env file
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
