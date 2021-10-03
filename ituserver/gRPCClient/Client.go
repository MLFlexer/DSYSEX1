package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "https://github.com/MLFlexer/DSYSEX1/tree/main/ituserver/ituserver"
)

const (
	address = "localhost:8080"
	defaultId = "BSDISYS1KU"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCoursesClient(conn)

	// Contact the server and print out its response.
	id := defaultId

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCourses(ctx, &pb.CourseRequest{Id: id})
	if err != nil {
		log.Fatalf("could not get course: %v", err)
	}
	log.Printf("course: %s", r.GetMessage())
}