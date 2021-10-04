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
	c.JSON(200, gin.H{
		"name": "Marco Polo",
		"id":   "007",
	})
}

func getStudent(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "Lars Larsen",
		"id":   "000001",
	})
}

func getCourse(c *gin.Context) {
	courses := Data.ReadCourses()
	c.JSON(200, courses.Courses)
}
