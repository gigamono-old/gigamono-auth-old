package main

import (
	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageflow/pkg/envs"
	"github.com/sageflow/sageflow/pkg/logs"

	"github.com/sageflow/sageauth/pkg/server"
)

func main() {
	// Set up log status file and load .env file.
	logs.SetStatusLogFile() // TODO. logs.SetStatusLogFile(config.Logging.Status.Filepath)
	envs.LoadEnvFile()      // TODO. Remove!

	// Connect to database.
	db := database.Connect() // TODO. database.Connect(config.db)

	// Start an auth gRPC server.
	server := server.NewAuthServer(db) // TODO. server.NewAuthServer(db, config)
	server.Listen("3003") // TODO. server.Listen(config.Server.Engine.Port)
}
