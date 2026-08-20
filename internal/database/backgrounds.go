package db

import (
	"database/sql"
	"errors"

	"github.com/aneeshsunganahalli/Kumo/internal/models"
)

func (db *DB) CreateBackground(b *models.Background) error {
	row := db.QueryRow(
		`INSERT INTO backgrounds (user_id, original_name, stored_filename, file_type)
		 VALUES (?, ?, ?, ?) RETURNING id, created_at`,
		b.UserID, b.OriginalName, b.StoredFileName, b.FileType,
	)
	return row.Scan(&b.ID, &b.CreatedAt)
}

func (db *DB) GetBackground(id int64) (*models.Background, error) {
	b := &models.Background{}
	err := db.QueryRow(
		`SELECT id, user_id, original_name, stored_filename, file_type, created_at
		 FROM backgrounds WHERE id = ?`, id,
	).Scan(&b.ID, &b.UserID, &b.OriginalName, &b.StoredFileName, &b.FileType, &b.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return b, err
}

func (db *DB) GetBackgroundsByUser(userID int64) ([]models.Background, error) {
	rows, err := db.Query(
		`SELECT id, user_id, original_name, stored_filename, file_type, created_at
		 FROM backgrounds WHERE user_id = ? ORDER BY created_at DESC`, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Background
	for rows.Next() {
		var b models.Background
		if err := rows.Scan(&b.ID, &b.UserID, &b.OriginalName, &b.StoredFileName, &b.FileType, &b.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, b)
	}
	return list, rows.Err()
}

func (db *DB) UpdateBackground(id int64, b *models.Background) error {
	res, err := db.Exec(
		`UPDATE backgrounds SET original_name = ?, stored_filename = ?, file_type = ? WHERE id = ?`,
		b.OriginalName, b.StoredFileName, b.FileType, id,
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
	b.ID = id
	return nil
}

func (db *DB) DeleteBackground(id int64) error {
	res, err := db.Exec(`DELETE FROM backgrounds WHERE id = ?`, id)
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