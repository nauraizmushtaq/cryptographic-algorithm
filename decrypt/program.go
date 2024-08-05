package decrypt

func NimbusDecrypt(str string) string {
	decryptStr := ""

	for _, c := range str {
		character := int(c)
		decryptStr += string(character - 3)
	}
	return decryptStr
}
