package accounts

import (
	"context"
	"net/mail"
	"os"

	"github.com/google/uuid"
)

// profileUpdater - is an implementation of ProfileUpdater interface
type profileUpdater struct {
	storage Storage
}

// LinkEmailToAccount implements ProfileUpdater
func (svc *profileUpdater) LinkEmailToAccount(ctx context.Context, id uuid.UUID, addr *mail.Address) error {
	// todo: implement
	panic("unimplemented")
}

// UpdateFirstName implements ProfileUpdater
func (svc *profileUpdater) UpdateFirstName(ctx context.Context, id uuid.UUID, firstName string) error {
	// todo: implement
	panic("unimplemented")
}

// UpdateSurname implements ProfileUpdater
func (svc *profileUpdater) UpdateSurname(ctx context.Context, id uuid.UUID, surname string) error {
	// todo: implement
	panic("unimplemented")
}

// LoadProfileImage implements ProfileUpdater
func (svc *profileUpdater) LoadProfileImage(ctx context.Context, id uuid.UUID, image os.File) error {
	// todo: implement
	panic("unimplemented")
}

func NewProfileUpdater(storage Storage) ProfileUpdater {
	return &profileUpdater{
		storage: storage,
	}
}
