package core

import (
	"context"
	"github.com/amahrouchi/serena/internal/core/handlers"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// NewHTTPServer creates a new HTTP server.
func NewHTTPServer(lc fx.Lifecycle) *echo.Echo {
	// Create the Echo server
	e := echo.New()

	// Register the server with the lifecycle
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Declare routes
			healthHandler := handlers.NewHealthzHandler()
			e.GET("/healthz", healthHandler.Handle())

			// Start the server
			e.Logger.Fatal(e.Start(":8080"))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Close()
		},
	})

	return e
}
