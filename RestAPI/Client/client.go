package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Print("Courses \n")
	fmt.Print(requestGetList("http://localhost:8080/course"))
	fmt.Print("\n")
	fmt.Print("Students \n")
	fmt.Print(requestGetList("http://localhost:8080/student"))
	fmt.Print("\n")
	fmt.Print("Teachers \n")
	fmt.Print(requestGetList("http://localhost:8080/teacher"))
	fmt.Print("\n")
}

func requestGetList(url string) string {
	r, error := http.Get(url)
	if error != nil {
		fmt.Println("Error: ", error)
	}

	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	myString := string(body[:])

	return myString
}
