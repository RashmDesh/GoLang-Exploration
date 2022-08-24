package main

import (
	"fmt"
)

var arr1 [5]int = [5]int{1, 2, 3, 4, 5}

func main() {

	fmt.Println(arr1)
	arr1[2] = 30
	fmt.Println(arr1)

	arr2 := [...]int{10, 11, 12, 13}

	fmt.Printf("%v , %T", arr2, arr2)

}
