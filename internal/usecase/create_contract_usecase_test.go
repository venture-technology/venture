package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/mocks"
)

func TestCreateContractUsecase_CreateContract(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *value.CreateContractParams
	}
	type fields struct {
		repositories *persistence.PostgresRepositories
	}
	tests := []struct {
		name    string
		setup   func(t *testing.T) *CreateContractUsecase
		args    args
		fields  fields
		wantErr bool
	}{
		{
			name: "there is successfull",
			args: args{
				ctx:     context.Background(),
				payload: &value.CreateContractParams{},
			},
			setup: func(t *testing.T) *CreateContractUsecase {
				kr := mocks.NewKidRepository(t)
				dr := mocks.NewDriverRepository(t)
				rr := mocks.NewResponsibleRepository(t)
				sr := mocks.NewSchoolRepository(t)

				kr.On("FindByResponsible", mock.Anything, mock.Anything).Return(true, nil)
				kr.On("Get", mock.Anything).Return(&entity.Kid{}, nil)
				dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
				sr.On("Get", mock.Anything).Return(&entity.School{}, nil)
				rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)

				worker := mocks.NewWorkerCreateContract(t)

				worker.On("Enqueue", mock.Anything).Return(nil)

				logger := mocks.NewLogger(t)
				return NewCreateContractUsecase(
					&persistence.PostgresRepositories{
						KidRepository:         kr,
						DriverRepository:      dr,
						SchoolRepository:      sr,
						ResponsibleRepository: rr,
					},
					worker,
					logger,
				)
			},
			wantErr: false,
		},
		{
			name: "there is kid find by responsible error",
			args: args{
				ctx:     context.Background(),
				payload: &value.CreateContractParams{},
			},
			setup: func(t *testing.T) *CreateContractUsecase {
				kr := mocks.NewKidRepository(t)
				dr := mocks.NewDriverRepository(t)
				rr := mocks.NewResponsibleRepository(t)
				sr := mocks.NewSchoolRepository(t)

				kr.On("FindByResponsible", mock.Anything, mock.Anything).Return(true, fmt.Errorf("database error"))
				kr.On("Get", mock.Anything).Return(&entity.Kid{}, nil)
				dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
				sr.On("Get", mock.Anything).Return(&entity.School{}, nil)
				rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)

				worker := mocks.NewWorkerCreateContract(t)

				logger := mocks.NewLogger(t)
				return NewCreateContractUsecase(
					&persistence.PostgresRepositories{
						KidRepository:         kr,
						DriverRepository:      dr,
						SchoolRepository:      sr,
						ResponsibleRepository: rr,
					},
					worker,
					logger,
				)
			},
			wantErr: true,
		},
		{
			name: "there is kid get error",
			args: args{
				ctx:     context.Background(),
				payload: &value.CreateContractParams{},
			},
			setup: func(t *testing.T) *CreateContractUsecase {
				kr := mocks.NewKidRepository(t)
				dr := mocks.NewDriverRepository(t)
				rr := mocks.NewResponsibleRepository(t)
				sr := mocks.NewSchoolRepository(t)

				kr.On("FindByResponsible", mock.Anything, mock.Anything).Return(true, nil)
				kr.On("Get", mock.Anything).Return(&entity.Kid{}, fmt.Errorf("database error"))
				dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
				sr.On("Get", mock.Anything).Return(&entity.School{}, nil)
				rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)

				worker := mocks.NewWorkerCreateContract(t)

				logger := mocks.NewLogger(t)
				return NewCreateContractUsecase(
					&persistence.PostgresRepositories{
						KidRepository:         kr,
						DriverRepository:      dr,
						SchoolRepository:      sr,
						ResponsibleRepository: rr,
					},
					worker,
					logger,
				)
			},
			wantErr: true,
		},
		{
			name: "there is driver get error",
			args: args{
				ctx:     context.Background(),
				payload: &value.CreateContractParams{},
			},
			setup: func(t *testing.T) *CreateContractUsecase {
				kr := mocks.NewKidRepository(t)
				dr := mocks.NewDriverRepository(t)
				rr := mocks.NewResponsibleRepository(t)
				sr := mocks.NewSchoolRepository(t)

				kr.On("FindByResponsible", mock.Anything, mock.Anything).Return(true, nil)
				kr.On("Get", mock.Anything).Return(&entity.Kid{}, nil)
				dr.On("Get", mock.Anything).Return(&entity.Driver{}, fmt.Errorf("database error"))
				sr.On("Get", mock.Anything).Return(&entity.School{}, nil)
				rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)

				worker := mocks.NewWorkerCreateContract(t)

				logger := mocks.NewLogger(t)
				return NewCreateContractUsecase(
					&persistence.PostgresRepositories{
						KidRepository:         kr,
						DriverRepository:      dr,
						SchoolRepository:      sr,
						ResponsibleRepository: rr,
					},
					worker,
					logger,
				)
			},
			wantErr: true,
		},
		{
			name: "there is responsible get error",
			args: args{
				ctx:     context.Background(),
				payload: &value.CreateContractParams{},
			},
			setup: func(t *testing.T) *CreateContractUsecase {
				kr := mocks.NewKidRepository(t)
				dr := mocks.NewDriverRepository(t)
				rr := mocks.NewResponsibleRepository(t)
				sr := mocks.NewSchoolRepository(t)

				kr.On("FindByResponsible", mock.Anything, mock.Anything).Return(true, nil)
				kr.On("Get", mock.Anything).Return(&entity.Kid{}, nil)
				dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
				sr.On("Get", mock.Anything).Return(&entity.School{}, nil)
				rr.On("Get", mock.Anything).Return(&entity.Responsible{}, fmt.Errorf("database error"))

				worker := mocks.NewWorkerCreateContract(t)

				logger := mocks.NewLogger(t)
				return NewCreateContractUsecase(
					&persistence.PostgresRepositories{
						KidRepository:         kr,
						DriverRepository:      dr,
						SchoolRepository:      sr,
						ResponsibleRepository: rr,
					},
					worker,
					logger,
				)
			},
			wantErr: true,
		},
		{
			name: "there is school get error",
			args: args{
				ctx:     context.Background(),
				payload: &value.CreateContractParams{},
			},
			setup: func(t *testing.T) *CreateContractUsecase {
				kr := mocks.NewKidRepository(t)
				dr := mocks.NewDriverRepository(t)
				rr := mocks.NewResponsibleRepository(t)
				sr := mocks.NewSchoolRepository(t)

				kr.On("FindByResponsible", mock.Anything, mock.Anything).Return(true, nil)
				kr.On("Get", mock.Anything).Return(&entity.Kid{}, nil)
				dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
				sr.On("Get", mock.Anything).Return(&entity.School{}, fmt.Errorf("database error"))
				rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)

				worker := mocks.NewWorkerCreateContract(t)

				logger := mocks.NewLogger(t)
				return NewCreateContractUsecase(
					&persistence.PostgresRepositories{
						KidRepository:         kr,
						DriverRepository:      dr,
						SchoolRepository:      sr,
						ResponsibleRepository: rr,
					},
					worker,
					logger,
				)
			},
			wantErr: true,
		},
		{
			name: "there is error returned because responsible and kid arent parents",
			args: args{
				ctx:     context.Background(),
				payload: &value.CreateContractParams{},
			},
			setup: func(t *testing.T) *CreateContractUsecase {
				kr := mocks.NewKidRepository(t)
				dr := mocks.NewDriverRepository(t)
				rr := mocks.NewResponsibleRepository(t)
				sr := mocks.NewSchoolRepository(t)

				kr.On("FindByResponsible", mock.Anything, mock.Anything).Return(false, nil)
				kr.On("Get", mock.Anything).Return(&entity.Kid{}, nil)
				dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
				sr.On("Get", mock.Anything).Return(&entity.School{}, fmt.Errorf("parents are false"))
				rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)

				worker := mocks.NewWorkerCreateContract(t)

				logger := mocks.NewLogger(t)
				return NewCreateContractUsecase(
					&persistence.PostgresRepositories{
						KidRepository:         kr,
						DriverRepository:      dr,
						SchoolRepository:      sr,
						ResponsibleRepository: rr,
					},
					worker,
					logger,
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := tt.setup(t)
			err := usecase.CreateContract(tt.args.ctx, tt.args.payload)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
