package responsible

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/mocks"
)

func TestCreate(t *testing.T) {
	mock := mocks.NewIResponsibleRepository(t)

	mock.On("Create", context.Background(), &entity.Responsible{Password: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"}).Return(nil)

	service := NewResponsibleUseCase(mock)

	err := service.Create(context.Background(), &entity.Responsible{})

	assert.Nil(t, err)

	mock.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}

func TestSaveCard(t *testing.T) {

}

func TestAuth(t *testing.T) {

}

func TestUpdatePaymentMethod(t *testing.T) {

}

func TestCreateCustomer(t *testing.T) {

}

func TestUpdateCustomer(t *testing.T) {

}

func TestDeleteCustomer(t *testing.T) {

}

func TestCreatePaymentMethod(t *testing.T) {

}

func TestAttachPaymentMethod(t *testing.T) {

}
