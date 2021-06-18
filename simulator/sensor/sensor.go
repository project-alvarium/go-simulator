package sensor

import (
	"github.com/project-alvarium/go-simulator/configuration"
	"github.com/project-alvarium/go-simulator/iota"
	"github.com/project-alvarium/go-simulator/libs"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type Sensor struct {
	subscriber *iota.Subscriber
	name string
	readingStore *iota.ReadingStore
}

func NewSensor(subscriber *iota.Subscriber, name string, readingStore *iota.ReadingStore) Sensor {
	return Sensor{ subscriber, name, readingStore }
}

func (sn Sensor) Schedule(delay time.Duration) {
	for i:= 0; i < 1000; i++ {
		sn.storeRawData()
		time.Sleep(delay * time.Second)
	}
}

func (sn *Sensor) storeRawData() {
	data := rand.Int63()

	// Prepare reading Id's in advance
	rl := libs.RandLib{Charset: configuration.LetterBytes}
	readingId := rl.StringWithCharset(10)

	log.Println("Sending ", readingId, " from ", sn.name)

	readingMessage := iota.NewReading(
		sn.name,
		readingId,
		strconv.FormatInt(data, 10),
		)

	sn.readingStore.AddReading(readingId, sn.name)
	sn.subscriber.SendMessage(readingMessage)

	/// **** Note: Does this reading ID approach work for your end? I'm not sure what the
	/// plan for that sensor insertion/annotation is going to be. I would propose storing
	/// them as a key/val mapping of readingId -> data for simplification purposes

/*
	insertResult, err := collections.InsertRawData(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Rawdata: ", insertResult)
*/
}
