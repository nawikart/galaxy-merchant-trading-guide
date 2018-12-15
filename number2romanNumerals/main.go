package number2romanNumerals

import (
		"../utils"
		"strings"
		"strconv"
		)


func Convert(numberStr string) (interface{}, bool){
	valid := true
	romanN := "" // roman numberals
	prevRomanS := "" // previous roman symbol
	prevprevRomanS := "" // previous previous roman symbol

	numberStr = strings.Replace(numberStr, " " , "", -1) // remove white space from the string of number
	var number int64 // variable to relocate a number convertion from string of number
	i, err := strconv.Atoi(numberStr) // convert string of number to a real number
	if err != nil {
		valid = false
	}else{
		number = int64(i)
	}
		
	for _, NnRS := range ArrayNnRomanS() { // 

		dlv := []string{"D", "L", "V"} // define array for roman symbol which can never be repeated and can never be subtracted.

		if !utils.InArray(NnRS.RomanS, dlv) { // repeat only for roman symbol other that "D" "L" "V"
			divi := int(number / NnRS.Number) // relocate the quotient results between the "number" and "1000" "100" "10" to "divi" variable
			number = number % NnRS.Number // relocate the quotient remainder results between the "number" and "1000" "100" "10" back to "number" variable 

			// The symbols "I", "X", "C", and "M" can be repeated three times in succession, but no more.
			// it will be Invalid if "divi" for symbol "M" is more than 3
			if NnRS.RomanS == "M" && divi > 3 {
				valid = false
			}
			if divi == 9 {
				romanN = romanN + NnRS.RomanS + prevprevRomanS

			}else if divi >= 6 && divi <= 8 {
				romanN = romanN + prevRomanS
				for i := 0; i < (divi - 5); i++ {
					romanN = romanN + NnRS.RomanS
				}				

			}else if divi == 4 {
				romanN = romanN + NnRS.RomanS + prevRomanS

			}else{
				for i := 0; i < divi; i++ {
					romanN = romanN + NnRS.RomanS
				}
			}
		}

		if number == 0 {
			break
		}
		
		prevprevRomanS = prevRomanS
		prevRomanS = NnRS.RomanS
	}	

	if !valid {
		return nil, false
	}else{
		return romanN, true
	}
}