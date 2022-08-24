// Accept two number and swap that number
package main

import "fmt"

func main() {

	fmt.Println("Enter two number for swapping :")

	fmt.Println("Enter First number")
	var number1 int
	fmt.Scanln(&number1)

	fmt.Println("Enter Secand number")
	var number2 int
	fmt.Scanln(&number2)

	fmt.Println("Number before wapping is :", number1, number2)

	var swapBucket, _ int

	swapBucket = number1
	number1 = number2
	number2 = swapBucket

	fmt.Println("Number after swapping :", number1, number2)
}
