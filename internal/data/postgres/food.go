package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/psodm/macroeats/internal/data"
)

type FoodStore struct {
	conn *pgx.Conn
}

func NewFoodStore(conn *pgx.Conn) *FoodStore {
	return &FoodStore{
		conn: conn,
	}
}

func (f *FoodStore) Insert(ctx context.Context, food *data.Food) error {
	return nil
}
