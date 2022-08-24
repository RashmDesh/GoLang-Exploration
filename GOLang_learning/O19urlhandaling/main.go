package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000?coursename=reactjs&paymentid=ghbj456"

func main() {

	fmt.Println("URL handaling")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of qparam is: %T", qparams)

	fmt.Println(qparams["coursename"])

	for _, value := range qparams {
		fmt.Println("Pram is :", value)

	}

	//create another url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutess",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
