package main

import "github.com/gin-gonic/gin"

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
	c.JSON(200, gin.H{
		"id":              "BSDISYS1KU",
		"name":            "Distributed Systems, BSc",
		"teacher":         "Agata Przybyszewska",
		"courseManager":   "Marco Carbone",
		"numParticipants": "195",
		"semester":        "Autum 2021",
		"ECTS":            "7.5",
		"programmingLang": "GoLang",
	})
}
