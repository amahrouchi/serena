package main

import (
	"github.com/amahrouchi/serena/internal/core/handlers"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// TODO:
//   - Use docker-compose and Air for the dev environment
//   - install FX

func main() {
	log.Info().Msg("Starting Serena node...")

	// TODO: start the block routine
	log.Info().Msg("Starting block producer...")

	// Start the API server
	log.Info().Msg("Starting API server...")
	e := echo.New()

	healthHandler := handlers.NewHealthzHandler()
	e.GET("/healthz", healthHandler.Handle())

	e.Logger.Fatal(e.Start(":8080")) // TODO: get the port from a config file or env variable (or both)

	log.Info().Msg("Serena node started!")
}
