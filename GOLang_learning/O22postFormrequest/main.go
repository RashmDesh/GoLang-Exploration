package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("make GET request in GoLang")
	PerformPostFormRequest()

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

func PerformPostFormRequest() {
	const myurl = "http://localhost:8080/postform"

	//formdata

	data := url.Values{}
	data.Add("Firstname", "rashmi")
	data.Add("Lastname", "Deshmukh")

	respnce, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)

	}

	defer respnce.Body.Close()
	content, _ := ioutil.ReadAll(respnce.Body)

	fmt.Println(string(content))
}
