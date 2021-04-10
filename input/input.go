package input

import (
	"encoding/json"
	"feedback/models"
	"fmt"
	"io/ioutil"
	"os"
)

type Input interface {
	NextEvent() *models.Event
}

type FileInput struct {
	FileName string
}

// NextEvent Returns input from a sample file
func (f *FileInput) NextEvent() *models.Event {
	jsonFile, err := os.Open(f.FileName)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, errRead := ioutil.ReadAll(jsonFile)

	if errRead != nil {
		fmt.Println(errRead)
	}

	var bundle Bundle

	errJson := json.Unmarshal(byteValue, &bundle)

	if errJson != nil {
		fmt.Println(errJson)
	}

	return createEvent(bundle)
}

func createEvent(bundle Bundle) *models.Event {
	event := &models.Event{}

	for _, entry := range bundle.Entry {
		switch t := entry.Resource.(type) {
		case *models.Patient:
			event.Patient = *t
		case *models.Doctor:
			event.Doctor = *t
		case *models.Diagnosis:
			event.Diagnosis = *t
		case *models.Appointment:
			event.Appointment = *t
		}
	}

	return event
}
