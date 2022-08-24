package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("make GET request in GoLang")
	PerformGetRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:8080/get"

	responce, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	//fmt.Println("Status code", responce.StatusCode)
	//fmt.Println("Content length is :", responce.ContentLength)

	content, _ := ioutil.ReadAll(responce.Body)

	//fmt.Println(string(content))

	//another wat to read url
	var responceString strings.Builder
	byteCount, _ := responceString.Write(content)

	fmt.Println("Bytecount is :", byteCount)
	fmt.Println(responceString.String())

}
