package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Orange"}
	fmt.Println(fruitList)

	fruitList = append(fruitList, "mango", "tomato")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:], "mango", "tomato")
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 10
	highScore[1] = 30
	highScore[2] = 10
	highScore[3] = 40
	//highScore[4] = 50 //error

	fmt.Println(highScore)

	highScore = append(highScore, 50, 90, 70)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
	fmt.Println(highScore)

	//remove value from slice based on index
	var index int = 2
	highScore = append(highScore[:index], highScore[index+1:]...)
	// print two slice 0-1 and 3-5 and remove 2
	fmt.Println(highScore)

}
