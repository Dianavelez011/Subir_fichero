package postgresql

import (
	"context"
	"fmt"
)

func (r Repository) Upsert(query string) error{
	defer r.Connection.Close()
	_, err := r.Connection.Exec(context.Background(),query)
	if err != nil{
		return fmt.Errorf("could not upsert values in Upsert: %s",err.Error())
	}
	return nil
}