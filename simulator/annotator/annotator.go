package annotator

import (
	"fmt"
	"github.com/project-alvarium/go-simulator/iota"
	"github.com/project-alvarium/go-simulator/libs"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/project-alvarium/go-simulator/collections"
	"github.com/project-alvarium/go-simulator/configuration"

	"github.com/dgrijalva/jwt-go"
)

type Annotator struct {
	sub *iota.Subscriber
}

func NewAnnotator(sub *iota.Subscriber) Annotator {
	return Annotator{sub}
}

func (annotator Annotator) StoreAnnotation(sensorId string, readingId string, annotation collections.Annotation, annotationName string) {
	rl := libs.RandLib{Charset: "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"}
	iss, _ := os.Hostname()
	iat := time.Now().String()
	an := annotation
	an.Iss = iss
	an.Sub = sensorId
	an.Iat = iat
	an.Jti = rl.StringWithCharset(10)
	an.Ann = annotationName
	an.Avl = rand.Float64()

	annotationMessage := iota.NewAnnotation(readingId, an)
	annotator.sub.SendMessage(annotationMessage)

	insertResult, err := collections.InsertAnnotation(an)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult)
}

func (annotator Annotator) RetrieveAnnotation() {

	result, _ := collections.FindAnnotation("Ownership")

	fmt.Printf("Found a single document: %+v\n", result)

	var token2 = setJWT(result)
	fmt.Print("\nThe JWT is:\n", token2)

	tokenString := token2
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
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
	atClaims["iss"] = ann.Iss
	atClaims["sub"] = ann.Sub
	atClaims["iat"] = ann.Iat
	atClaims["jti"] = ann.Jti
	atClaims["ann"] = ann.Ann
	atClaims["avl"] = ann.Avl
	// atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return ""
	}
	return token
}
