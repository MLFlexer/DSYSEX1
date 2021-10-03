package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "https://github.com/MLFlexer/DSYSEX1/tree/main/ituserver/ituserver"
)



const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedCourcesServer
}

func (s *server) GetCourses(ctx context.Context, in *pb.CourseRequest) (*pb.CourseReply, error) {
	return &pb.CourseReply{
		Id:"BSDISYS1KU",
		Name:"Distributed Systems, BSc",
		Teacher:"Agata Przybyszewska",
		CourseManager:"Marco Carbone",
		NumParticipants:195,
		Semester:"Autum 2021",
		ECTS:7.5,
		ProgrammingLang: "Golang"
	}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCoursesServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}