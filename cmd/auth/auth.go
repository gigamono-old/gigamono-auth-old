package main

import (
	"github.com/sageflow/sageflow/pkg/secrets"
	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageflow/pkg/configs"
	"github.com/sageflow/sageflow/pkg/logs"

	"github.com/sageflow/sageauth/pkg/server"
)

func main() {
	// Set up log status file.
	logs.SetStatusLogFile()

	// Load sageflow config file.
	config, err := configs.LoadSageflowConfig()
	if err != nil {
		logs.FmtPrintln("Unable to load config file:", err)
	}

	// Set up secret manager,
	secrets, err := secrets.NewManager(&config)
	if err != nil {
		logs.FmtPrintln("Unable to create secrets manager:", err)
	}

	// Connect to database.
	db, err := database.Connect(secrets)
	if err != nil {
		logs.FmtPrintln("Unable to connect to database:", err)
	}

	// Start an auth gRPC server.
	server := server.NewAuthServer(&db) // TODO. server.NewAuthServer(db, config)
	server.Listen("3003") // TODO. server.Listen(config.Server.Engine.Port)
}
