package main

import (
	"flag"
	"fmt"
)

func main() {
	num := flag.Int("value", 2, "looping through")
	flag.Parse()

	value := *num
	i := 0

	for i < value {
		fmt.Println("GO")
		i++
	}
}
// go run .\main.go -value 7  