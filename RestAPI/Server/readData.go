package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Students struct {
	Students []Student `json:"students"`
}

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Enrollments struct {
	Enrollments []Enrollment `json:"Enrollments"`
}

type Enrollment struct {
	StudentId string `json:"studentid"`
	CourseId  string `json:"courseid"`
}

type Teachers struct {
	Teachers []Teacher `json:"teachers"`
}

type Teacher struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Teachings struct {
	Teachings []Teaching `json:"teaching"`
}

type Teaching struct {
	TeacherId string `json:"Teacherid"`
	CourseId  string `json:"courseid"`
}

type TeacherRatings struct {
	TeacherRatings []TeacherRating `json:"teacherratings"`
}

type TeacherRating struct {
	StudentId string `json:"studentid"`
	TeacherId string `json:"teacherid"`
	Rating    int    `json:"rating"`
}

type Courses struct {
	Courses []Course `json:"courses"`
}

type Course struct {
	Id              string  `json:"id"`
	Name            string  `json:"name"`
	NumParticipants int     `json:"numparticipants"`
	Semester        string  `json:"semester"`
	ECTS            float32 `json:"ECTS"`
	ProgrammingLang string  `json:"programminglang"`
}
type CourseRatings struct {
	CourseRatings []CourseRating `json:"courseratings"`
}

type CourseRating struct {
	StudentId string `json:"studentid"`
	CourseId  string `json:"courseid"`
	Rating    int    `json:"rating"`
}

func main() {

	jsonFile, err := os.Open("Data/database.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hey that's a pretty cool file")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var courses Courses

	json.Unmarshal([]byte(byteValue), &courses)

	for i := 0; i < len(courses.Courses); i++ {
		fmt.Println("Course Code: " + courses.Courses[i].Id)
		fmt.Println("Course Name: " + courses.Courses[i].Name)
		fmt.Println("Number of Participants: " + strconv.Itoa(courses.Courses[i].NumParticipants))
	}
}
