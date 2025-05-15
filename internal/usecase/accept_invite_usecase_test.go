package usecase

import (
	"testing"
)

func TestAcceptInviteUsecase_Accept_Invite(t *testing.T) {
	// id := "123"
	// t.Run("if Invite repository returns error", func(t *testing.T) {
	// 	logger := mocks.NewLogger(t)
	// 	repository := mocks.NewInviteRepository(t)

	// 	usecase := NewAcceptInviteUseCase(
	// 		&persistence.PostgresRepositories{
	// 			InviteRepository: repository,
	// 		},
	// 		logger,
	// 	)
	// 	repository.On("Accept", mock.Anything).Return(errors.New("database error"))

	// 	err := usecase.AcceptInvite(id)

	// 	assert.EqualError(t, err, "database error")
	// })

	// t.Run("when invite repository returns success", func(t *testing.T) {
	// 	logger := mocks.NewLogger(t)
	// 	repository := mocks.NewInviteRepository(t)

	// 	usecase := NewAcceptInviteUseCase(
	// 		&persistence.PostgresRepositories{
	// 			InviteRepository: repository,
	// 		},
	// 		logger,
	// 	)
	// 	repository.On("Accept", mock.Anything).Return(nil)

	// 	err := usecase.AcceptInvite(id)

	// 	assert.Nil(t, err)
	// })

	tests := []struct {
		name    string
		setup   func(t *testing.T) *AcceptInviteUseCase
		wantErr bool
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
		})
	}

}
