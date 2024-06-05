package postgres

import "database/sql"

type IngredientSectionStore struct {
	DB *sql.DB
}

type IngredientStore struct {
	DB *sql.DB
}
