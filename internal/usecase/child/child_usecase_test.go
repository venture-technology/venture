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
}

func TestChildUseCase_FindAll(t *testing.T) {

}

func TestChildUseCase_Update(t *testing.T) {

}

func TestChildUseCase_Delete(t *testing.T) {

}
