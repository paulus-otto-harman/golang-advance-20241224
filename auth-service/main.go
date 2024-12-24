package main

import (
	pb "20241224/auth-service/proto"
	"20241224/auth-service/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &service.AuthService{})

	if err = grpcServer.Serve(lis); err != nil {
		log.Println(err)
		return
	}
}
