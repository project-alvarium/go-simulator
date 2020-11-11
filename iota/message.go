package iota

import (
	"encoding/json"
	"fmt"
	"github.com/project-alvarium/go-simulator/collections"
)

type TangleMessage struct {
	message string
}

func NewReading(sensorId string, readingId string, data string) TangleMessage {
	message := "{ \"sensor_id\": \"" + sensorId + "\", \"reading_id\": \"" + readingId +
		"\", \"data\": \"" + data + "\" }"
	return TangleMessage{ message }
}

func NewAnnotation(readingId string, ann collections.Annotation) TangleMessage {
	message := "{ \"reading_id\": \"" + readingId + "\", \"annotation\": " + AnnotationToString(ann) + " }"
	return TangleMessage{ message }
}

func AnnotationToString(ann collections.Annotation) string {
	return "{ \"header\": " + Header("RS256", "JWT") +
		", \"payload\": " + Payload(ann) +
		", \"signature\": " + Signature("A signature") + " }"
}

func Header(alg string, typ string) string {
	return "{ \"alg\": \"" + alg + "\", \"typ\": \"" + typ + "\" }"
}

func Signature(sig string) string {
	return "\"" + sig + "\""
}

func Payload(ann collections.Annotation) string {
	j, err := json.Marshal(ann)
	if err != nil {
		fmt.Println("Error marshalling annotation json")
	}
	return string(j)
}
