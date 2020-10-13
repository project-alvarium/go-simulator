package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/project-alvarium/go-simulator/api"
	"github.com/project-alvarium/go-simulator/collections"
	"github.com/project-alvarium/go-simulator/configuration"
	"github.com/project-alvarium/go-simulator/simulator/annotator"
	"github.com/project-alvarium/go-simulator/simulator/configfile"
	"github.com/project-alvarium/go-simulator/simulator/sensor"

)

func simulateSensor(frequency int64) {
	sensor := sensor.Sensor{}

	stop := sensor.Schedule(time.Duration(frequency) * time.Second)
	time.Sleep(25 * time.Second)
	stop <- true
	time.Sleep(25 * time.Second)

	fmt.Println("Done")
}

func main() {
	fmt.Println("Starting go-simulator...")
	httpRouter := api.NewRouter()
	configuration.InitConfig()
	srv := &http.Server{
		Handler: httpRouter,
		Addr:    "127.0.0.1:" + fmt.Sprint(configuration.Config.HTTPPort),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	collections.Database()
	cf := configfile.ConfigFile{}
	cf.SetConfigurationFile()
	cf = parseData()
	simulateSensor(cf.EmissionFrequency)
	annotation := annotator.Annotation{}
	// annotation.StoreAnnotation()
	annotation.RetrieveAnnotation()

	log.Fatal(srv.ListenAndServe())
	log.Println("listening")
}

func readFromFile(filename string) string {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}

	return string(data)
}

func parseData() configfile.ConfigFile {
	var configuartions = readFromFile("test.txt")
	var Data configfile.ConfigFile
	json.Unmarshal([]byte(configuartions), &Data)
	fmt.Print("Example for config file fields: \n The Sensor Name is: ", Data.GatewayName, "\n")
	return Data

}
