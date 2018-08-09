package main

import (
	"log"
	"net"

	pb "github.com/vladimirvivien/workbench/grpc/protobuf"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// GreetService type implements gRPC service
type GreetService struct{}

// Greet returns a greeting msg
func (c *GreetService) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {

	name := req.GetName()
	return &pb.GreetResponse{Message: "Hello," + name}, nil

}

func main() {
	lstnr, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to start server:", err)
	}

	grpcSvr := grpc.NewServer()
	pb.RegisterGreeterServer(grpcSvr, &GreetService{})
	grpcSvr.Serve(lstnr)
}
