package main

import (
	"log"
	"net"

	"github.com/Bryan-BC/go-course-microservice/pkg/config"
	"github.com/Bryan-BC/go-course-microservice/pkg/db"
	"github.com/Bryan-BC/go-course-microservice/pkg/pb"
	"github.com/Bryan-BC/go-course-microservice/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Panicf("Error loading config file, %s \n", err)
	}

	db := db.Init(c.DBURL)

	listener, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Panicf("Error listening, %s \n", err)
	}

	log.Printf("Listening on port %s \n", c.Port)

	s := services.Server{
		DBPointer: &db,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCourseServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Panicf("Error serving course microservice, %s \n", err)
	}
}
