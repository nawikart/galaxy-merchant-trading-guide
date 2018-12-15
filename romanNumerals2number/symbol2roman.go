package romanNumerals2number

// define a function "roman symbol to number" to assign a arabic number for a roman symbol
func romanS2number(symbol string) int64{
	_map := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	return int64(_map[symbol])
}