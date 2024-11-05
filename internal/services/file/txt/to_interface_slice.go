package txt


func (s Service)ToInterfaceSlice(sliceString [][]string) [][]interface{} {
	//translate slice string to slice interface
	//return slice interface
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