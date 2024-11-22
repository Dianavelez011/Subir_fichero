package txt

import "fmt"

func (h Handler) SaveMainTable() error{
	err := h.FileService.SaveMainTable()

	if err != nil {
		return fmt.Errorf("error in handler SaveMainTable: %s",err.Error())
	}
	return nil
}