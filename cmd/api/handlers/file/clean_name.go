package file

func (h Handler)CleanName(fileName string)string{
	cleanName := h.TxtService.CleanName(fileName)
	return cleanName
}