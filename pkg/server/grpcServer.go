package server

import (
	"context"
	"net"

	"github.com/sageflow/sageauth/internal/proto"
	"github.com/sageflow/sageflow/pkg/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (server *AuthServer) grpcServe(listener net.Listener) error {
	grpcServer := grpc.NewServer() // Create a gRPC server.

	// Register gRPC service.
	proto.RegisterAuthServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener) // Listen for requests.
}

// Ping says Pong.
func (server *AuthServer) Ping(ctx context.Context, msg *proto.AuthNull) (*proto.PingResponse, error) {
	authMsg := "Auth replies: Pong!"
	logs.FmtPrintln(authMsg)
	response := proto.PingResponse{
		Content: authMsg,
	}
	return &response, nil
}

// GetSignUpToken
func (server *AuthServer) GetSignUpToken(ctx context.Context, msg *proto.UserTokenRequest) (*proto.UserTokenResponse, error) {
	response := proto.UserTokenResponse{
		AccessToken:  "",
		RefreshToken: "",
	}
	return &response, nil
}

// GetSignInToken
func (server *AuthServer) GetSignInToken(ctx context.Context, msg *proto.UserTokenRequest) (*proto.UserTokenResponse, error) {
	response := proto.UserTokenResponse{
		AccessToken:  "",
		RefreshToken: "",
	}
	return &response, nil
}

// RefreshAccessToken
func (server *AuthServer) RefreshAccessToken(ctx context.Context, msg *proto.AccessTokenRequest) (*proto.UserTokenResponse, error) {
	response := proto.UserTokenResponse{
		AccessToken:  "",
		RefreshToken: "",
	}
	return &response, nil
}
