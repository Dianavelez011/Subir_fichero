package connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(dbUri string) (*pgxpool.Pool, error) {
	connection, err := pgxpool.New(context.Background(), dbUri)

	if err != nil {
		return nil, fmt.Errorf("failed in connection database: %w", err)
	}

	err = connection.Ping(context.Background())
	if err != nil{
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}


	return connection,nil
}
