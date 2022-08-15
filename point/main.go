package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {
	points := &point{}
	setPoint(points)
	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}
