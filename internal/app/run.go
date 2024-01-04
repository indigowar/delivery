package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/config"
)

// Run(*config.Config) - is the application's main code
func Run(cfg *config.Config) {
	router := echo.New()

	router.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>It works</h1>")
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Forced to stop the server: %s\n", err.Error())
	}

	log.Println("Graceful shutdown has ended")
}
