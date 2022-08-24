package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time demo")

	presetTime := time.Now()
	fmt.Println(presetTime)

	fmt.Println(presetTime.Format("01-02-2021 Monday"))

	createDate := time.Date(2020, time.August, 10, 20, 20, 0, 0, time.Local)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2020 Monday"))
}
