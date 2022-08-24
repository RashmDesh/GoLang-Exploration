package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files in GoLang")
	content := "This need to go in file.Leaning GoLang"

	file, err := os.Create("./mylocalfile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is", length)

	defer file.Close()

	readFile("./mylocalfile.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Content of file:", databyte)
	fmt.Println("Content of file:", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}

}
