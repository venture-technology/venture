package repository

import (
	"github.com/venture-technology/venture/internal/entity"
)

type IResponsibleRepository interface {
	Create(responsible *entity.Responsible) error
	Get(cpf string) (*entity.Responsible, error)
	Update(responsible *entity.Responsible, attributes map[string]interface{}) error
	Delete(cpf string) error
	SaveCard(cpf, cardToken, paymentMethodId string) error
	Auth(responsible *entity.Responsible) (*entity.Responsible, error)
	FindByEmail(email string) (*entity.Responsible, error)
}
