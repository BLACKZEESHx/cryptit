package decrypt

import "strings"

func Nimbus(str string) string {
	decrypted := ""
	for _, c := range str {
		ascii_code := int(c)
		character := ascii_code - 3
		decrypted += string(character)
	}
	de := strings.Replace(decrypted, ">=.", "", -1)
	de = strings.Replace(de, "N", "", 1)
	return de
}
