package piscine

func ToLower(s string) string {
	new := []byte(s)
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			new[i] = s[i] - 'A' + 'a'
		}
	}
	return string(new)
}
