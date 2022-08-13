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
	if len(an) <= len(as) {
		an = append(an, af)
	}
	return an
}

// SplitWhiteSpaces("RSWj<hjS3)IiI Q7qMJtCUy\"LtF c38?%F5hf\\6Q| <Ki: %V/_C-ET 718SLx7ofj.]\" (=\"If7K8IRL26 \\[/Nv0_c/ER>\" 67^?oO_DwOAz ") == []string{"RSWj<hjS3)IiI", "Q7qMJtCUy\"LtF", "c38?%F5hf\\6Q|", "<Ki:", "%V/_C-ET", "718SLx7ofj.]\"", "(=\"If7K8IRL26", "\\[/Nv0_c/ER>\"", "67^?oO_DwOAz", ""} instead of []string{"RSWj<hjS3)IiI", "Q7qMJtCUy\"LtF", "c38?%F5hf\\6Q|", "<Ki:", "%V/_C-ET", "718SLx7ofj.]\"", "(=\"If7K8IRL26", "\\[/Nv0_c/ER>\"", "67^?oO_DwOAz"}
