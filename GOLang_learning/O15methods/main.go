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
	//fmt.Println("Rashmi details are :=", rashmi)
	fmt.Printf("Rashmi details are := %+v\n", rashmi)
	//fmt.Printf("Name is %v and email is %v", rashmi.Name, rashmi.Email)

	rashmi.GetStatus()
	rashmi.NewMail()
	fmt.Printf("Rashmi details are := %+v\n", rashmi)

}

//if you want to change value of user use pointer *User
func (u User) GetStatus() {
	fmt.Println("Is User active :", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is:", u.Email)
}
