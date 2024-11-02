package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgreSql struct {
}



func (p PostgreSql) connect(url string) (*pgxpool.Pool,error){
	dbpool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("don't connect to data base: %w", err)
	}
	//  defer dbpool.Close()
	fmt.Println("Conexion establecida")
	return dbpool, nil
}
