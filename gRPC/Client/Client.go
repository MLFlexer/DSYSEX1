package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "gRPC/Proto"
)

const (
	address   = "localhost:8080"
	defaultId = "BSDISYS1KU"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCourcesClient(conn)

	// Contact the server and print out its response.
	id := defaultId

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCourses(ctx, &pb.CourseRequest{Id: id})
	if err != nil {
		log.Fatalf("could not get course: %v", err)
	}
	log.Printf("Course id: %s", r.GetId())
	log.Printf("Course name: %s", r.GetName())
	log.Printf("Course Manager: %s", r.GetCourseManager())
	log.Printf("Course programming language: %s", r.GetProgrammingLang())
	log.Printf("Course semester: %s", r.GetSemester())
	log.Printf("Course ECTS: %s", fmt.Sprintf("%f", r.GetECTS()))
	log.Printf("Number of participants: %s", fmt.Sprint(r.GetNumParticipants()))
}
