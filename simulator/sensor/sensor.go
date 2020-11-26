package sensor

import (
	"fmt"
	"github.com/project-alvarium/go-simulator/collections"
	"github.com/project-alvarium/go-simulator/iota"
	"github.com/project-alvarium/go-simulator/simulator/configfile"
	"log"

	"math/rand"
	"strconv"
	"time"
)

type Sensor struct {
	subscriber *iota.Subscriber
	config configfile.ConfigFile
	count int
}

func NewSensor(subscriber *iota.Subscriber, cf configfile.ConfigFile) Sensor {
	return Sensor{ subscriber, cf, 0 }
}

func (sn Sensor) Schedule(delay time.Duration) {
	for {
		sn.storeRawData()
		time.Sleep(delay * time.Second)
	}
}

func (sn *Sensor) storeRawData() {
	data := rand.Int63()
	sn.count += 1

	fmt.Println("Sending data ", data, " from ", sn.config.SensorName)
	readingMessage := iota.NewReading(
		sn.config.SensorName,
		strconv.Itoa(sn.count),
		strconv.FormatInt(data, 10),
		)

	sn.subscriber.SendMessage(readingMessage)

	/// **** Note: Does this reading ID approach work for your end? I'm not sure what the
	/// plan for that sensor insertion/annotation is going to be. I would propose storing
	/// them as a key/val mapping of readingId -> data for simplification purposes


	insertResult, err := collections.InsertRawData(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Rawdata: ", insertResult)

}
