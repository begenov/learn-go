package piscine

func ToUpper(s string) string {
	new := []byte(s)
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			new[i] = s[i] - 'a' + 'A'
		}
	}
	return string(new)
}
