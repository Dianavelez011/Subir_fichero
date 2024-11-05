package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r Repository) SelectAll(query string) (pgx.Rows, error) {
	rows, err := r.Connection.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("query failed in SelectAll: %w", err)
	}
	//Cerrar conexion
	defer r.Connection.Close()

	return rows, nil
}