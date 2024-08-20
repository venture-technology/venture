package child

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/mocks"
)

func TestChildUseCase_Create(t *testing.T) {
	mock := mocks.NewIChildRepository(t)
	mock.On("Create", context.Background(), &entity.Child{}).Return(nil)
	service := NewChildUseCase(mock)
	err := service.Create(context.Background(), &entity.Child{})
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}

func TestChildUseCase_Get(t *testing.T) {
	mock := mocks.NewIChildRepository(t)
	rg := "274020361"
	mock.On("Get", context.Background(), &rg).Return(&entity.Child{}, nil)
	service := NewChildUseCase(mock)
	_, err := service.Get(context.Background(), &rg)
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}

func TestChildUseCase_FindAll(t *testing.T) {
	mock := mocks.NewIChildRepository(t)
	cpf := "22216252000"
	mock.On("FindAll", context.Background(), &cpf).Return([]entity.Child{}, nil)
	service := NewChildUseCase(mock)
	_, err := service.FindAll(context.Background(), &cpf)
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}

func TestChildUseCase_Update(t *testing.T) {
	mock := mocks.NewIChildRepository(t)
	mock.On("Update", context.Background(), &entity.Child{}).Return(nil)
	service := NewChildUseCase(mock)
	err := service.Update(context.Background(), &entity.Child{})
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}

func TestChildUseCase_Delete(t *testing.T) {
	mock := mocks.NewIChildRepository(t)
	rg := "274020361"
	mock.On("Delete", context.Background(), &rg).Return(nil)
	service := NewChildUseCase(mock)
	err := service.Delete(context.Background(), &rg)
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}
