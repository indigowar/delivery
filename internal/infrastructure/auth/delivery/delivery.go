package delivery

import (
	"context"

	"github.com/indigowar/delivery/internal/usecases/auth"
)

type Delivery interface {
	Run() error
	Shutdown(ctx context.Context) error

	AddService(service auth.Service)
}
