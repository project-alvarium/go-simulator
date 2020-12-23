![Project Alvarium](README.assets/ProjectAlvarium.png)
# Simulator Example (Golang)

In this readme you shall be able to run a basic simulator that stores random annotations, into our previously setup tangle


For simulation purposes, we have a function that outputs a random config file for parsing, you can edit the values directly.

## Direct to simulator/configfile and access the gen.go file

```golang
type ConfigFile struct {
	SensorID          string
	SensorName        string
	GatewayName       string
	ServerName        string
	StorageName       string
	SensorType        string
	TangleLocation    string
	AnnotationOwners  []Owner
	Annotations       []collections.Annotation
	IOTAStreamID      string
	EmissionFrequency int64 `json:"ef"`
	NodeConfig        NodeConfig
	SubConfig         SubConfig
	// private string // An unexported field is not encoded.
	Created time.Time
}
```
This is the struct defining the fields of the expected configfile, you can add or remove based on what the actual ones are.

## Edit the ConfigFile Values based on prefrences
```golang

func setRandomData() ConfigFile {
	rl := libs.RandLib{Charset: configuration.LetterBytes}
	cf := ConfigFile{}
	cf.SensorID = rl.StringWithCharset(8)
	cf.SensorName = "TestSensor3"
	cf.GatewayName = "TestGateWay"
	cf.ServerName = "TestServer"
	cf.StorageName = "TestStorage"
	cf.SensorType = "Binary"
	cf.TangleLocation = "Test"
	cf.AnnotationOwners = []Owner{Owner{Name: "IOTA", PrivateKey: "IOTAKey"}, {Name: "IBM", PrivateKey: "IBMKey"}, {Name: "Dell", PrivateKey: "DellKey"}}
	cf.Annotations = []collections.Annotation{{Ann: "policy"}, {Ann: "ownership"}}
	cf.IOTAStreamID = "s7g37gd"
	cf.EmissionFrequency = 10
	cf.Created = time.Now()
	cf.SubConfig = NewSubConfig(configuration.AnnAddress)
	cf.NodeConfig = NewNodeConfig(configuration.NodeUrl, configuration.NodeMwm)

	return cf

}
```

The function "setRandomData()" simply sets initial random values to the fields, these would be replaced later on with the actual values of the edge device.

## Direct to main.go 


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
