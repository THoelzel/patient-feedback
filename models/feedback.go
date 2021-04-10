package models

import (
	"context"
	"feedback/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Feedback struct {
	Rating              int
	DiagnosisUnderstood bool
	Improvement         string
	Feelings            string
	Diagnosis           string
}

var FeedbackTranslation = map[int]string{
	1:  "would definitely not recommend",
	2:  "would not recommend",
	3:  "would strongly hesitate to recommend",
	4:  "would hesitate to recommend",
	5:  "feel neutral about recommending",
	6:  "maybe will recommend",
	7:  "probably would recommend",
	8:  "would recommend",
	9:  "would not hesitate to recommend",
	10: "would strongly recommend",
}

type FeedbackModel struct {
	Pool *pgxpool.Pool
}

func (m *FeedbackModel) Insert(f *Feedback) {
	u := uuid.New()

	// explicitly store in utc
	utcTime := time.Now().UTC()
	_, err := m.Pool.Exec(context.Background(),
		"insert into zone1.feedback"+
			"(id, rating, diagnosis_understood, explanation_improvement, feelings, diagnosis, time_reported) "+
			"values ($1, $2, $3, $4, $5, $6, $7)",
		u, f.Rating, f.DiagnosisUnderstood, f.Improvement, f.Feelings, f.Diagnosis, utcTime)

	if err != nil {
		logger.Logger().Errorw("DB insert failed", "error", err)
	}
}

func (m *FeedbackModel) Stats() (s []Stats) {

	r, err := m.Pool.Query(context.Background(),
		"select round(avg(rating),3) as rating, round(avg(diagnosis_understood::int),3), diagnosis "+
			"from zone1.feedback group by diagnosis order by rating asc")

	if err != nil {
		logger.Logger().Errorw("Failed to retrieve stats", "error", err)
	}

	var sRows []Stats

	for r.Next() {
		sRow := Stats{}
		//fmt.Println(r.Values())
		errScan := r.Scan(&sRow.Rating, &sRow.Understandable, &sRow.Diagnosis)

		if errScan != nil {
			logger.Logger().Errorw("Failed to read stats", "error", errScan)
		}

		sRows = append(sRows, sRow)
	}

	return sRows
}
