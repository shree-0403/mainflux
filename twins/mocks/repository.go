// Code generated by mockery v2.42.3. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	twins "github.com/absmach/magistrala/twins"
	mock "github.com/stretchr/testify/mock"
)

// TwinRepository is an autogenerated mock type for the TwinRepository type
type TwinRepository struct {
	mock.Mock
}

// Remove provides a mock function with given fields: ctx, twinID
func (_m *TwinRepository) Remove(ctx context.Context, twinID string) error {
	ret := _m.Called(ctx, twinID)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, twinID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveAll provides a mock function with given fields: ctx, owner, offset, limit, name, metadata
func (_m *TwinRepository) RetrieveAll(ctx context.Context, owner string, offset uint64, limit uint64, name string, metadata twins.Metadata) (twins.Page, error) {
	ret := _m.Called(ctx, owner, offset, limit, name, metadata)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAll")
	}

	var r0 twins.Page
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, uint64, string, twins.Metadata) (twins.Page, error)); ok {
		return rf(ctx, owner, offset, limit, name, metadata)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, uint64, string, twins.Metadata) twins.Page); ok {
		r0 = rf(ctx, owner, offset, limit, name, metadata)
	} else {
		r0 = ret.Get(0).(twins.Page)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, uint64, uint64, string, twins.Metadata) error); ok {
		r1 = rf(ctx, owner, offset, limit, name, metadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveByAttribute provides a mock function with given fields: ctx, channel, subtopic
func (_m *TwinRepository) RetrieveByAttribute(ctx context.Context, channel string, subtopic string) ([]string, error) {
	ret := _m.Called(ctx, channel, subtopic)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveByAttribute")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]string, error)); ok {
		return rf(ctx, channel, subtopic)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []string); ok {
		r0 = rf(ctx, channel, subtopic)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, channel, subtopic)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveByID provides a mock function with given fields: ctx, twinID
func (_m *TwinRepository) RetrieveByID(ctx context.Context, twinID string) (twins.Twin, error) {
	ret := _m.Called(ctx, twinID)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveByID")
	}

	var r0 twins.Twin
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (twins.Twin, error)); ok {
		return rf(ctx, twinID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) twins.Twin); ok {
		r0 = rf(ctx, twinID)
	} else {
		r0 = ret.Get(0).(twins.Twin)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, twinID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, twin
func (_m *TwinRepository) Save(ctx context.Context, twin twins.Twin) (string, error) {
	ret := _m.Called(ctx, twin)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, twins.Twin) (string, error)); ok {
		return rf(ctx, twin)
	}
	if rf, ok := ret.Get(0).(func(context.Context, twins.Twin) string); ok {
		r0 = rf(ctx, twin)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, twins.Twin) error); ok {
		r1 = rf(ctx, twin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, twin
func (_m *TwinRepository) Update(ctx context.Context, twin twins.Twin) error {
	ret := _m.Called(ctx, twin)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, twins.Twin) error); ok {
		r0 = rf(ctx, twin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTwinRepository creates a new instance of TwinRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTwinRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TwinRepository {
	mock := &TwinRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
