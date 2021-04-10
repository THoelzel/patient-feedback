package main

import (
	"feedback/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"syscall"
	"testing"
)

type mockFeedbackModel struct{}

func (m *mockFeedbackModel) Stats() (s []models.Stats) {

	return []models.Stats{
		{Rating: 5.5, Understandable: 0.8, Diagnosis: "A1.0"},
	}
}

func (m *mockFeedbackModel) Insert(f *models.Feedback) {

}

func TestDisplayStats(t *testing.T) {
	env = &Env{
		feedback: &mockFeedbackModel{},
	}

	out, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()
	os.Stdout = out

	// defer stdout reset in case it fails before the later call
	defer resetStdOut()
	DisplayStats()

	b, _ := ioutil.ReadFile(out.Name())

	// reset stdout so the output of the assert is not hidden
	resetStdOut()

	expected := `+-----------------------+----------------+-----------------------+
| ICD-10 DIAGNOSIS CODE | AVERAGE RATING | RATE OF UNDERSTANDING |
+-----------------------+----------------+-----------------------+
| A1.0                  |          5.500 |                 0.800 |
+-----------------------+----------------+-----------------------+
`

	assert.Equal(t, expected, string(b))
}

func resetStdOut() {
	os.Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
}
