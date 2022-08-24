package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status string
	Age    int
}

func main() {
	rashmi := User{"Rashmi", "rashmi.deshmukh@cctech.co.in", "married", 28}
	fmt.Println("Rashmi details are :=", rashmi)
	fmt.Printf("Rashmi details are := %+v\n", rashmi)
	fmt.Printf("Name is %v and email is %v", rashmi.Name, rashmi.Email)

}
