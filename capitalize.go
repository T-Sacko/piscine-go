// You can edit this code!
// Click here and start typing.
package piscine

func Capitalize(s string) string {

	str := []rune(s)

	if str[0] >= 'a' && str[0] <= 'z' {
		str[0] -= 32
	}

	for i := 1; i < len(str); i++ {

		pe := str[i-1]
		if pe >= 'a' && pe <= 'z' || pe >= 'A' && pe <= 'Z' || pe >= 48 && pe <= 57 {
			if str[i] >= 'A' && str[i] <= 'Z' {
				str[i] += 32
			}

		} else {
			if str[i] >= 'a' && str[i] <= 'z' {
				str[i] -= 32
			}

		}
	}

	return string(str)

}
