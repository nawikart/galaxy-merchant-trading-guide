package utils

func InArray(x string, a []string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}