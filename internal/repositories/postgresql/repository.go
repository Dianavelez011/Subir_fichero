package postgresql

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	Connection *pgxpool.Pool
}