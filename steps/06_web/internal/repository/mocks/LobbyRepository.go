// Code generated by mockery v2.33.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/rezaAmiri123/kingscomp/steps/06_web/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// LobbyRepository is an autogenerated mock type for the LobbyRepository type
type LobbyRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, id
func (_m *LobbyRepository) Get(ctx context.Context, id entity.ID) (entity.Lobby, error) {
	ret := _m.Called(ctx, id)

	var r0 entity.Lobby
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.ID) (entity.Lobby, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.ID) entity.Lobby); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Lobby)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.ID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MGet provides a mock function with given fields: ctx, ids
func (_m *LobbyRepository) MGet(ctx context.Context, ids ...entity.ID) ([]entity.Lobby, error) {
	_va := make([]interface{}, len(ids))
	for _i := range ids {
		_va[_i] = ids[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []entity.Lobby
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...entity.ID) ([]entity.Lobby, error)); ok {
		return rf(ctx, ids...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...entity.ID) []entity.Lobby); ok {
		r0 = rf(ctx, ids...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Lobby)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...entity.ID) error); ok {
		r1 = rf(ctx, ids...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, ent
func (_m *LobbyRepository) Save(ctx context.Context, ent entity.Lobby) error {
	ret := _m.Called(ctx, ent)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Lobby) error); ok {
		r0 = rf(ctx, ent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewLobbyRepository creates a new instance of LobbyRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLobbyRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *LobbyRepository {
	mock := &LobbyRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
