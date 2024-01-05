package postgres

import (
	"context"
	"log"

	"github.com/indigowar/delivery/internal/domain"
	"github.com/jmoiron/sqlx"
)

type getAccountByPhone struct {
	db *sqlx.DB
}

// GetByPhone implements domain.GetAccountByPhoneUseCase.
func (u *getAccountByPhone) GetByPhone(ctx context.Context, phone string) (*domain.Account, error) {
	query := `SELECT a.id, a.phone, a.password FROM accounts a WHERE a.phone = $1`
	result := make([]domain.Account, 0)

	if err := u.db.Select(&result, query, phone); err != nil {
		// todo: add better handling the error
		log.Println(err)
		return nil, err
	}
	return &result[0], nil
}

func NewGetAccountByPhoneUseCase(db *sqlx.DB) domain.GetAccountByPhoneUseCase {
	return &getAccountByPhone{
		db: db,
	}
}
