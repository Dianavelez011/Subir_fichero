package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r Repository) CopyFrom(columns []string, values [][]interface{},tableName string) error {
	

	//Iniciar una transacción
	trans, err := r.Connection.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	//Cerrar la conexion 
	defer r.Connection.Close()
	

	//Copiar datos a la base de datos
	_, err = trans.CopyFrom(context.Background(), pgx.Identifier{tableName}, columns, pgx.CopyFromRows(values))

	if err != nil {
		trans.Rollback(context.Background())
		return fmt.Errorf("copyFrom failed: %w", err)
	}

	// Confirmar la transacción si no hubo errores
	if err := trans.Commit(context.Background()); err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	fmt.Println("Rows copied successfully")
	return nil
}
