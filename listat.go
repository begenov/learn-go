package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	i := l
	count := 0
	for i != nil {
		if pos == count {
			return i
		}
		count++
		i = i.Next
	}
	return nil
}
