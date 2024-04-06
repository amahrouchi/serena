package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting Serena node...")

	// TODO: start the block routine
	log.Info().Msg("Starting block producer...")

	// Start the API server
	log.Info().Msg("Starting API server...")
	e := echo.New()
	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
	e.Logger.Fatal(e.Start(":8080")) // TODO: get the port from a config file or env variable (or both)

	log.Info().Msg("Serena node started!")
}
