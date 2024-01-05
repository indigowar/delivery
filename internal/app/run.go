package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	session "github.com/spazzymoto/echo-scs-session"

	"github.com/indigowar/delivery/internal/config"
	repository "github.com/indigowar/delivery/internal/repository/postgres"
	"github.com/indigowar/delivery/internal/services/auth"
	"github.com/indigowar/delivery/pkg/postgres"
)

// Run(*config.Config) - is the application's main code
func Run(cfg *config.Config) {
	postgresConnection, err := postgres.CreateConnection(cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Db, cfg.Postgres.User, cfg.Postgres.Password)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %s", err.Error())
	}

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	router := echo.New()

	setupRoutes(router, sessionManager, postgresConnection)

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

func setupRoutes(r *echo.Echo, sm *scs.SessionManager, connection *sqlx.DB) {
	_ = repository.NewGetAccountByIDUseCase(connection)
	getAccountByPhone := repository.NewGetAccountByPhoneUseCase(connection)
	createAccount := repository.NewCreateAccountUseCase(connection)

	// services
	authService := auth.NewService(getAccountByPhone, createAccount)

	r.Use(session.LoadAndSave(sm))

	r.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>It works</h1>")
	})

	// auth

	authHandler := auth.NewHandler(authService, sm)
	r.GET("/login", authHandler.ServeLoginPage("/login"))
	r.POST("/login", authHandler.HandleLoginRequest())

	r.GET("/register", authHandler.ServeRegistrationPage("/register"))
	r.POST("/register", authHandler.HandleRegisterRequest())

	// the other

}
