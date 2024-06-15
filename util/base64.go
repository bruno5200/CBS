package util

// adds padding to a base64 string
func AddPaddingToBase64(input string) string {
	padLen := len(input) % 4
	if padLen == 2 {
		return input + "=="
	} else if padLen == 3 {
		return input + "="
	}
	return input
}
