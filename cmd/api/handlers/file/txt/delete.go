package txt

func (h Handler)Delete(mainFilePath string)error{

	err := h.FileService.Delete(mainFilePath);

	if err != nil{
		return err
	}

	return nil

}