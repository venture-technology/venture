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
	pm := "pm_1PpggZLfFDLpePGL4Rl81sWf"
	mock := mocks.NewIResponsibleRepository(t)
	responsible := entity.Responsible{
		CPF: "12345678910",
		CreditCard: entity.CreditCard{
			CardToken: "tok_visa",
			Default:   true,
		}}
	mock.On("SaveCard", context.Background(), &responsible.CPF, &responsible.CreditCard.CardToken, &pm).Return(nil)
	service := NewResponsibleUseCase(mock)
	err := service.SaveCard(context.Background(), &responsible.CPF, &responsible.CreditCard.CardToken, &pm)
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}

func TestResponsibleUseCase_CreateCustomer(t *testing.T) {
	mock := mocks.NewIResponsibleRepository(t)
	responsible := entity.Responsible{
		Name:  "CreateCustomer",
		Email: "CreateCustomer@gmail.com",
		Phone: "+55 11 123456789",
	}
	service := NewResponsibleUseCase(mock)
	_, err := service.CreateCustomer(context.Background(), &responsible)
	if err != nil {
		t.Errorf("erro ao criar customer na stripe")
	}
}

func TestResponsibleUseCase_UpdateCustomer(t *testing.T) {
	mock := mocks.NewIResponsibleRepository(t)
	responsible := entity.Responsible{
		Name:  "CreateCustomer",
		Email: "CreateCustomer@gmail.com",
		Phone: "+55 11 123456789",
	}
	service := NewResponsibleUseCase(mock)
	cus, err := service.CreateCustomer(context.Background(), &responsible)
	if err != nil {
		t.Errorf("erro ao criar customer na stripe")
	}
	responsible.CustomerId = cus.ID
	_, err = service.UpdateCustomer(context.Background(), &responsible)
	if err != nil {
		t.Errorf("erro ao atualizar customer")
	}
}

func TestResponsibleUseCase_DeleteCustomer(t *testing.T) {
	mock := mocks.NewIResponsibleRepository(t)
	responsible := entity.Responsible{
		Name:  "CreateCustomer",
		Email: "CreateCustomer@gmail.com",
		Phone: "+55 11 123456789",
	}
	service := NewResponsibleUseCase(mock)
	cus, err := service.CreateCustomer(context.Background(), &responsible)
	if err != nil {
		t.Errorf("erro ao criar customer na stripe")
	}
	responsible.CustomerId = cus.ID
	_, err = service.DeleteCustomer(context.Background(), responsible.CustomerId)
	if err != nil {
		t.Errorf("erro ao deleter customer")
	}
}

func TestResponsibleUseCase_CreatePaymentMethod(t *testing.T) {
	mock := mocks.NewIResponsibleRepository(t)
	responsible := entity.Responsible{
		Name:  "CreateCustomer",
		Email: "CreateCustomer@gmail.com",
		Phone: "+55 11 123456789",
		CreditCard: entity.CreditCard{
			CardToken: "tok_visa",
			Default:   true,
		},
	}
	service := NewResponsibleUseCase(mock)
	cus, err := service.CreateCustomer(context.Background(), &responsible)
	if err != nil {
		t.Errorf("erro ao criar customer na stripe")
	}
	responsible.CustomerId = cus.ID
	_, err = service.CreatePaymentMethod(context.Background(), &responsible.CreditCard.CardToken)
	if err != nil {
		t.Errorf("erro ao criar metodo de pagamento")
	}
}

func TestResponsibleUseCase_AttachPaymentMethod(t *testing.T) {
	mock := mocks.NewIResponsibleRepository(t)
	responsible := entity.Responsible{
		Name:  "UpdatePaymentMethodDefault",
		Email: "UpdatePaymentMethodDefault@gmail.com",
		Phone: "+55 11 123456789",
		CreditCard: entity.CreditCard{
			CardToken: "tok_visa",
			Default:   true,
		}}
	service := NewResponsibleUseCase(mock)
	cus, err := service.CreateCustomer(context.Background(), &responsible)
	if err != nil {
		t.Errorf("erro ao criar customer na stripe")
	}
	pm, err := service.CreatePaymentMethod(context.Background(), &responsible.CreditCard.CardToken)
	if err != nil {
		t.Errorf("erro ao criar metodo de pagamento")
	}
	_, err = service.AttachPaymentMethod(context.Background(), &cus.ID, &pm.ID, responsible.CreditCard.Default)
	if err != nil {
		t.Errorf("erro ao atachar metodo de pagamento no customer")
	}
}
