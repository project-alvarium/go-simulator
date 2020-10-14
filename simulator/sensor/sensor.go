package sensor

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/project-alvarium/go-simulator/collections"
)

type Sensor struct {
}

func (sn Sensor) Schedule(delay time.Duration) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			StoreRawData()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func StoreRawData() {
	insertResult, err := collections.InsertRawData(rand.Int63())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Rawdata: ", insertResult)

}
