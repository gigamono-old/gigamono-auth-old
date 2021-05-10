package server

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// AuthServer is a grpc server for managing authentication and authorisation.
type AuthServer struct {
	*gin.Engine
	inits.App
}

// NewAuthServer creates a new server instance.
func NewAuthServer(app inits.App) AuthServer {
	return AuthServer{
		Engine: gin.Default(),
		App:    app,
	}
}

// Listen makes the server listen on specified port.
func (server *AuthServer) Listen() error {
	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return server.grpcServe() })
	grp.Go(func() error { return server.httpServe() })
	return grp.Wait()
}
