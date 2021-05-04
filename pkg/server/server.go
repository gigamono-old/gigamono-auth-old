package server

import (
	"fmt"
	"net"

	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
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
	// Listener on TCP port.
	listener, err := net.Listen("tcp", fmt.Sprint(":", server.Config.Services.Types.Auth.Port))
	if err != nil {
		return err
	}

	// Create multiplexer and delegate content-types.
	multiplexer := cmux.New(listener)
	grpcListener := multiplexer.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := multiplexer.Match(cmux.HTTP1Fast())

	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return server.grpcServe(grpcListener) })
	grp.Go(func() error { return server.httpServe(httpListener) })
	grp.Go(func() error { return multiplexer.Serve() })
	return grp.Wait()
}
