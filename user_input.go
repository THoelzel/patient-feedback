package main

import (
	"bufio"
	"feedback/logger"
	"feedback/models"
	"fmt"
	"log"
	"os"
	"strings"
)

func UserInput(e *models.Event, in *os.File) *models.Feedback {

	if in == nil {
		in = os.Stdin
	}

	var feedback models.Feedback

	feedback.Rating = askForRating(e, in)
	feedback.DiagnosisUnderstood = askIfDiagnosisUnderstood(e, in)

	if !feedback.DiagnosisUnderstood {
		feedback.Improvement = askForSuggestion(e, in)
	}

	feedback.Feelings = askForFeelings(e, in)

	// keep track of the diagnosis this feedback relates to
	feedback.Diagnosis = e.Diagnosis.Code.Coding[0].Code

	printSummary(&feedback, e)

	return &feedback
}

func askForRating(e *models.Event, in *os.File) int {
	q := "Hi %s, on a scale of 1-10, would you recommend Dr %s to a friend or family member? " +
		"1 = Would not recommend, 10 = Would strongly recommend\n"
	fmt.Printf(q, e.Patient.Name[0].Given[0], e.Doctor.Name[0].Family)

	return readInt(in)
}

func askIfDiagnosisUnderstood(e *models.Event, in *os.File) bool {
	q := "Thank you. You were diagnosed with '%s'. " +
		"Did Dr %s explain how to manage this diagnosis in a way you could understand? (y)es/(n)o\n"
	fmt.Printf(q, e.Diagnosis.Code.Coding[0].Name, e.Doctor.Name[0].Family)

	return readBool(in)
}

func askForSuggestion(e *models.Event, in *os.File) string {
	q := "Could you please provide a suggestion as to how Dr. %s could have improved the explanation:\n"
	fmt.Printf(q, e.Doctor.Name[0].Family)

	return readString(in)
}

func askForFeelings(e *models.Event, in *os.File) string {
	q := "We appreciate the feedback, one last question: How do you feel about being diagnosed with '%s'?\n"
	fmt.Printf(q, e.Diagnosis.Code.Coding[0].Name)

	return readString(in)
}

func printSummary(f *models.Feedback, e *models.Event) {
	u := "understandable"
	if !f.DiagnosisUnderstood {
		u = "not " + u
	}

	summary := `Thanks again! Here’s what we heard:
You %[1]s Dr %[2]s.
Dr %[2]s explained the diagnosis in a way that was %[3]s.
Your feelings about the diagnosis can be summarized as:
%[4]s
`

	fmt.Printf(summary, models.FeedbackTranslation[f.Rating], e.Doctor.Name[0].Family, u, f.Feelings)
}

func readBool(in *os.File) bool {
	var response string

	_, err := fmt.Fscanln(in, &response)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		fmt.Println("We're sorry but we didn't get what you meant, please type (y)es or (n)o and then press enter>")
		return readBool(in)
	}
}

func readString(in *os.File) string {
	var s string
	r := bufio.NewReader(in)

	// No input verification as we assume valid input
	s, err := r.ReadString('\n')
	if err != nil {
		logger.Logger().Errorw("Unable to read user input", "error", err)
		return ""
	}

	return strings.TrimSuffix(s, "\n")
}

func readInt(in *os.File) int {
	var n int

	// No input verification as we assume proper input
	_, err := fmt.Fscanln(in, &n)
	if err != nil {
		// TODO
		return 5
	}

	return n
}
