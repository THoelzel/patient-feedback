package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func DisplayStats() {
	stats := env.feedback.Stats()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ICD-10 Diagnosis Code", "Average Rating", "Rate of Understanding"})

	for _, s := range stats {
		table.Append(
			[]string{s.Diagnosis, fmt.Sprintf("%.3f", s.Rating), fmt.Sprintf("%.3f", s.Understandable)},
		)
	}
	table.Render()
}
