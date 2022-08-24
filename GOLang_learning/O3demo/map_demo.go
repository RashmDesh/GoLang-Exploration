package main

import "fmt"

var empsal = map[string]int{"rashmi": 1,
	"rahul": 2,
	"karan": 3,
}

func main() {
	fmt.Println(empsal)

	fmt.Println(empsal["karan"])
	empsal["karan"] = 10

	fmt.Println(empsal["karan"])
}
