package annotator

import (
	"fmt"
	"github.com/project-alvarium/go-simulator/iota"
	"github.com/project-alvarium/go-simulator/libs"
	"github.com/project-alvarium/go-simulator/simulator/configfile"
	"math/rand"
	"os"
	"time"

	"github.com/project-alvarium/go-simulator/collections"
	"github.com/project-alvarium/go-simulator/configuration"

	"github.com/dgrijalva/jwt-go"
)

type Annotator struct {
	sub *iota.Subscriber
	config configfile.ConfigFile
	readingStore *iota.ReadingStore
}

var annotated []string

func NewAnnotator(sub *iota.Subscriber, config configfile.ConfigFile, readings *iota.ReadingStore) Annotator {
	fmt.Println("Made a new annotator")
	return Annotator{sub, config, readings }
}

func (annotator *Annotator) Schedule(delay time.Duration) {
	for i := 0; i < 1000; i++ {
		readingId, sensorId := annotator.readingStore.GetNext()

		if readingId != "" {
			cf := &annotator.config
			for y:= 0; y < 4; y++ {
				if rand.Intn(2) == 1 {
					annotator.StoreAnnotation(sensorId, readingId, cf.Annotations[y])
					time.Sleep(3 * time.Second)
				}
			}
			annotator.readingStore.Remove(readingId)
		}
		time.Sleep(delay * time.Second)
	}
}
func (annotator *Annotator) StoreAnnotation(sensorId string, readingId string, annotation collections.Annotation) {
	fmt.Println("Sending annotation for ", readingId, " from ", sensorId)

	rl := libs.RandLib{Charset: configuration.LetterBytes}
	iss, _ := os.Hostname()
	iat := time.Now().UnixNano() / int64(time.Millisecond)
	an := annotation
	an.Iss = iss
	an.Sub = sensorId
	an.Iat = iat
	an.Jti = rl.StringWithCharset(10)
	an.Avl = 2

	annotationMessage := iota.NewAnnotation(readingId, an)
	annotator.sub.SendMessage(annotationMessage)
/*
	insertResult, err := collections.InsertAnnotation(an)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult)*/
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
