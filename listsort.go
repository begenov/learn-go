package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSort(l *NodeI) *NodeI {
	cmpt := 0
	var first *NodeI

	if l == nil || l.Next == nil {
		return l
	}

	for l != nil {
		next := l.Next
		if cmpt == 0 {
			first = l
			cmpt++
		}

		for next != nil {
			if l.Data > next.Data {
				l.Data, next.Data = next.Data, l.Data
			}
			next = next.Next
		}
		l = l.Next
	}
	return first
}
