package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languanges", languages)
	fmt.Println("JS shots for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languanges", languages)

	//loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n ", key, value)

	}
}
