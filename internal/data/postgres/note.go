package postgres

import "database/sql"

type NoteStore struct {
	DB *sql.DB
}
