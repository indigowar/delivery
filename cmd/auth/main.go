package main

import (
	"context"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"

	accountsrest "github.com/indigowar/delivery/internal/infrastructure/auth/credentials_validator/accounts_rest"
	"github.com/indigowar/delivery/internal/infrastructure/auth/delivery"
	"github.com/indigowar/delivery/internal/infrastructure/auth/delivery/rest"
	"github.com/indigowar/delivery/internal/infrastructure/auth/storage/redis"
	"github.com/indigowar/delivery/internal/infrastructure/auth/token_generator/jwt"
	"github.com/indigowar/delivery/internal/usecases/auth"
)

func main() {
	storage, err := createStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	credentialsValidator, err := createCredentialsValidator()
	if err != nil {
		log.Fatal(err)
	}

	tokenGenerator := jwt.NewTokenGenerator()

	service, err := createService(storage, credentialsValidator, tokenGenerator)
	if err != nil {
		log.Fatal(err)
	}

	var delivery delivery.Delivery = rest.NewDelivery(80)

	delivery.AddService(service)

	go func() {
		if err := delivery.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := delivery.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Service is stopped")
}

func createStorage() (*redis.Storage, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	return redis.NewStorage(host, port, password)
}

func createCredentialsValidator() (auth.CredentialsValidator, error) {
	accountsHost, err := url.Parse(os.Getenv("AUTH_CREDENTIAL_VALIDATOR_HOST")) // todo: accounts host
	if err != nil {
		return nil, err
	}

	return accountsrest.NewCredentialsValidator(accountsHost), nil
}

func createService(
	storage auth.SessionStorage,
	validator auth.CredentialsValidator,
	generator auth.ShortLiveTokenGenerator,
) (auth.Service, error) {
	secret := os.Getenv("AUTH_SECRET")

	sessionTtlInHours, err := strconv.Atoi(os.Getenv("AUTH_SESSION_TTL"))
	if err != nil {
		return nil, err
	}

	sessionTtl := time.Duration(sessionTtlInHours) * time.Hour

	accessTtlInMinutes, err := strconv.Atoi(os.Getenv("AUTH_ACCESS_TTL"))
	if err != nil {
		return nil, err
	}

	accessTtl := time.Duration(accessTtlInMinutes) * time.Minute

	return auth.NewServiceImplementation(
		auth.WithStorage(storage),
		auth.WithCredentialsValidator(validator),
		auth.WithTokenGenerator(generator),

		auth.WithSecret([]byte(secret)),

		auth.WithSessionLifetime(sessionTtl),
		auth.WithAccessTokenLifetime(accessTtl),
	), nil
}
