package iota

type ReadingStore struct {
	Readings map[string]string
}

func NewReadingStore() ReadingStore {
	readings := make(map[string]string)
	return ReadingStore{readings }
}

func (store *ReadingStore) AddReading(readingId string, sensorId string) {
	store.Readings[readingId] = sensorId
}

func (store *ReadingStore) GetNext() (string, string) {
	for readingId, sensorId := range store.Readings {
		return readingId, sensorId
	}
	return "", ""
}

func (store *ReadingStore) Remove(readingId string) {
	delete(store.Readings, readingId)
}