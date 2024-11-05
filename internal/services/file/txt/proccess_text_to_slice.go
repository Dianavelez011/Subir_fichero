package txt

var incompleteArrays [][]string

// fillIncompleteArray llena un arreglo incompleto
func (s Service)FillIncompleteSlice(rows []string) []string {
	quantity := s.TabRow - len(incompleteArrays[0])

	for i := 0; i < quantity; i++ {
		incompleteArrays[0] = append(incompleteArrays[0], rows[i])
		if i+1 == quantity {
			break
		}
	}

	return incompleteArrays[0]
}


func (s Service)ProccessTextToSlice(text string) [][]string{
	//convert text to slice
	//filter slice
	//fild incomplete slice
	//convert slice to double slice

	// Convert text to slice
	rows := s.convertTextToSlice(text)
	// (text, "»", "NULL")
	incompleteArrays = [][]string{}

	// Filter rows
	filterRows := s.FilterFields(rows)
	// (rows, []string{
	// 	addColumn,
	// 	"NULL",
	// }, 53)


	// Initial double array
	doubleArray := [][]string{}
	offset := 0

	for i := 0; i < len(filterRows); i += s.TabRow {
		if len(incompleteArrays) > 0 {
			arrayComplete := s.FillIncompleteSlice(filterRows)
			// Agrega el array completo a doubleArray
			doubleArray = append(doubleArray, arrayComplete)
			incompleteArrays = [][]string{} // Vacía el array de incompletos
		}

		col := filterRows[i:min(i+s.TabRow, len(filterRows))] // Obtener un sub-slice

		if len(col) > s.TabRow {
			col = col[:s.TabRow] // Limita la fila a 53 elementos
		}else if len(col) == s.TabRow {
			doubleArray = append(doubleArray, col)
		} else {
			incompleteArrays = append(incompleteArrays, col)
		}

		offset += s.TabRow
	}
	return doubleArray
}