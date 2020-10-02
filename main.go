package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sim/go-simulator/simulator/annotator"
	"sim/go-simulator/simulator/configfile"
)

func main() {

	// storeAnnotation()
	cf := configfile.ConfigFile{}
	cf.SetConfigurationFile()
	parseData()
	an := annotator.Annotation{"Policy", rand.Float64()}
	// an.StoreAnnotation()
	an.RetrieveAnnotation()

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
	var Data configfile.ConfigFile
	json.Unmarshal([]byte(configuartions), &Data)
	fmt.Print("The Sensor Name is: ", Data.GatewayName)
}
