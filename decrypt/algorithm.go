package decrypt

func Nimbus(str string) string {
	decrypted := ""
	for _,c := range str {
		asciicode := int(c)
		character := string(asciicode - 1)
		decrypted += character
	}
	return decrypted
}

