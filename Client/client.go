package main

import (
	"fmt"
	"io"
	"net/http"
)

type Course struct {
	id              string
	name            string
	teacher         string
	courseManager   string
	numParticipants int
	ECTS            float64
	programmingLang string
}

func main() {
	r, error := http.Get("http://localhost:8080/course")
	if error != nil {
		fmt.Println("Error: ", error)
	}
	/*
			err := json.NewDecoder(r.Body).Decode(&c)
		    if err != nil {
		        http.Error(w, err.Error(), http.StatusBadRequest)
		        return
		    }
	*/
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	myString := string(body[:])
	fmt.Print(myString)
}
