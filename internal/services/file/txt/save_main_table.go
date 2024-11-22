package txt

import (
	"fmt"
	"strings"
)

func (s Service) SaveMainTable() error{
	//create query
	//execute query
	//return error
	// columns := strings.Join(s.Columns,",")
	columns := `"` + strings.Join(s.Columns, `", "`) + `"`
	setString := ""
	valuesString := ""
	for _, value := range s.Columns {
		// setString += `tgt."` + value + `"= src."` + value + `", `
		setString += `"`+value+`"= src."` + value + `", `
		valuesString += `src."` + value + `", `
	}
	//Delete last , in setString
	setString = strings.TrimSuffix(setString,", ")
	valuesString = strings.TrimSuffix(valuesString,", ")

	query := `MERGE INTO public.` + s.TableName + ` AS tgt 
	USING ` + s.TemporaryTable + ` AS src 
	ON (` + s.Querys["condition_merge"] + `) WHEN MATCHED THEN UPDATE SET ` +setString + ` 
	WHEN NOT MATCHED THEN INSERT(` + columns + `) VALUES(`+ valuesString + `);`

	fmt.Println(query)

	if err := s.Repo.Upsert(query); err != nil{
		return fmt.Errorf("error in SaveMainTable: %s",err.Error())
	}
	return nil
}
