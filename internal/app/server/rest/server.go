package rest

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/insaneadinesia/go-boilerplate/internal/app/container"
	"github.com/labstack/echo/v4"
)

func StartRestHttpService(container *container.Container) {
	server := echo.New()

	SetupMiddleware(server, container)
	SetupRouter(server, container)

	// Start server
	go func() {
		if err := server.Start(fmt.Sprintf(":%d", container.Config.AppHTTPPort)); err != nil {
			server.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
