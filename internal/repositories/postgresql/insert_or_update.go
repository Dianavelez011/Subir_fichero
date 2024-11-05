package postgresql

import (
	"context"
	"fmt"
)

func (r Repository) InsertOrUpdate(query string, values []interface{}) error {
	//----------------------------------------------
	_, err := r.Connection.Exec(context.Background(), query, values...)
	if err != nil {
		return fmt.Errorf("failed to execute insert or update query: %w", err)
	}
	//Cerrar la conexion
	defer r.Connection.Close()

	fmt.Println("data inserted or updated successfully")
	return nil
}