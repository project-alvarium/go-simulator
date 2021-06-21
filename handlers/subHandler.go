package handlers

import (
	"github.com/project-alvarium/go-simulator/iota"
	"github.com/project-alvarium/go-simulator/simulator/annotator"
	"github.com/project-alvarium/go-simulator/simulator/configfile"
	"github.com/project-alvarium/go-simulator/simulator/sensor"
)

func CreateSubscriber(cf configfile.ConfigFile) iota.Subscriber {
	sub := iota.NewSubscriber(&cf.NodeConfig, &cf.SubConfig)
	sub.AwaitKeyload()

	return sub
}

func CreateSensor(sub *iota.Subscriber, cf configfile.ConfigFile, readings *iota.ReadingStore) sensor.Sensor {
	return sensor.NewSensor(sub, cf.SensorName, readings)
}

func CreateAnnotator(sub *iota.Subscriber, cf configfile.ConfigFile, readings *iota.ReadingStore) annotator.Annotator {
	return annotator.NewAnnotator(sub, cf, readings)
}
