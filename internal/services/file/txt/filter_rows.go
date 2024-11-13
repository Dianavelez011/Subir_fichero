package txt

func (s Service)FilterRows(rows [][]string)[][]string{
 for index,row := range rows{
	for _,field := range row{
		if field == "NOINCLUIDO"{
			rows = append(rows[:index],rows[index+1:]...)
		}
	}
 }

 return rows

}