// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	context "context"

	user "github.com/insaneadinesia/go-boilerplate/internal/app/usecase/user"
	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, req
func (_m *UserUsecase) Create(ctx context.Context, req user.CreateUpdateUserRequest) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, user.CreateUpdateUserRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, reqUUID
func (_m *UserUsecase) Delete(ctx context.Context, reqUUID string) error {
	ret := _m.Called(ctx, reqUUID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, reqUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, req
func (_m *UserUsecase) GetAll(ctx context.Context, req user.GetAllUserRequest) (user.GetAllUserResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 user.GetAllUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, user.GetAllUserRequest) (user.GetAllUserResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, user.GetAllUserRequest) user.GetAllUserResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(user.GetAllUserResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, user.GetAllUserRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetail provides a mock function with given fields: ctx, reqUUID
func (_m *UserUsecase) GetDetail(ctx context.Context, reqUUID string) (user.UserDetailResponse, error) {
	ret := _m.Called(ctx, reqUUID)

	if len(ret) == 0 {
		panic("no return value specified for GetDetail")
	}

	var r0 user.UserDetailResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (user.UserDetailResponse, error)); ok {
		return rf(ctx, reqUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) user.UserDetailResponse); ok {
		r0 = rf(ctx, reqUUID)
	} else {
		r0 = ret.Get(0).(user.UserDetailResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, reqUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, reqUUID, req
func (_m *UserUsecase) Update(ctx context.Context, reqUUID string, req user.CreateUpdateUserRequest) error {
	ret := _m.Called(ctx, reqUUID, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, user.CreateUpdateUserRequest) error); ok {
		r0 = rf(ctx, reqUUID, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
