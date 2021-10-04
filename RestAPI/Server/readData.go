package main

import (
	"fmt"
	"os"
)

func main() {

	jsonFile, err := os.Open("Data/students.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hey that's a pretty cool file")

	defer jsonFile.Close()
}
