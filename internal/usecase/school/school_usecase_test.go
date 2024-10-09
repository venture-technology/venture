package school

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/mocks"
)

func TestSchoolUse_Create(t *testing.T) {
	mock := mocks.NewISchoolRepository(t)
	mock.On("Create", context.Background(), &entity.School{}).Return(nil)
	service := NewSchoolUseCase(mock, nil)
	err := service.Create(context.Background(), &entity.School{})
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}

func TestSchoolUse_Get(t *testing.T) {
	mock := mocks.NewISchoolRepository(t)
	cnpj := "98088292000160"
	mock.On("Get", context.Background(), &cnpj).Return(&entity.School{}, nil)
	service := NewSchoolUseCase(mock, nil)
	_, err := service.Get(context.Background(), &cnpj)
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}

func TestSchoolUse_FindAll(t *testing.T) {
	mock := mocks.NewISchoolRepository(t)
	mock.On("FindAll", context.Background()).Return([]entity.School{}, nil)
	service := NewSchoolUseCase(mock, nil)
	_, err := service.FindAll(context.Background())
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}

func TestSchoolUse_Update(t *testing.T) {
	mock := mocks.NewISchoolRepository(t)
	school := entity.School{Name: "adsasd", CNPJ: "98088292000160", Email: "jorge@gmail.com"}
	mock.On("Update", context.Background(), &school).Return(nil)
	service := NewSchoolUseCase(mock, nil)
	err := service.Update(context.Background(), &school)
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}

func TestSchoolUse_Delete(t *testing.T) {
	mock := mocks.NewISchoolRepository(t)
	cnpj := "98088292000160"
	mock.On("Delete", context.Background(), &cnpj).Return(nil)
	service := NewSchoolUseCase(mock, nil)
	err := service.Delete(context.Background(), &cnpj)
	assert.Nil(t, err)
	mock.AssertExpectations(t)
}
