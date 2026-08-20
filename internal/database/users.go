package db

import (
	"database/sql"
	"errors"
	"time"

	"github.com/aneeshsunganahalli/Kumo/internal/models"
)

func (db *DB) CreateUser(u *models.Users) error {
	row := db.QueryRow(
		`INSERT INTO users (username, email) VALUES (?, ?) RETURNING id, created_at`,
		u.Username, u.Email,
	)
	return row.Scan(&u.ID, &u.CreatedAt)
}

func (db *DB) GetUser(id int64) (*models.Users, error) {
	u := &models.Users{}
	err := db.QueryRow(
		`SELECT id, username, email, created_at FROM users WHERE id = ?`, id,
	).Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return u, err
}

func (db *DB) UpdateUser(id int64, u *models.Users) error {
	res, err := db.Exec(
		`UPDATE users SET username = ?, email = ? WHERE id = ?`,
		u.Username, u.Email, id,
	)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	u.ID = id
	return nil
}

func (db *DB) CreatePreferences(p *models.UIPreference) error {
	row := db.QueryRow(
		`INSERT INTO ui_preferences (user_id, font_family, theme_mode, fallback_color, active_background_id, active_audio_id)
		 VALUES (?, ?, ?, ?, ?, ?)
		 RETURNING updated_at`,
		p.UserID, p.FontFamily, p.ThemeMode, p.FallbackColor, p.ActiveBackgroundID, p.ActiveAudioID,
	)
	return row.Scan(&p.UpdatedAt)
}

func (db *DB) GetPreferences(userID int64) (*models.UIPreference, error) {
	p := &models.UIPreference{}
	err := db.QueryRow(
		`SELECT user_id, font_family, theme_mode, fallback_color, active_background_id, active_audio_id, updated_at
		 FROM ui_preferences WHERE user_id = ?`, userID,
	).Scan(&p.UserID, &p.FontFamily, &p.ThemeMode, &p.FallbackColor,
		&p.ActiveBackgroundID, &p.ActiveAudioID, &p.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return p, err
}

func (db *DB) UpdatePreferences(userID int64, p *models.UIPreference) error {
	p.UpdatedAt = time.Now()
	res, err := db.Exec(
		`UPDATE ui_preferences
		 SET font_family = ?, theme_mode = ?, fallback_color = ?,
		     active_background_id = ?, active_audio_id = ?, updated_at = ?
		 WHERE user_id = ?`,
		p.FontFamily, p.ThemeMode, p.FallbackColor,
		p.ActiveBackgroundID, p.ActiveAudioID, p.UpdatedAt,
		userID,
	)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	p.UserID = userID
	return nil
}
