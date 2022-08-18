package main

import (
	"fmt"
)

func main() {
	x := 5
	y := &x
	z := &y
	a := &z

	w := 2
	b := &w

	u := 7
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j

	k := 6
	l := &k
	m := &l
	n := &m
	d := &n

	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

	Enigma(a, b, c, d)

	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
}

func Enigma(a ***int, b *int, c *******int, d ****int) {
	a1 := ***a
	b1 := *b
	c1 := *******c
	d1 := ****d
	***a = c1
	*******c = d1
	****d = b1
	*b = a1
}
