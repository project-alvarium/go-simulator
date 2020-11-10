package iota

type TangleMessage struct {
	message string
}

func NewReading(sensorId string, readingId string, data string) TangleMessage {
	message := "{ \"sensor_id\": \"" + sensorId + "\", \"reading_id\": \"" + readingId +
		"\", \"data\": \"" + data + "\" }"
	return TangleMessage{ message }
}

func NewAnnotation(readingId string, data string) TangleMessage {
	message := "{ \"reading_id\": \"" + readingId + "\", \"data\":\"" + data + "\" }"
	return TangleMessage{ message }
}
