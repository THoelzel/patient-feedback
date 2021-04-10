package main

import (
	"context"
	"feedback/input"
	"feedback/logger"
	"feedback/models"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

type Env struct {
	feedback interface {
		Insert(f *models.Feedback)
		Stats() []models.Stats
	}
}

var env *Env

func main() {
	pool := startPool()

	env = &Env{
		feedback: &models.FeedbackModel{Pool: pool},
	}

	// by default we ask for patient feedback
	err := feedbackCmd.Execute()
	if err != nil {
		logger.Logger().Errorw("Failed to start root command", "error", err)
		return
	}
}

func collectData() {
	file := &input.FileInput{FileName: "patient-feedback-raw-data.json"}
	event := file.NextEvent()
	feedback := UserInput(event, os.Stdin)

	env.feedback.Insert(feedback)
}

func startPool() *pgxpool.Pool {

	c := loadConfig()

	connString := "postgresql://%s:%s@%s:%d/%s"

	pool, err := pgxpool.Connect(context.Background(), fmt.Sprintf(connString, c.User, c.Pass, c.Host, c.Port, c.DbName))

	if err != nil {
		logger.Logger().Errorw("Unable to connect to database", "error", err)
		os.Exit(1)
	}

	return pool
}
