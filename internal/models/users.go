package models

import "time"

type Users struct {
	ID           int64     `db:"id" json:"id"`
	Username     string    `db:"username" json:"username"`
	Email        string    `db:"email" json:"email"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type UIPreference struct {
	UserID             int64     `db:"user_id" json:"user_id"`
	FontFamily         string    `db:"font_family" json:"font_family"`
	ThemeMode          string    `db:"theme_mode" json:"theme_mode"`
	FallbackColor      string    `db:"fallback_color" json:"fallback_color"`
	ActiveBackgroundID *int64    `db:"active_background_id" json:"active_background_id"` // Pointer because it can be NULL
	ActiveAudioID      *int64    `db:"active_audio_id" json:"active_audio_id"`
	UpdatedAt          time.Time `db:"updated_at" json:"updated_at"`
}

