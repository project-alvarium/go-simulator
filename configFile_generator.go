package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	IOTAStreamID      string
	EmissionFrequency int64 `json:"ref"`
	// private string // An unexported field is not encoded.
	Created time.Time
}

func main() {

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
