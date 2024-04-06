package main

import (
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting Serena node...")

	// TODO: start the API server
	log.Info().Msg("Starting API server...")

	// TODO: start the block routine
	log.Info().Msg("Starting block producer...")

	log.Info().Msg("Serena node started!")
}
