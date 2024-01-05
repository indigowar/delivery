package postgres

import (
	"context"
	"log"

	"github.com/indigowar/delivery/internal/domain"
	"github.com/jmoiron/sqlx"
)

type createAccount struct {
	db *sqlx.DB
}

// Create implements domain.CreateAccountUseCase.
func (u *createAccount) Create(ctx context.Context, account *domain.Account) error {
	query := `INSERT INTO accounts(id, phone, password) VALUES($1, $2, $3)`

	if _, err := u.db.Exec(query, account.ID(), account.Phone(), account.Password()); err != nil {
		// todo: add better error handling
		log.Println(err)
		return err
	}

	return nil
}

func NewCreateAccountUseCase(db *sqlx.DB) domain.CreateAccountUseCase {
	return &createAccount{
		db: db,
	}
}
