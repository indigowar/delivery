package postgres

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/indigowar/delivery/internal/domain"
	"github.com/jmoiron/sqlx"
)

type getAccountByID struct {
	db *sqlx.DB
}

// GetByID implements domain.GetAccountByIDUseCase.
func (u *getAccountByID) GetByID(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	query := `SELECT a.id, a.phone, a.password FROM accounts a WHERE a.id = $1`
	result := make([]domain.Account, 0)

	if err := u.db.Select(&result, query, id); err != nil {
		// todo: add better handling the error
		log.Println(err)
		return nil, err
	}
	return &result[0], nil
}

func NewGetAccountByIDUseCase(db *sqlx.DB) domain.GetAccountByIDUseCase {
	return &getAccountByID{
		db: db,
	}
}
