package models

import "time"

type Background struct {
	ID             int64     `db:"id" json:"id"`
	UserID         int64     `db:"user_id" json:"user_id"`
	OriginalName   string    `db:"original_name" json:"original_name"`
	StoredFileName string    `db:"stored_filename" json:"stored_filename"` 
	FileType       string    `db:"file_type" json:"file_type"`             
	CreatedAt      time.Time `db:"created_at" json:"created_at"`           
}


type Audio struct {
	ID             int64     `db:"id" json:"id"`
	UserID         int64     `db:"user_id" json:"user_id"`
	OriginalName   string    `db:"original_name" json:"original_name"`
	StoredFileName string    `db:"stored_filename" json:"stored_filename"`
	FileType       string    `db:"file_type" json:"file_type"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}


type Quote struct {
	ID        int64     `db:"id" json:"id"`
	Content   string    `db:"content" json:"content"`
	Author    string    `db:"author" json:"author"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

