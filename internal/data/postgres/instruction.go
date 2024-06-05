package postgres

import "database/sql"

type InstructionStore struct {
	DB *sql.DB
}
