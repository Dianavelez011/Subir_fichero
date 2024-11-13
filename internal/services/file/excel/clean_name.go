package excel

import (
	"strings"
	"unicode"
)

func (s Service)CleanName(fileName string) string{
	//delete numbers in file
	return strings.Map(func(r rune) rune {
		// Retorna -1 para eliminar el carácter si es un dígito
		if unicode.IsDigit(r) {
			return -1
		}
		// De lo contrario, mantiene el carácter
		return r
	}, fileName)
}