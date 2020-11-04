package main

/*
#cgo CFLAGS: -I./iota
#cgo LDFLAGS: -L./iota -liota_streams_c
#include <channels.h>
*/
import "C"
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

func main() {
	//VERY simple demonstration that the IOTA C bindings are included and callable
	C.drop_str(C.CString("A"))
	//After "make build" and "make run", you will see the statement below indicating the
	//above call was made successfully even though it doesn't do anything.
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
	sensor := sensor.Sensor{}
	go sensor.Schedule(time.Duration(cf.EmissionFrequency))
	annotation := annotator.Annotation{}
	// annotation.StoreAnnotation()
	annotation.RetrieveAnnotation(cf.SensorID)

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
