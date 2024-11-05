package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (r Repository) Select(query string, value interface{}) (pgx.Row, error) {

	row := r.Connection.QueryRow(context.Background(), query, value)
	defer r.Connection.Close()
	return row, nil
}