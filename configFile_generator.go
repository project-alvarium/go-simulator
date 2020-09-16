package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type ConfigFile struct {
	SensorName        string
	GatewayName       string
	ServerName        string
	StorageName       string
	SensorType        string
	TangleLocation    string
	AnnotationOwners  []string
	Annotations       []string
	IOTAStreamId      string
	EmissionFrequency int64 `json:"ref"`
	// private string // An unexported field is not encoded.
	Created time.Time
}

func main() {
	CF1 := ConfigFile{
		SensorName:        "TestSensor",
		GatewayName:       "TestGateWay",
		ServerName:        "TestServer",
		StorageName:       "TestStorage",
		SensorType:        "Binary",
		TangleLocation:    "Testttt",
		AnnotationOwners:  []string{"apple", "ibm", "dell"},
		Annotations:       []string{"policy", "ownership", "date"},
		IOTAStreamId:      "s7g37gd",
		EmissionFrequency: 10,
		Created:           time.Now(),
	}

	var jsonData []byte
	jsonData, err := json.Marshal(CF1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	writeToFile(string(jsonData))

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
