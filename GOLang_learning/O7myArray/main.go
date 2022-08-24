package main

import "fmt"

func main() {
	fmt.Println("Array demo")

	var friuitlist [4]string

	friuitlist[0] = "Aplle"
	friuitlist[1] = "oragne"
	friuitlist[3] = "tomato"

	fmt.Println("Fruit list is :", friuitlist)
	fmt.Println("Fruit list length is :", len(friuitlist))

	var number = [...]int{1, 2, 3, 4, 5, 6} // array with no size
	fmt.Println("Number list is :", number)

	number[0] = 10
	fmt.Println("Updated Number list is :", number)

}
