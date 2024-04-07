package core

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// RegisterHooks registers the lifecycle hooks, starts/stops the Echo server.
func RegisterHooks(lc fx.Lifecycle, e *echo.Echo) {
	// Register the server with the lifecycle
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start the server
			go func() {
				err := e.Start(":8080") // TODO: make the port configurable
				if err != nil {
					e.Logger.Errorf("Echo server failed to start. error=%v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
