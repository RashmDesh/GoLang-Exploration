package main

import "fmt"

func main() {
	//var number int
	number := 10
	fmt.Println("Numbe is: ", number)

	var ptr *int

	ptr = &number
	fmt.Println("Pointer actual value is: ", ptr)
	fmt.Println("Pointer refrance value is: ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value", *ptr)

}
