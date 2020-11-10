package main

/*
#cgo CFLAGS: -I./iota/include -DIOTA_STREAMS_CHANNELS_CLIENT
//Choose one of the 2 below for compilation. Use .so for linux and .dylib for mac
#cgo LDFLAGS: ./iota/include/libiota_streams_c.so
//#cgo LDFLAGS: ./iota/include/libiota_streams_c.dylib
#include <channels.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/project-alvarium/go-simulator/api"
	"github.com/project-alvarium/go-simulator/configuration"
	"github.com/project-alvarium/go-simulator/iota"
	"github.com/project-alvarium/go-simulator/simulator/configfile"
	"github.com/project-alvarium/go-simulator/simulator/sensor"
)

var subs []iota.Subscriber

func main() {
	SetupShutdownHandler(&subs)
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

	for x:=0;x<10;x++ {
		// Create a new configuration for subscriber/sensor
		cf := configfile.ConfigFile{}
		cf.SetConfigurationFile()
		cf = parseData()

		// Create a subscriber instance and await connection
		subscriber := iota.NewSubscriber(cf.NodeConfig, cf.SubConfig)
		subscriber.AwaitKeyload()

		// Add subscriber to array for dropping on shutdown
		subs = append(subs, subscriber)

		// Create a new sensor with subscriber embedded
		new_sensor := sensor.NewSensor(&subscriber, cf)
		go new_sensor.Schedule(time.Duration(cf.EmissionFrequency))
	}
	//collections.Database()
	//annotation := annotator.Annotation{}

	//annotation.StoreAnnotation()
	//annotation.RetrieveAnnotation(cf.SensorID)

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

func SetupShutdownHandler(subs *[]iota.Subscriber) {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		fmt.Println("Shutdown called\nDropping Subscribers")
		for _, sub := range *subs {
			sub.Drop()
		}
		fmt.Println("Dropped\nExiting...")
		os.Exit(0)
	}()
}
