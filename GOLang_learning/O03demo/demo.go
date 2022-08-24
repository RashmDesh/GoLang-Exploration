package main

import "fmt"

var number1 int = 10
var number2 int = 30
var add int

func main() {
	number3 := 50
	number4 := 10

	add = number3 + number4
	fmt.Println("add", add)

	add = number1 + number2
	fmt.Println("add", add)

	fmt.Printf("%v,%T", number1, number1)
	fmt.Printf("%v,%T", number3, number3)
}
