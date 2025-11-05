package grpc

import (
	"fmt"
	"net"

	userpb "github.com/JustSelection/project-protos/proto/users"
	"github.com/JustSelection/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc user.Service) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("could not listen on port 50051: %w", err)
	}

	grpcServer := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpcServer, NewHandler(svc))

	fmt.Println("gRPC server listening on port :50051")
	return grpcServer.Serve(listener)
}
