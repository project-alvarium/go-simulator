package main

/*
#cgo CFLAGS: -I./iota/include -DIOTA_STREAMS_CHANNELS_CLIENT
//Choose one of the 2 below for compilation. Use .so for linux and .dylib for mac
#cgo LDFLAGS: ./iota/include/libiota_streams_c.so
#include <channels.h>
*/
import "C"
import (
	"fmt"
	"github.com/project-alvarium/go-simulator/iota"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/project-alvarium/go-simulator/api"
	"github.com/project-alvarium/go-simulator/configuration"
)

func main() {
	var subs = iota.NewSubStore()
	var readings = iota.NewReadingStore()

	SetupShutdownHandler(subs)
	//VERY simple demonstration that the IOTA C bindings are included and callable
	C.drop_str(C.CString("A"))
	//After "make build" and "make run", you will see the statement below indicating the
	//above call was made successfully even though it doesn't do anything.
	log.Println("Starting go-simulator...")
	httpRouter := api.NewRouter(&subs, &readings)
	configuration.InitConfig()
	srv := &http.Server{
		Handler: httpRouter,
		Addr:    "127.0.0.1:" + fmt.Sprint(configuration.Config.HTTPPort),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

/*
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
     */

	// Create a new sensor with subscriber embedded
	//newSensor := sensor.NewSensor(&sensorSubscriber, cf, readingIds)
	// Create a new annotator with subscriber embedded
	//newAnnotator := annotator.NewAnnotator(&annSubscriber, cf, readingIds)

	// Schedule emissions
	//go newSensor.Schedule(time.Duration(cf.EmissionFrequency))
	//go newAnnotator.Schedule(time.Duration(cf.EmissionFrequency))

	//collections.Database()
	//annotator.RetrieveAnnotation(cf.SensorID)

	log.Println("listening")
	log.Fatal(srv.ListenAndServe())
}



func SetupShutdownHandler(subs iota.SubStore) {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		log.Println("Shutdown called")
		log.Println("Dropping Subscribers")
		subs.DropSubs()
		log.Println("Dropped")
		log.Println("Exiting...")
		os.Exit(0)
	}()
}
