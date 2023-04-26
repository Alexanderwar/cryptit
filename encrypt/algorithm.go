package encrypt

func Nimbus(str string) string {
	encrypted := ""
	for _,c := range str {
		asciicode := int(c)
		character := string(asciicode + 1)
		encrypted += character
	}
	return encrypted
}

