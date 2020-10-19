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

func (sn Sensor) Schedule(delay time.Duration) {
	for {
		storeRawData()
		time.Sleep(delay * time.Second)
	}
}

func storeRawData() {
	insertResult, err := collections.InsertRawData(rand.Int63())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Rawdata: ", insertResult)

}
