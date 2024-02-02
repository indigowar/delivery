package postgres

import (
	"net/mail"
	"net/url"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/indigowar/delivery/internal/entities"
	"github.com/indigowar/delivery/pkg/postgres/fields"
)

type account struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	Phone           string    `gorm:"unique"`
	Password        string
	Email           *fields.MailAdress `gorm:"column:email"`
	FirstName       *string            `gorm:"column:first_name"`
	Surname         *string            `gorm:"column:surname"`
	ProfileImageUrl *fields.URL        `gorm:"column:profile_image_url"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func fromEntity(model *entities.Account) account {
	return account{
		ID:              model.ID(),
		Phone:           model.Phone(),
		Password:        model.Password(),
		Email:           (*fields.MailAdress)(model.Email()),
		FirstName:       model.FirstName(),
		Surname:         model.Surname(),
		ProfileImageUrl: (*fields.URL)(model.ProfileImageUrl()),
	}
}

func toEntity(a account) *entities.Account {
	return entities.NewAccountWith(
		a.ID,
		a.Phone,
		a.Password,
		(*mail.Address)(a.Email),
		a.FirstName,
		a.Surname,
		(*url.URL)(a.ProfileImageUrl),
	)
}
