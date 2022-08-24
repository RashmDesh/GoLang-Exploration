package main

import "fmt"

func main() {
	fmt.Println("Main function")
	greeter()
	result := adder(10, 20)
	fmt.Println("Addition is :", result)

	//proResult := proAdder(2, 4, 8, 9)
	fmt.Println("enter any list of number :", proAdder(2, 4, 8, 9))
}
func greeter() {
	fmt.Println("I am in function greeter")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) int {

	total := 0

	for _, val := range values {
		total += val

	}
	return total
}
