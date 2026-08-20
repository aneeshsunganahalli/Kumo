package db

import (
	"database/sql"
	"errors"

	"github.com/aneeshsunganahalli/Kumo/internal/models"
)

func (db *DB) CreateAudio(a *models.Audio) error {
	row := db.QueryRow(
		`INSERT INTO audio_tracks (user_id, original_name, stored_filename, file_type)
		 VALUES (?, ?, ?, ?) RETURNING id, created_at`,
		a.UserID, a.OriginalName, a.StoredFileName, a.FileType,
	)
	return row.Scan(&a.ID, &a.CreatedAt)
}

func (db *DB) GetAudio(id int64) (*models.Audio, error) {
	a := &models.Audio{}
	err := db.QueryRow(
		`SELECT id, user_id, original_name, stored_filename, file_type, created_at
		 FROM audio_tracks WHERE id = ?`, id,
	).Scan(&a.ID, &a.UserID, &a.OriginalName, &a.StoredFileName, &a.FileType, &a.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return a, err
}

func (db *DB) GetAudiosByUser(userID int64) ([]models.Audio, error) {
	rows, err := db.Query(
		`SELECT id, user_id, original_name, stored_filename, file_type, created_at
		 FROM audio_tracks WHERE user_id = ? ORDER BY created_at DESC`, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Audio
	for rows.Next() {
		var a models.Audio
		if err := rows.Scan(&a.ID, &a.UserID, &a.OriginalName, &a.StoredFileName, &a.FileType, &a.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, rows.Err()
}

func (db *DB) UpdateAudio(id int64, a *models.Audio) error {
	res, err := db.Exec(
		`UPDATE audio_tracks SET original_name = ?, stored_filename = ?, file_type = ? WHERE id = ?`,
		a.OriginalName, a.StoredFileName, a.FileType, id,
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
	a.ID = id
	return nil
}

func (db *DB) DeleteAudio(id int64) error {
	res, err := db.Exec(`DELETE FROM audio_tracks WHERE id = ?`, id)
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
	return nil
}