package main

import "fmt"

func main() {

	//create slice

	days := []string{"sunday", "monday", "tushday", "wednesday", "thursday", "friday"}

	fmt.Println(days)

	//for i := 0; i < len(days); i++ {
	//	fmt.Println(days[i])
	//}

	//for i := range days {
	//	fmt.Println(days[i])
	//}

	//index:=0
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

}
