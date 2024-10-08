// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	entity "github.com/venture-technology/venture/internal/entity"
)

// IChildRepository is an autogenerated mock type for the IChildRepository type
type IChildRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, child
func (_m *IChildRepository) Create(ctx context.Context, child *entity.Child) error {
	ret := _m.Called(ctx, child)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Child) error); ok {
		r0 = rf(ctx, child)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, rg
func (_m *IChildRepository) Delete(ctx context.Context, rg *string) error {
	ret := _m.Called(ctx, rg)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *string) error); ok {
		r0 = rf(ctx, rg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx, cpf
func (_m *IChildRepository) FindAll(ctx context.Context, cpf *string) ([]entity.Child, error) {
	ret := _m.Called(ctx, cpf)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []entity.Child
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string) ([]entity.Child, error)); ok {
		return rf(ctx, cpf)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string) []entity.Child); ok {
		r0 = rf(ctx, cpf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Child)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(ctx, cpf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindResponsibleByChild provides a mock function with given fields: ctx, rg
func (_m *IChildRepository) FindResponsibleByChild(ctx context.Context, rg *string) (*entity.Responsible, error) {
	ret := _m.Called(ctx, rg)

	if len(ret) == 0 {
		panic("no return value specified for FindResponsibleByChild")
	}

	var r0 *entity.Responsible
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string) (*entity.Responsible, error)); ok {
		return rf(ctx, rg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string) *entity.Responsible); ok {
		r0 = rf(ctx, rg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Responsible)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(ctx, rg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, rg
func (_m *IChildRepository) Get(ctx context.Context, rg *string) (*entity.Child, error) {
	ret := _m.Called(ctx, rg)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *entity.Child
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string) (*entity.Child, error)); ok {
		return rf(ctx, rg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string) *entity.Child); ok {
		r0 = rf(ctx, rg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Child)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(ctx, rg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, child
func (_m *IChildRepository) Update(ctx context.Context, child *entity.Child) error {
	ret := _m.Called(ctx, child)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Child) error); ok {
		r0 = rf(ctx, child)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIChildRepository creates a new instance of IChildRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIChildRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IChildRepository {
	mock := &IChildRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
