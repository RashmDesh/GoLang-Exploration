package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter input:")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n') // read string up \n

	fmt.Println("Your input is:", input)

	//if inpur is in number formate
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("1 is added in numRating", numRating+1)
	}

}
