package main

import (
	"github.com/gin-gonic/gin"

	"github.com/MLFlexer/DSYSEX1/Data"
)

func main() {
	r := gin.Default()
	r.GET("/student", getStudent)
	r.GET("/teacher", getTeacher)
	r.GET("/course", getCourse)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getTeacher(c *gin.Context) {
	teachers := Data.ReadTeachers()
	c.JSON(200, teachers.Teachers)
}

func getStudent(c *gin.Context) {
	students := Data.ReadStudents()
	c.JSON(200, students.Students)
}

func getCourse(c *gin.Context) {
	courses := Data.ReadCourses()
	c.JSON(200, courses.Courses)
}
