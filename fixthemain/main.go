package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	return true
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?")
	return true
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return true
}

func main() {
	door := &Door{}

	if IsDoorOpen(door) {
		IsDoorClose(door)
	}
	if IsDoorClose(door) {
		IsDoorClose(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
}

type Door struct {
	state string
}
