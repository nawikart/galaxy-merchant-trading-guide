package utils

// a function to check if a string is in an array of string
func InArray(x string, a []string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}