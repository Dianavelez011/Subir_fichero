package txt

func (s Service)convertTextToSlice(text string) []string {
	array := make([]string, 0)
	counterCharacter := 0
	newWord := ""
	targetChar := []rune(s.Character)[0] // Toma el primer carácter de la cadena.

	for _, char := range text {
		if char == targetChar {
			counterCharacter++
			continue
		}

		switch {
		case counterCharacter == 1:
			array = append(array, newWord)
			newWord = ""
			counterCharacter = 0
		case counterCharacter > 1:
			array = append(array, newWord)
			for i := 0; i < counterCharacter-1; i++ {
				array = append(array, s.DefaultValue)
			}
			newWord = ""
			counterCharacter = 0
		}

		newWord += string(char)
	}

	// Asegurarse de agregar la última palabra si existe
	if newWord != "" {
		array = append(array, newWord)
	}

	return array
}