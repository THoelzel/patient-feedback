package user_input

import (
	"feedback/models"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserInput(t *testing.T) {

	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, "7\ny\nNo feelings\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatal(err)
	}
	e := &models.Event{
		Patient:   models.Patient{Name: []models.Name{{Given: []string{"John"}}}},
		Doctor:    models.Doctor{Name: []models.Name{{Family: "Sample"}}},
		Diagnosis: models.Diagnosis{Code: models.Code{Coding: []models.DiagnosisDetails{{Name: "Unspecified Illness"}}}},
	}

	f := UserInput(e, in)

	assert.Equal(t, 7, f.Rating)
	assert.True(t, f.DiagnosisUnderstood)
	//assert.Equal(t, "Use simpler language", f.Improvement)
	assert.Equal(t, "No feelings", f.Feelings)
}
