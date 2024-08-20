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
	mock := mocks.NewIResponsibleRepository(t)
	cpf := "12345678910"
	mock.On("Get", context.Background(), &cpf).Return(&entity.Responsible{}, nil)
	service := NewResponsibleUseCase(mock)
	_, err := service.Get(context.Background(), &cpf)
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}

func TestResponsibleUseCase_Upadte(t *testing.T) {
	mock := mocks.NewIResponsibleRepository(t)
	responsible := entity.Responsible{Name: "adsasd", CPF: "12345678910", Email: "jorge@gmail.com"}
	Cresponsible := entity.Responsible{Name: "adsasd", CPF: "12345678910", Email: "jorge@gmail.com"}
	mock.On("Update", context.Background(), &responsible, &Cresponsible).Return(nil)
	service := NewResponsibleUseCase(mock)
	err := service.Update(context.Background(), &responsible, &Cresponsible)
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}

func TestResponsibleUseCase_Delete(t *testing.T) {
	mock := mocks.NewIResponsibleRepository(t)
	cpf := "12345678910"
	mock.On("Delete", context.Background(), &cpf).Return(nil)
	service := NewResponsibleUseCase(mock)
	err := service.Delete(context.Background(), &cpf)
	assert.Nil(t, err)
	mock.AssertExpectations(t)
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
