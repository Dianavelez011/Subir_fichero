package excel

import (
	"fmt"
	"regexp"
	"strconv"
)

var(
	
	validateNumber = make(map[string]int)
	keysValidateNumber []string
)

func (s Service)ValidateChunk(fileName string,word string) (error){
	//get part file number
	//clean name


	re := regexp.MustCompile(fmt.Sprintf(`%s(\d+)`, word))
	matches := re.FindStringSubmatch(fileName)
	if len(matches) < 2 {
		return fmt.Errorf("could not found a number after '%s' in ValidateChunk", word)
	}

	numStr := matches[1] // Extraemos el número después de 'parte'

	number, err := strconv.Atoi(numStr)
	if err != nil {
		return fmt.Errorf("error converting the string to int: %s", err.Error())
	}

	//clean name
	cleanFileName := s.CleanName(fileName);
	valueValidateNumber,exists := validateNumber[cleanFileName]

	if !exists{
		validateNumber[cleanFileName] = 1
		valueValidateNumber = validateNumber[cleanFileName]
		keysValidateNumber =  append(keysValidateNumber,cleanFileName)
	}

	fmt.Printf("Numero: %d valueValidate: %d",number,valueValidateNumber)
	
	if  number != valueValidateNumber{
		return fmt.Errorf("file send sequence failed")
	}

	if len(keysValidateNumber) >= 5{
		for i:=0; i<2; i++{
			delete(validateNumber,keysValidateNumber[i])
		}
		keysValidateNumber = keysValidateNumber[2:]
	}

	validateNumber[cleanFileName] +=1 
	return nil

}