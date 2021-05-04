package main

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/logs"

	"github.com/gigamono/gigamono-auth/pkg/server"
)

func main() {
	// Initialises app.
	app, err := inits.NewApp("Auth")
	if err != nil {
		logs.FmtPrintln("Unable to initialize auth:", err)
		return
	}

	// Start an auth gRPC server.
	server := server.NewAuthServer(app)
	if err := server.Listen(); err != nil {
		logs.FmtPrintln("Unable to listen on port specified:", err)
	}
}
