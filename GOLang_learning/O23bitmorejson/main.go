package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"` // Nmae is changes to courseNmae
	Price    int
	Platform string
	Password string `json:"_"` //no need to password
	Tags     []string
}

func main() {
	fmt.Println("Converting any type of data into json")
	//EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCouses := []course{
		{"React js Booycamp", 2000, "LearnCode", "123456789", []string{"web-dev", "js"}},
		{"Data science  Booycamp", 2000, "LearnCode", "123456789", []string{"data science", "statistics"}},
		{"Data eng Booycamp", 2000, "LearnCode", "123456789", []string{"cloude", "data"}},
	}

	//package this data as json data

	finalJson, err := json.MarshalIndent(lcoCouses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"courseName": "Data eng Booycamp",
		"Price": 2000,
		"Platform": "LearnCode",
		"_": "123456789",
		"Tags": ["cloude","data"]
	}
	
	`)

	var lcoCourse course

	checkvalid := json.Valid(jsonDataFromWeb)

	if checkvalid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json not valid")
	}

	// some cases where you just want to add data to key value
	//var myOnlineData map[string]interface{}
	//json.Unmarshal(json)
}
