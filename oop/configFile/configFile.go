package configFile

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
	IOTAStreamId      string
	EmissionFrequency int64 `json:"ref"`
	// private string // An unexported field is not encoded.
	Created time.Time
}

func (e ConfigFile) LeavesRemaining() {
	fmt.Printf("%s %s has %d names \n", e.SensorName, e.GatewayName)
}

func (cf ConfigFile) createRandomFile(t1 string, t2 string, t3 string, t4 string, t5 string, t6 string, t7 string, t8 int64) {
	CF1 := ConfigFile{
		SensorName:        t1,
		GatewayName:       t2,
		ServerName:        t3,
		StorageName:       t4,
		SensorType:        t5,
		TangleLocation:    t6,
		AnnotationOwners:  []string{"apple", "ibm", "dell"},
		Annotations:       []string{"policy", "ownership", "date"},
		IOTAStreamId:      t7,
		EmissionFrequency: t8,
		Created:           time.Now(),
	}

	var jsonData []byte
	jsonData, err := json.Marshal(CF1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	writeToFile(string(jsonData))
	// readFromFile("test.txt")

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

func readFromFile(filename string) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
