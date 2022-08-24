// handaling web request
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Local web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("respnse is of type %T", response)

	defer response.Body.Close() //CALLERS RESPONSIBLITY TO CLOSE CONNECTION

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)

	}

	content := string(databytes)
	fmt.Println(content)

}
