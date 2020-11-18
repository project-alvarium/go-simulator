![Project Alvarium](README.assets/ProjectAlvarium.png)
# Simulator Example (Golang)

In this readme you shall be able to run a basic simulator that stores random annotations, into our previously setup tangle

## Direct to where the main.go file resides


```golang

  cf2 := configfile.ConfigFile{}
	cf2.SetConfigurationFile()
	cf2 = parseData()

```
In here you will find that the entry point will begin by exporting a config file and get the parsed values

```golang
	// Create a subscriber instance for annotator and await connection
	annSubscriber := iota.NewSubscriber(cf2.NodeConfig, cf2.SubConfig)
	annSubscriber.AwaitKeyload()

	// Add subscriber to array for dropping on shutdown
	subs = append(subs, annSubscriber)

```
Setting the iota stream to subscribe to the annotations stream

```golang
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
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"}
	newAnnotator.StoreAnnotation(cf.SensorID, rl.StringWithCharset(8))
```
Finally, initializing the sensor loop and storing the outcome annotation.

## Run the following commands in the main directory:
If running for the first time
```
make build
```
then
```
make run
```

After building for the first time you can use: 
```
go run main.go
```
