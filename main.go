package main

/*
#cgo CFLAGS: -I./iota/include -DIOTA_STREAMS_CHANNELS_CLIENT
//Choose one of the 2 below for compilation. Use .so for linux and .dylib for mac
#cgo LDFLAGS: ./iota/include/libiota_streams_c.so
//#cgo LDFLAGS: -L./iota/include -liota_streams_c
#include <channels.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/project-alvarium/go-simulator/libs"
	"github.com/project-alvarium/go-simulator/simulator/annotator"
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

	// Create a new configuration for subscriber/sensor
	cf := configfile.ConfigFile{}
	cf.SetConfigurationFile()
	cf = parseData()

	// Create a subscriber instance for annotator and await connection
	sensorSubscriber := iota.NewSubscriber(cf.NodeConfig, cf.SubConfig)
	sensorSubscriber.AwaitKeyload()

	// Add subscriber to array for dropping on shutdown
	subs = append(subs, sensorSubscriber)

	// Create a new configuration for annotator
	cf2 := configfile.ConfigFile{}
	cf2.SetConfigurationFile()
	cf2 = parseData()

	// Create a subscriber instance for annotator and await connection
	annSubscriber := iota.NewSubscriber(cf2.NodeConfig, cf2.SubConfig)
	annSubscriber.AwaitKeyload()

	// Add subscriber to array for dropping on shutdown
	subs = append(subs, annSubscriber)

	// Create a new sensor with subscriber embedded
	newSensor := sensor.NewSensor(&sensorSubscriber, cf)
	// Create a new annotator with subscriber embedded
	newAnnotator := annotator.NewAnnotator(&annSubscriber)
	go newSensor.Schedule(time.Duration(cf.EmissionFrequency))

	rl := libs.RandLib{Charset: "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" }
	newAnnotator.StoreAnnotation(cf.SensorID, rl.StringWithCharset(8))

	//collections.Database()
	//annotator.RetrieveAnnotation(cf.SensorID)

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
