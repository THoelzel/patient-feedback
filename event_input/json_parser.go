package event_input

import (
	"encoding/json"
	"errors"
	"feedback/logger"
	"feedback/models"
	"fmt"
	"strings"
)

// additional data types required for unmarshalling

type GenericResource interface {
}

type Bundle struct {
	Id           string  `json:"id"`
	ResourceType string  `json:"resourceType"`
	Entry        []Entry `json:"entry"`
}

type Entry struct {
	Resource    GenericResource `json:"-"`
	RawResource json.RawMessage `json:"resource"`
}

func (e *Entry) UnmarshalJSON(b []byte) error {

	type entry Entry

	err := json.Unmarshal(b, (*entry)(e))
	if err != nil {
		fmt.Println(err)
		return err
	}

	var r models.Resource
	errRes := json.Unmarshal(e.RawResource, &r)
	if errRes != nil {
		fmt.Println(errRes)
		return errRes
	}

	var g GenericResource
	switch strings.ToLower(r.ResourceType) {
	case "patient":
		g = &models.Patient{}
	case "doctor":
		g = &models.Doctor{}
	case "diagnosis":
		g = &models.Diagnosis{}
	case "appointment":
		g = &models.Appointment{}
	default:
		logger.Logger().Debugw("Unknown resource type", "resource", r.ResourceType)
		return errors.New("unknown type")
	}

	errGen := json.Unmarshal(e.RawResource, g)
	if errGen != nil {
		fmt.Println(errGen)
		return errGen
	}

	e.Resource = g

	return nil
}
