package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	connect(url string) (*pgxpool.Pool, error)
}

type ConnectionDb struct {
}

func (cd *ConnectionDb) NewConnection(db Database, url string) (*pgxpool.Pool, error) {
	dbpool, err := db.connect(url)
	if err != nil {
		return nil, fmt.Errorf("failed in db connections plis try again %w",err)
	}

	err = dbpool.Ping(context.Background())
	if err != nil{
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	return dbpool, nil
}

// func ConnectPool(route string)(*pgxpool.Pool, error){
//  dbpool,err := pgxpool.New(context.Background(), route)
//  if err != nil {
// 	return nil, fmt.Errorf("don't connect to data base: %w",err)
//  }
// //  defer dbpool.Close()

//  fmt.Println("Conexion establecida")
//  return dbpool,nil;
// }
