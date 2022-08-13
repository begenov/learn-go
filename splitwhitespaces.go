package piscine

func SplitWhiteSpaces(s string) []string {
	as := []rune(s)
	var an []string
	af := ""
	b := false
	for i := 0; i < len(as); i++ {
		if as[i] == ' ' && b {
			an = append(an, af)
			af = ""
			b = false
		}

		if as[i] != ' ' {
			af += string(as[i])
			b = true
		}
	}

	an = append(an, af)
	return an
}
