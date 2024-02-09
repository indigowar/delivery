package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/indigowar/delivery/internal/infrastructure/accounts/delivery"
	"github.com/indigowar/delivery/internal/infrastructure/accounts/delivery/rest"
	"github.com/indigowar/delivery/internal/infrastructure/accounts/storage/postgres"
	"github.com/indigowar/delivery/internal/usecases/accounts"
)

func main() {

	// added due to errors, when database container is started,
	// but db isn't initialized.
	//
	// Without it there is possbillity that service will not start.
	time.Sleep(1 * time.Second)

	storage, err := createPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	finder := accounts.NewFinder(storage)
	registrator := accounts.NewRegistrator(storage)
	credentialsValidator := accounts.NewCredentialsValidator(storage)
	profileUpdater := accounts.NewProfileUpdater(storage, nil)

	var delivery delivery.Delivery = rest.NewDelivery(80)

	delivery.AddFinder(finder)
	delivery.AddRegistrator(registrator)
	delivery.AddCredentialsValidator(credentialsValidator)
	delivery.AddProfileUpdater(profileUpdater)

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
		log.Fatalf("Failed to shutdown the delivery: %s\n", err)
	}

	log.Println("Service is stopped")
}

func createPostgresStorage() (*postgres.Storage, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	return postgres.NewStorage(host, port, user, password, dbName)
}
