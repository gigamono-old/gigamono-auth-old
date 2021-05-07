package server

import (
	"context"
	"fmt"
	"net"

	"github.com/gigamono/gigamono/pkg/services/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (server *AuthServer) grpcServe(listener net.Listener) error {
	grpcServer := grpc.NewServer() // Create a gRPC server.

	// Register gRPC service.
	generated.RegisterAuthServer(grpcServer, server)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener) // Listen for requests.
}

// SayHello replies with message.
func (server *AuthServer) SayHello(ctx context.Context, msg *generated.Message) (*generated.Message, error) {
	serverMsg := "Auth replies: " + msg.Content
	fmt.Println(serverMsg)
	response := generated.Message{
		Content: serverMsg,
	}
	return &response, nil
}
