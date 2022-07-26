// Code generated by mockery v2.12.0. DO NOT EDIT.

package server

import (
	models "MEND/pkg/models"
	context "context"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *MockUserRepository) Create(ctx context.Context, user models.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockUserRepository) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockUserRepository) Get(ctx context.Context, id int) (*models.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, user
func (_m *MockUserRepository) Update(ctx context.Context, id int, user models.User) error {
	ret := _m.Called(ctx, id, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, models.User) error); ok {
		r0 = rf(ctx, id, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUserRepository(t testing.TB) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
