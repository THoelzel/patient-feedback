package models

type Name struct {
	Text   string   `json:"text"`
	Family string   `json:"family"`
	Given  []string `json:"given"`
}

type Patient struct {
	Resource
	Name []Name `json:"name"`
}

type Doctor struct {
	Resource
	Name []Name `json:"name"`
}

type Appointment struct {
	Resource
}

type Diagnosis struct {
	Status string `json:"status"`
	Code   Code   `json:"code"`
	Resource
}

type Code struct {
	Coding []DiagnosisDetails `json:"coding"`
}

type DiagnosisDetails struct {
	System string `json:"system"`
	Code   string `json:"code"`
	Name   string `json:"name"`
}

type Resource struct {
	ResourceType string `json:"resourceType"`
	Id           string `json:"id"`
}

type Event struct {
	Patient     Patient
	Doctor      Doctor
	Appointment Appointment
	Diagnosis   Diagnosis
}
