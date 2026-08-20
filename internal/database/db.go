package db

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("record not found")

type DB struct {
	*sql.DB
}

func New(sqlDB *sql.DB) *DB {
	return &DB{sqlDB}
}


