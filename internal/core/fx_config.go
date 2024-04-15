package core

import (
	"context"
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"strconv"
)

// registerHooks registers the lifecycle hooks, starts/stops the Echo server.
func registerHooks(
	lc fx.Lifecycle,
	e *echo.Echo,
	c *configuration.Config,
	logger *zerolog.Logger,
) {
	// Register the server with the lifecycle
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start the server
			go func() {
				logger.Info().Msgf("Starting server on port %d", c.Port)
				err := e.Start(":" + strconv.Itoa(c.Port))
				if err != nil {
					e.Logger.Errorf("Echo server failed to start. error=%+v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
