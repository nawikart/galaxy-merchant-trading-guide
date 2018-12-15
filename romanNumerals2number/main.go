package romanNumerals2number

import (
		"strings"
		"../utils"
		)


func Convert(symbols string) (interface{}, bool){
	symbols = strings.ToUpper(strings.Replace(symbols, " " , "", -1))

	var prevSymbol, repeatedSymbol string
	var prevRoman, prevprevRoman, repeatedSymbolCount, sum int64
	var valid, hvSubtraction bool

	dlv := []string{"D", "L", "V"}
	allowedSymbols := []string{"I", "V", "X", "L", "C", "D", "M"}	
	sum = 0
	prevRoman = 0
	prevprevRoman = 0
	repeatedSymbolCount = 1
	valid = true
	for _, e := range symbols {

		if !utils.InArray(string(e), allowedSymbols) {
			valid = false
		}

		if repeatedSymbol == string(e){
			repeatedSymbolCount ++			
		}else{
			repeatedSymbol = string(e)
			repeatedSymbolCount = 1
		}

		if repeatedSymbolCount > 1 && utils.InArray(string(e), dlv)  {
			valid = false
		}else  if repeatedSymbolCount > 3 {
			valid = false
		}

		if prevRoman > 0 && prevRoman < romanS2number(string(e)) {
			if prevprevRoman > 0 && prevprevRoman < romanS2number(string(e)){
				valid = false
			}
			if hvSubtraction {
				valid = false
			}else{
				hvSubtraction = true
			}
				
			if allowedSubstraction(string(e)) != prevSymbol {
				valid = false
			}

			sum = sum + (romanS2number(string(e)) - prevRoman - prevRoman)
		}else{
			hvSubtraction = false
			sum = sum + romanS2number(string(e))
		}
		prevprevRoman = prevRoman
		prevRoman = romanS2number(string(e))
		prevSymbol = string(e)
	}

	if !valid {
		return nil, valid
	}else{
		return sum, valid
	}
}