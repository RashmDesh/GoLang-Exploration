package main

import "fmt"

func main() {
	var number1, number2 int

	fmt.Println("Enter first number :")

	fmt.Scanln(&number1)

	fmt.Println("Enter secand number number :")
	fmt.Scanln(&number2)

	//Addition of two number
	addition := number1 + number2

	fmt.Println("Addition of two number is:", addition)

}
