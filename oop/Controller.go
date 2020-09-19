package main

import (
	"oop/configFile"
	"time"
)

func main() {
	c := configFile.ConfigFile{
		SensorName:        "t1",
		GatewayName:       "t2",
		ServerName:        "t3",
		StorageName:       "t4",
		SensorType:        "t5",
		TangleLocation:    "t6",
		AnnotationOwners:  []string{"apple", "ibm", "dell"},
		Annotations:       []string{"policy", "ownership", "date"},
		IOTAStreamId:      "t7",
		EmissionFrequency: 10,
		Created:           time.Now(),
	}
	c.LeavesRemaining()

}
