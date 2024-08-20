package responsible

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/mocks"
)

func TestResponsibleUseCase_Create(t *testing.T) {
	mock := mocks.NewIResponsibleRepository(t)

	mock.On("Create", context.Background(), &entity.Responsible{}).Return(nil)

	service := NewResponsibleUseCase(mock)

	err := service.Create(context.Background(), &entity.Responsible{})

	assert.Nil(t, err)

	mock.AssertExpectations(t)
}

func TestResponsibleUseCase_Get(t *testing.T) {

}

func TestResponsibleUseCase_Upadte(t *testing.T) {

}

func TestResponsibleUseCase_Delete(t *testing.T) {

}

func TestResponsibleUseCase_SaveCard(t *testing.T) {

}

func TestResponsibleUseCase_UpdatePaymentMethod(t *testing.T) {

}

func TestResponsibleUseCase_CreateCustomer(t *testing.T) {

}

func TestResponsibleUseCase_UpdateCustomer(t *testing.T) {

}

func TestResponsibleUseCase_DeleteCustomer(t *testing.T) {

}

func TestResponsibleUseCase_CreatePaymentMethod(t *testing.T) {

}

func TestResponsibleUseCase_AttachPaymentMethod(t *testing.T) {

}
