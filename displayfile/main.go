package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 0 {
		fmt.Println("File name missing")
	} else if len(arg) > 1 {
		fmt.Println("Too many arguments")
	} else {
		file, _ := ioutil.ReadFile(arg[0])
		fmt.Println(string(file))
	}
}
