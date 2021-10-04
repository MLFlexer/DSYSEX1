package Data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

func ReadCourses() Courses {

	jsonFile, err := os.Open("Data/database.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var courses Courses

	json.Unmarshal([]byte(byteValue), &courses)

	return courses
}
