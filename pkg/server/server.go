package server

import (
	"context"
	"fmt"
	"net"

	"github.com/sageflow/sageauth/internal/proto"
	"github.com/sageflow/sageflow/pkg/inits"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// AuthServer is a grpc server for managing authentication and authorisation.
type AuthServer struct {
	inits.App
}

// NewAuthServer creates a new server instance.
func NewAuthServer(app inits.App) AuthServer {
	return AuthServer{app}
}

// Listen starts a new gRPC server that listens on specified port.
func (server *AuthServer) Listen() error {
	// Listen on port using TCP.
	listener, err := net.Listen("tcp", fmt.Sprint(":", server.Config.Server.Auth.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer() // Create a gRPC server.

	// Register gRPC service.
	proto.RegisterAuthServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener) // Listen for requests.
}

// SayHello says Hello
func (server *AuthServer) SayHello(ctx context.Context, msg *proto.Message) (*proto.Message, error) {
	authMsg := "Auth replies: " + msg.Content
	fmt.Println(authMsg)
	response := proto.Message{
		Content: authMsg,
	}
	return &response, nil
}
