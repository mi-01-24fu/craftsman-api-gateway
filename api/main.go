package main

import (
	"craftsman-api-gateway/controller"
	signuppb "craftsman-api-gateway/pkg/grpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	protocol := "tcp"
	port := 8081

	listenPort, err := net.Listen(protocol, fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	signuppb.RegisterNewMemberRegistrationServiceServer(server, controller.NewNewMemberRegistrationService())

	reflection.Register(server)

	server.Serve(listenPort)
}
