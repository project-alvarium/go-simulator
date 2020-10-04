package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sim/go-simulator/api"
	"sim/go-simulator/collections"
	"sim/go-simulator/configuration"
	"sim/go-simulator/simulator/annotator"
	"sim/go-simulator/simulator/configfile"
	"sim/go-simulator/simulator/sensor"
	"time"
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
