package piscine

func Capitalize(s string) string {
	str := []byte(tolower(s))
	ass := true
	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' && ass {
			str[i] -= 32
			ass = false
		} else if (str[i] < 97 || str[i] > 122) && (str[i] < 48 || str[i] > 57) {
			ass = true
		} else {
			ass = false
		}
	}

	return string(str)
}

func tolower(s string) string {
	new := []byte(s)
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			new[i] = s[i] - 'A' + 'a'
		}
	}
	return string(new)
}
