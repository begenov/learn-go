package piscine

func ConcatParams(args []string) string {
	str := ""
	for i := 0; i < len(args); i++ {
		for _, v := range args[i] {
			str += string(v)
		}
		if i != len(args)-1 {
			str += string('\n')
		}
	}
	return str
}
