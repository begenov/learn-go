package main

import (
	"github.com/01-edu/z01"
)

type Door struct {
	OPEN  string
	CLOSE string
}

func main() {
	door := Door{}

	OPENDOOR(door)
	if IsDoorClose(door) {
		OPENDOOR(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
}

func OPENDOOR(door Door) {
	door.OPEN = "Door Opening..."
	for _, j := range door.OPEN {
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}

func IsDoorClose(door Door) bool {
	PrintStr("is the Door closed ?")
	return false
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func IsDoorOpen(Door Door) bool {
	PrintStr("is the Door opened ?")
	return true
}

func CloseDoor(door Door) {
	PrintStr("Door Closing...")
}
