package auth

import (
	"context"
	"testing"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/mocks"
)

func TestAuthUsecase_LoginResponsible(t *testing.T) {
	mock := mocks.NewIResponsibleRepository(t)
	fakeResponsible := entity.Responsible{
		Email: "teste@gmail.com",
	}
	mock.On("FindByEmail", context.Background(), &fakeResponsible.Email).Return(&entity.Responsible{}, nil)
	service := NewAuthUseCase(nil, nil, mock, nil)
	_, err := service.LoginResponsible(context.Background(), &fakeResponsible.Email)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestAuthUsecase_LoginDriver(t *testing.T) {
	mock := mocks.NewIDriverRepository(t)
	fakeDriver := entity.Driver{
		Email: "teste@gmail.com",
	}
	mock.On("FindByEmail", context.Background(), &fakeDriver.Email).Return(&entity.Driver{}, nil)
	service := NewAuthUseCase(nil, mock, nil, nil)
	_, err := service.LoginDriver(context.Background(), &fakeDriver.Email)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestAuthUsecase_LoginSchool(t *testing.T) {
	mock := mocks.NewISchoolRepository(t)
	fakeSchool := entity.School{
		Email: "teste@gmail.com",
	}
	mock.On("FindByEmail", context.Background(), &fakeSchool.Email).Return(&entity.School{}, nil)
	service := NewAuthUseCase(mock, nil, nil, nil)
	_, err := service.LoginSchool(context.Background(), &fakeSchool.Email)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
