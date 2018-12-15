package romanNumerals2number

// define a function for allowed substraction
// "I" can be subtracted from "V" and "X" only.
// "X" can be subtracted from "L" and "C" only.
// "C" can be subtracted from "D" and "M" only.
// "V", "L", and "D" can never be subtracted.
func allowedSubstraction(symbol string) string{
	_map := map[string]string{
		"V": "I",
		"X": "I",
		"L": "X",
		"C": "X",
		"D": "C",
		"M": "C",
	}
	return _map[symbol]
}