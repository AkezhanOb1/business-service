package main

import (
	"log"
	"net"

	pb "github.com/AkezhanOb1/business-service/api"
	"github.com/AkezhanOb1/business-service/config"
	"google.golang.org/grpc"

	"github.com/AkezhanOb1/business-service/service"
)

func main() {
	lis, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", config.ServerAddress)

	s := grpc.NewServer()
	pb.RegisterBusinessServicesServer(s, &service.Server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
