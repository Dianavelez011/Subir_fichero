package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Querys struct {
	ConnectionDb
	Url      string
	Database PostgreSql
}

func (q *Querys) constructor() *Querys {
	q.Url = "postgres://admin:admin123@localhost:5434/db_buscador"
	q.Database = PostgreSql{}
	return q
}

func toInterface(sliceString [][]string) [][]interface{} {
	var convertedInterface [][]interface{}
	for _, row := range sliceString {
		var interfaceRow []interface{}
		for _, col := range row {
			interfaceRow = append(interfaceRow, col)
		}
		convertedInterface = append(convertedInterface, interfaceRow)
	}

	return convertedInterface

}

func (q Querys) InsertOrUpdate(query string, values []interface{}) error {
	dbpool, err := q.NewConnection(q.constructor().Database, q.constructor().Url)
	if err != nil {
		return err
	}

	_, err = dbpool.Exec(context.Background(), query, values...)
	if err != nil {
		return fmt.Errorf("failed to execute insert or update query: %w", err)
	}

	fmt.Println("data inserted or updated successfully")
	return nil
}

func (q Querys) SelectAll(query string)(pgx.Rows,error){
	dbpool, err := q.NewConnection(q.constructor().Database, q.constructor().Url)
	if err != nil {
		return nil,err
	}

	rows, err := dbpool.Query(context.Background(),query)
	if err != nil{
		return nil, fmt.Errorf("query failed in SelectAll: %w",err)
	}

	return rows, nil
}

func (q Querys) Select(query string,value interface{})(pgx.Row,error){
	dbpool, err := q.NewConnection(q.constructor().Database, q.constructor().Url)
	if err != nil {
		return nil,err
	}
	row := dbpool.QueryRow(context.Background(),query,value)
	return row,nil
}

func (q Querys) CopyFrom(columns []string, values [][]string)error {
	dbpool, err := q.NewConnection(q.constructor().Database, q.constructor().Url)
	if err != nil {
		return err
	}

	//Iniciar una transacción
	tx, err := dbpool.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w",err)
	}
	//Asegura la reversion en caso de error
	defer tx.Rollback(context.Background())
	//Cerrar la conexion de la base de datos
	defer dbpool.Close()

	rowsInterface := toInterface(values)

	_, err = tx.CopyFrom(context.Background(), pgx.Identifier{"ruv_victimas"}, columns, pgx.CopyFromRows(rowsInterface))

	if err != nil {
		tx.Rollback(context.Background())
		return fmt.Errorf("copyFrom failed: %w", err)
	}

	// Confirmar la transacción si no hubo errores
	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	fmt.Println("Rows copied successfully")
	return nil
}
