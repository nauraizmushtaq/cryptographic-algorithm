package encrypt

func NimbusEnrypt(str string) string { //Capatical function name start export the func outside of the package, otherwise the pkg will remain
	// within available within the package.
	encryptedStr := ""
	for _, c := range str {
		asciiCode := int(c)
		characterAscii := string(asciiCode + 3)
		encryptedStr += characterAscii
	}
	return encryptedStr
}
