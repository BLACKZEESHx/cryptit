package encrypt

func Nimbus(str string) string {
	encrypted := ""
	for _, c := range str {
		ascii_code := int(c)
		character := ascii_code + 3
		encrypted += string(character)
	}
	encrypted = "Q" + encrypted + "A@1"
	return encrypted
}
