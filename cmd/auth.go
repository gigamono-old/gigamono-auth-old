package main

import (
	"github.com/sageflow/sageflow/pkg/inits"
	"github.com/sageflow/sageflow/pkg/logs"

	"github.com/sageflow/sageauth/pkg/server"
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
