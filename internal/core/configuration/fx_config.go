package configuration

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"strconv"
)

// RegisterHooks registers the lifecycle hooks, starts/stops the Echo server.
func RegisterHooks(
	lc fx.Lifecycle,
	e *echo.Echo,
	c *Config,
	logger *zerolog.Logger,
) {
	// Register the server with the lifecycle
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start the server
			go func() {
				logger.Info().Msgf("Starting HTTP server on port %d", c.Port)
				err := e.Start(":" + strconv.Itoa(c.Port))
				if err != nil {
					e.Logger.Warn("The Echo server may have stopped unexpectedly. error=%+v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info().Msg("Shutting down HTTP server")
			return e.Shutdown(ctx)
		},
	})
}
