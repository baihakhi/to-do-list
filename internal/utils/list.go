package utils

// IsContain check if there is value in listStr,
// return true if value match string in listStr.
func IsContain(value string, listStr []string) bool {
	for _, v := range listStr {
		if v == value {
			return true
		}
	}
	return false
}
