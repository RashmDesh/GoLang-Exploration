package main

import (
	"fmt"
	"time"
)

func main() {
	go numberSeies(11)
	numberSeies(15)

}

func numberSeies(value int) {

	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(value)
	}

}
