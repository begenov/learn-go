package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	for i := range arg {
		if arg[i] == "galaxy" || arg[i] == "galaxy 01" || arg[i] == "01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
