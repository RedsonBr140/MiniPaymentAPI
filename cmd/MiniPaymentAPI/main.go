package main

import (
	"context"
	"fmt"

	"github.com/RedsonBr140/MiniPaymentAPI/internal/app"
	"github.com/RedsonBr140/MiniPaymentAPI/internal/config"
	"github.com/RedsonBr140/MiniPaymentAPI/internal/database"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.Info("Initializing...")

	// Load the .env file
	appConfig, err := config.InitConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	log.Info(".env has been loaded.")

	// Connect to the database
	ctx := context.Background()

	pg, err := database.Initialize(&ctx, fmt.Sprintf("postgres://%s:%s@%s/%s",
		appConfig.Postgres.User,
		appConfig.Postgres.Password,
		appConfig.Postgres.Host,
		appConfig.Postgres.DBName))

	if err != nil {
		log.Fatal("Connecting to the database: ", err)
	}

	defer pg.DB.Close(ctx)

	log.Info("Connected to the PostgreSQL database!")

	// Load routes
	router := app.InitRouter()
	// Start API server
	if err := router.Run(); err != nil {
		log.Fatal("Initializing server: ", err)
	}
}
