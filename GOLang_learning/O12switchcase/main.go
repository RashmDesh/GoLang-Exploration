package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6)
	fmt.Println("value of dice", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open game.")
	case 2:
		fmt.Println("Move 2 spot")
	case 3:
		fmt.Println("Move 3 spot")
		//fallthrough // consider next case too
	case 4:
		fmt.Println("Move 4 spot")

	case 5:
		fmt.Println("Move 5 spot")
	case 6:
		fmt.Println("Move 6 spot and roll dice agian")
	default:
		fmt.Println("what you want to do! Play again")

	}

}
