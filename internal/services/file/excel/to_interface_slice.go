package excel

func (s Service)ToInterfaceSlice(sliceString [][]string) [][]interface{} {
	//fill rows incomplets
	//translate slice string to slice interface
	//return slice interface
	
	var convertedInterface [][]interface{}
	for _,row := range sliceString {
		if len(row) == (s.TabRow - 1) {
				row = append(row, "NULL")
		}

		var interfaceRow []interface{}
		for _, col := range row {
			interfaceRow = append(interfaceRow, col)
		}
		convertedInterface = append(convertedInterface, interfaceRow)
	}

	return convertedInterface

}