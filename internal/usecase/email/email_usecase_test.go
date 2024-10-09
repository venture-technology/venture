package email

import (
	"context"
	"fmt"
	"testing"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/mocks"
)

func TestEmailUseCase_Record(t *testing.T) {

	t.Run("when return sucess", func(t *testing.T) {
		caller := mocks.NewIEmailRepository(t)
		email := entity.Email{
			Body:      "body",
			Subject:   "subject",
			Recipient: "eu@omg.com",
		}
		caller.On("Record", context.Background(), &email).Return(nil)
		useCase := NewEmailUseCase(caller, nil, nil)
		err := useCase.Record(context.Background(), &email)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("when return error", func(t *testing.T) {
		caller := mocks.NewIEmailRepository(t)
		email := entity.Email{
			Body:      "body",
			Subject:   "subject",
			Recipient: "eu@omg.com",
		}
		caller.On("Record", context.Background(), &email).Return(fmt.Errorf("error happens"))
		useCase := NewEmailUseCase(caller, nil, nil)
		err := useCase.Record(context.Background(), &email)
		if err == nil {
			t.Error("Not returned error")
		}
	})

}
