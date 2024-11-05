package file

func ToInterfaceSlice(sliceString [][]string) [][]interface{} {
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