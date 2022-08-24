package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("make GET request in GoLang")
	PerformPostRequest()

}

func PerformPostRequest() {
	const myurl = "http://localhost:8080/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"golang",
			"price":0
		}
	`)

	responce, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	content, _ := ioutil.ReadAll(responce.Body)
	fmt.Println(string(content))
}
