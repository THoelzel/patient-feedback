package main

import (
	"github.com/spf13/cobra"
)

// root command
var feedbackCmd = &cobra.Command{
	Use:   "patient-feedback",
	Short: "Collect patient feedback",
	Run: func(cmd *cobra.Command, args []string) {
		collectData()
	},
}

var analysisCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze the collected patient feedback",
	Run: func(cmd *cobra.Command, args []string) {
		DisplayStats()
	},
}

func init() {
	feedbackCmd.AddCommand(analysisCmd)
}
