package postgresql

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
)

func (r Repository) CopyFrom(columns []string, values [][]interface{},tableName string,channel chan map[string]interface{},wg *sync.WaitGroup)  {
	
	response := map[string]interface{}{
		"error": nil,
		"db_connection": r.Connection,
	}
	defer wg.Done()
	//Iniciar una transacción
	trans, err := r.Connection.Begin(context.Background())
	if err != nil {
		response["error"] = fmt.Errorf("failed to begin transaction: %w", err)
		channel <- response
		return
	}

	//Cerrar la conexion 
	// defer r.Connection.Close()
	

	//Copiar datos a la base de datos
	_, err = trans.CopyFrom(context.Background(), pgx.Identifier{tableName}, columns, pgx.CopyFromRows(values))

	if err != nil {
		trans.Rollback(context.Background())
		response["error"] = fmt.Errorf("copyFrom failed: %w", err)
		channel <- response
		return
	}

	// Confirmar la transacción si no hubo errores
	if err := trans.Commit(context.Background()); err != nil {
		response["error"] = fmt.Errorf("transaction commit failed: %w", err)
		channel <- response
		return
	}

	fmt.Println("rows copied succesfully in CopyFrom")
	channel <- response

}
