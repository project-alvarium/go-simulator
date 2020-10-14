package configfile

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Owner struct {
	Name       string
	PrivateKey string
}

type Annotation struct {
	Name     string
	NodePath string
	Owner    Owner
}

type ConfigFile struct {
	SensorName        string
	GatewayName       string
	ServerName        string
	StorageName       string
	SensorType        string
	TangleLocation    string
	AnnotationOwners  []Owner
	Annotations       []Annotation
	IOTAStreamID      string
	EmissionFrequency int64 `json:"ef"`
	// private string // An unexported field is not encoded.
	Created time.Time
}

func (cf ConfigFile) SetConfigurationFile() {

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
	// parseData()
	//Then we move on with the flow

}

func setRandomData() ConfigFile {
	cf := ConfigFile{}
	cf.SensorName = "TestSensor3"
	cf.GatewayName = "TestGateWay"
	cf.ServerName = "TestServer"
	cf.StorageName = "TestStorage"
	cf.SensorType = "Binary"
	cf.TangleLocation = "Testttt"
	cf.AnnotationOwners = []Owner{Owner{Name: "IOTA", PrivateKey: "IOTAKey"}, {Name: "IBM", PrivateKey: "IBMKey"}, {Name: "Dell", PrivateKey: "DellKey"}}
	cf.Annotations = []Annotation{Annotation{Name: "policy", NodePath: "policyNodePath", Owner: Owner{Name: "IOTA", PrivateKey: "IOTAKey"}}, Annotation{Name: "ownership", NodePath: "ownershipNodePath", Owner: Owner{Name: "Dell", PrivateKey: "DellKey"}}}
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
