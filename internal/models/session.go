package models

import "time"

type Session struct {
	ID                int64     `db:"id" json:"id"`
	UserID            int64     `db:"user_id" json:"user_id"`
	SessionDate       time.Time `db:"session_date" json:"session_date"`
	TotalStudySeconds int       `db:"total_study_seconds" json:"total_study_seconds"`
	Notes             *string   `db:"notes" json:"notes"` 
}