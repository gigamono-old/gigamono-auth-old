package server

import (
	"context"
	"net"

	"github.com/gigamono/gigamono/pkg/logs"
	"github.com/gigamono/gigamono/pkg/services/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (server *AuthServer) grpcServe(listener net.Listener) error {
	grpcServer := grpc.NewServer() // Create a gRPC server.

	// Register gRPC service.
	generated.RegisterAuthServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener) // Listen for requests.
}

// Ping says Pong.
func (server *AuthServer) Ping(ctx context.Context, msg *generated.Empty) (*generated.PingResponse, error) {
	authMsg := "Auth replies: Pong!"
	logs.FmtPrintln(authMsg)
	response := generated.PingResponse{
		Content: authMsg,
	}
	return &response, nil
}

// GetSignUpToken
func (server *AuthServer) GetSignUpToken(ctx context.Context, msg *generated.UserTokenRequest) (*generated.UserTokenResponse, error) {
	response := generated.UserTokenResponse{
		AccessToken:  "",
		RefreshToken: "",
	}
	return &response, nil
}

// GetSignInToken
func (server *AuthServer) GetSignInToken(ctx context.Context, msg *generated.UserTokenRequest) (*generated.UserTokenResponse, error) {
	logs.FmtPrintln("Successfully Called")
	response := generated.UserTokenResponse{
		AccessToken:  "",
		RefreshToken: "",
	}
	return &response, nil
}

// RefreshAccessToken
func (server *AuthServer) RefreshAccessToken(ctx context.Context, msg *generated.AccessTokenRequest) (*generated.UserTokenResponse, error) {
	response := generated.UserTokenResponse{
		AccessToken:  "",
		RefreshToken: "",
	}
	return &response, nil
}
