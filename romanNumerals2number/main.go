package romanNumerals2number

import (
		"strings"
		"../utils"
		)


// a function to convert roman numerals to number
func Convert(symbols string) (interface{}, bool){
	// make the symbols uppercase and remove all white space
	symbols = strings.ToUpper(strings.Replace(symbols, " " , "", -1))

	var prevSymbol, repeatedSymbol string
	var prevNumber, prevprevNumber, repeatedSymbolCount, sum int64
	var valid, hvSubtraction bool

	// define array for roman symbol which can never be repeated and can never be subtracted.
	dlv := []string{"D", "L", "V"}

	// define array for allowed roman symbols 
	allowedSymbols := []string{"I", "V", "X", "L", "C", "D", "M"}	
	sum = 0
	prevNumber = 0
	prevprevNumber = 0
	repeatedSymbolCount = 1
	valid = true

	// looping the characters of symbols, eg. MMXIX
	for _, e := range symbols { // value of "e" will be "M" "M" "X" "I" "X"

		// check if each character of symbols is correct base on allowed roman symbols
		if !utils.InArray(string(e), allowedSymbols) {
			valid = false
		}

		// counting the repeatedSymbol
		if repeatedSymbol == string(e){
			repeatedSymbolCount ++			
		}else{
			repeatedSymbol = string(e)
			repeatedSymbolCount = 1
		}


		// check if the symbol repeated 
		// and if it is including one of the symbol that is not allowed to be repeated then it is invalid
		if repeatedSymbolCount > 1 && utils.InArray(string(e), dlv)  {
			valid = false

		// or any symbol repeated more than 3, it is invalid
		}else  if repeatedSymbolCount > 3 {
			valid = false
		}

		// if previous number of symbol > 0, mean this current symbol is not the first character
		// and previous number lower than current number of symbol
		if prevNumber > 0 && prevNumber < romanS2number(string(e)) {

			// if the previous two symbols is lower than currect symbol, it is invalid
			if prevprevNumber > 0 && prevprevNumber < romanS2number(string(e)){
				valid = false
			}

			// if previously also have a substraction, it is invalid
			if hvSubtraction {
				valid = false
			}else{

				// note this is have substraction
				hvSubtraction = true
			}
			
			
			if allowedSubstraction(string(e)) != prevSymbol {
				valid = false
			}

			sum = sum + (romanS2number(string(e)) - prevNumber - prevNumber)
		}else{
			hvSubtraction = false
			sum = sum + romanS2number(string(e))
		}
		prevprevNumber = prevNumber
		prevNumber = romanS2number(string(e))
		prevSymbol = string(e)
	}

	if !valid {
		return nil, valid
	}else{
		return sum, valid
	}
}