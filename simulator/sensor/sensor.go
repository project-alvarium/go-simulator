package sensor

import (
	"fmt"
	"github.com/project-alvarium/go-simulator/collections"
	"github.com/project-alvarium/go-simulator/configuration"
	"github.com/project-alvarium/go-simulator/libs"
	"log"

	"github.com/project-alvarium/go-simulator/iota"
	"github.com/project-alvarium/go-simulator/simulator/configfile"

	"math/rand"
	"strconv"
	"time"
)

type Sensor struct {
	subscriber *iota.Subscriber
	config configfile.ConfigFile
}

func NewSensor(subscriber *iota.Subscriber, cf configfile.ConfigFile) Sensor {
	return Sensor{ subscriber, cf }
}

func (sn Sensor) Schedule(delay time.Duration) {
	for {
		sn.storeRawData()
		time.Sleep(delay * time.Second)
	}
}

func (sn Sensor) storeRawData() {
	data := rand.Int63()
	rl := libs.RandLib{Charset: configuration.LetterBytes}
	readingId := rl.StringWithCharset(10)

	fmt.Println("Sending data ", data, " from ", sn.config.SensorID)
	readingMessage := iota.NewReading(
		sn.config.SensorID,
		readingId,
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
