package configfile

import (
	"encoding/json"
	"fmt"
	"github.com/project-alvarium/go-simulator/collections"
	"github.com/project-alvarium/go-simulator/configuration"
	"github.com/project-alvarium/go-simulator/libs"
	"log"
	"math/rand"
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

type NodeConfig struct {
	Url string
}

type SubConfig struct {
	Seed       string
	Encoding   string
	AnnAddress string
}

type ConfigFile struct {
	SensorID          string
	SensorName        string
	GatewayName       string
	ServerName        string
	StorageName       string
	SensorType        string
	TangleLocation    string
	AnnotationOwners  []Owner
	Annotations       []collections.Annotation
	IOTAStreamID      string
	EmissionFrequency int64 `json:"ef"`
	NodeConfig        NodeConfig
	SubConfig         SubConfig
	// private string // An unexported field is not encoded.
	Created time.Time
}

func (cf *ConfigFile) SetConfigurationFile(name string, subConfig SubConfig, nodeConfig NodeConfig) {

	var cf1 = setData(name, subConfig, nodeConfig)
	cf = &cf1

	var jsonData []byte
	jsonData, err := json.Marshal(&cf)
	if err != nil {
		log.Println(err)
	}

	//After the configuration file data are set, it should be exported in a JSON formated string to be used
	writeToFile(name, string(jsonData))
	//This is our simulator entry point where it reads the configuration file, then parses the required data
	// parseData()
	//Then we move on with the flow

}

func setData(name string, subConfig SubConfig, nodeConfig NodeConfig) ConfigFile {
	rl := libs.RandLib{Charset: configuration.LetterBytes}
	cf := ConfigFile{}
	cf.SensorID = rl.StringWithCharset(8)
	cf.SensorName = name
	cf.GatewayName = "TestGateWay"
	cf.ServerName = "TestServer"
	cf.StorageName = "TestStorage"
	cf.SensorType = "Binary"
	cf.TangleLocation = "Test"
	cf.AnnotationOwners = []Owner{Owner{Name: "IOTA", PrivateKey: "IOTAKey"}, {Name: "IBM", PrivateKey: "IBMKey"}, {Name: "Dell", PrivateKey: "DellKey"}}
	cf.Annotations = []collections.Annotation{{Ann: "policy"}, {Ann: "ownership"}, {Ann: "boot"}, {Ann: "store"}}
	cf.IOTAStreamID = "s7g37gd"
	cf.EmissionFrequency = 10
	cf.Created = time.Now()
	cf.SubConfig = subConfig
	cf.NodeConfig = nodeConfig

	return cf
}

func writeToFile(name, s string) {
	f, err := os.Create(fmt.Sprintf("%s-test.txt", name))
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

func NewNodeConfig(url string) NodeConfig {
	return NodeConfig{url}
}

func NewSubConfig(annAddress string) SubConfig {
	bytes := make([]byte, 64)
	rand.Seed(time.Now().UnixNano())
	for i := range bytes {
		bytes[i] = configuration.LetterBytes[rand.Intn(len(configuration.LetterBytes))]
	}

	seed := string(bytes)
	fmt.Println("Seed: ", seed)
	encoding := "utf-8"

	return SubConfig{seed, encoding, annAddress}
}
