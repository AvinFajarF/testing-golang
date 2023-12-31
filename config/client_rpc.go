package config


import (
	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addrRPC = flag.String("gRPC address", "localhost:8080", "Address for dialing gRPC Server")
)

var errorRPC error
var RpcClient *grpc.ClientConn

func ConnectRPC() {
	flag.Parse()

	RpcClient, errorRPC = grpc.Dial(*addrRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if errorRPC != nil {
		panic("Failed to connect gRPC Server!")
	}

}