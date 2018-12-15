package number2romanNumerals

// define struct of pairs Number and Roman Symbol
type NnRomanS struct{
	Number int64
	RomanS string // roman symbol
}

// a function to define array of pairs Number and Roman Symbol
func ArrayNnRomanS() []NnRomanS{
	var arr []NnRomanS
	arr = append(arr, NnRomanS{Number: 1000, RomanS: "M"})
	arr = append(arr, NnRomanS{Number: 500, RomanS: "D"})
	arr = append(arr, NnRomanS{Number: 100, RomanS: "C"})
	arr = append(arr, NnRomanS{Number: 50, RomanS: "L"})
	arr = append(arr, NnRomanS{Number: 10, RomanS: "X"})
	arr = append(arr, NnRomanS{Number: 5, RomanS: "V"})
	arr = append(arr, NnRomanS{Number: 1, RomanS: "I"})
	return arr
}