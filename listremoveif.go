package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil {
		return
	}
	c := l.Head
	var previous *NodeL

	for c != nil {
		if c.Data == data_ref {
			if previous == nil {
				l.Head = c.Next
			} else {
				previous.Next = c.Next
				c = previous
			}
		} else {
			previous = c
		}
		c = c.Next
	}
}
