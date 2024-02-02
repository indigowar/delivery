package delivery

import (
	"context"

	"github.com/indigowar/delivery/internal/usecases/accounts"
)

type Delivery interface {
	Run() error
	Shutdown(ctx context.Context) error

	AddFinder(finder accounts.Finder)
	AddRegistrator(registrator accounts.Registrator)
	AddCredentialsValidator(validator accounts.CredentialsValidator)
	AddProfileUpdater(updater accounts.ProfileUpdater)
}
