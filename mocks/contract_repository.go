// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	entity "github.com/venture-technology/venture/internal/entity"

	uuid "github.com/google/uuid"
)

// IContractRepository is an autogenerated mock type for the IContractRepository type
type IContractRepository struct {
	mock.Mock
}

// Cancel provides a mock function with given fields: ctx, id
func (_m *IContractRepository) Cancel(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Cancel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, contract
func (_m *IContractRepository) Create(ctx context.Context, contract *entity.Contract) error {
	ret := _m.Called(ctx, contract)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Contract) error); ok {
		r0 = rf(ctx, contract)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Expired provides a mock function with given fields: ctx, id
func (_m *IContractRepository) Expired(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Expired")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllByCnh provides a mock function with given fields: ctx, cnh
func (_m *IContractRepository) FindAllByCnh(ctx context.Context, cnh *string) ([]entity.Contract, error) {
	ret := _m.Called(ctx, cnh)

	if len(ret) == 0 {
		panic("no return value specified for FindAllByCnh")
	}

	var r0 []entity.Contract
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string) ([]entity.Contract, error)); ok {
		return rf(ctx, cnh)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string) []entity.Contract); ok {
		r0 = rf(ctx, cnh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Contract)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(ctx, cnh)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllByCnpj provides a mock function with given fields: ctx, cnpj
func (_m *IContractRepository) FindAllByCnpj(ctx context.Context, cnpj *string) ([]entity.Contract, error) {
	ret := _m.Called(ctx, cnpj)

	if len(ret) == 0 {
		panic("no return value specified for FindAllByCnpj")
	}

	var r0 []entity.Contract
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string) ([]entity.Contract, error)); ok {
		return rf(ctx, cnpj)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string) []entity.Contract); ok {
		r0 = rf(ctx, cnpj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Contract)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(ctx, cnpj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllByCpf provides a mock function with given fields: ctx, cpf
func (_m *IContractRepository) FindAllByCpf(ctx context.Context, cpf *string) ([]entity.Contract, error) {
	ret := _m.Called(ctx, cpf)

	if len(ret) == 0 {
		panic("no return value specified for FindAllByCpf")
	}

	var r0 []entity.Contract
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string) ([]entity.Contract, error)); ok {
		return rf(ctx, cpf)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string) []entity.Contract); ok {
		r0 = rf(ctx, cpf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Contract)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(ctx, cpf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *IContractRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Contract, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *entity.Contract
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*entity.Contract, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.Contract); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Contract)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSimpleContractByTitle provides a mock function with given fields: ctx, title
func (_m *IContractRepository) GetSimpleContractByTitle(ctx context.Context, title *string) (*entity.Contract, error) {
	ret := _m.Called(ctx, title)

	if len(ret) == 0 {
		panic("no return value specified for GetSimpleContractByTitle")
	}

	var r0 *entity.Contract
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string) (*entity.Contract, error)); ok {
		return rf(ctx, title)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string) *entity.Contract); ok {
		r0 = rf(ctx, title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Contract)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIContractRepository creates a new instance of IContractRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIContractRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IContractRepository {
	mock := &IContractRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}