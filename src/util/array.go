package util

func Contains(s []string, e string) int {
	for idx, a := range s {
		if a == e {
			return idx
		}
	}
	return -1
}
