package main

import (
	"net"

	pb "github.com/AvinFajarF/proto"
	"github.com/AvinFajarF/service"
	"google.golang.org/grpc"
)

func main(){
	listen, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	grpc := grpc.NewServer()

	pb.RegisterUserServiceServer(grpc, &service.UserService{})
	grpc.Serve(listen)

}