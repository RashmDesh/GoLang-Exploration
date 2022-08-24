package main

import "fmt"

var result string

func main() {
	logincount := 20

	if logincount < 18 {
		result = "under aged"
	} else if logincount > 30 {
		result = "over aged"
	} else {
		result = "applicable"
	}

	fmt.Println(result)

}
