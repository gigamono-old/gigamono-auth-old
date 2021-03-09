package server

import (
	"net"
	"net/http"

	routes "github.com/sageflow/sageauth/internal/routes/v1"
)

func (server *AuthServer) httpServe(listener net.Listener) error {
	server.setRoutes() // Set routes.

	// Use http server.
	httpServer := &http.Server{
		Handler: server,
	}

	return httpServer.Serve(listener)
}

func (server *AuthServer) setRoutes() {
	v1 := server.Group("/v1")
	routes.HandleTokenRoutes(v1, &server.App)
}
