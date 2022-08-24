package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("len", len(os.Args))

	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
}

//go run .\main.go hello I ma rashmi
