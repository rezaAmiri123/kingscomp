// Code generated by mockery v2.33.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/rezaAmiri123/kingscomp/steps/09_game/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// Lobby is an autogenerated mock type for the Lobby type
type Lobby struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, id
func (_m *Lobby) Get(ctx context.Context, id entity.ID) (entity.Lobby, error) {
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
func (_m *Lobby) MGet(ctx context.Context, ids ...entity.ID) ([]entity.Lobby, error) {
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

// MSet provides a mock function with given fields: ctx, ents
func (_m *Lobby) MSet(ctx context.Context, ents ...entity.Lobby) error {
	_va := make([]interface{}, len(ents))
	for _i := range ents {
		_va[_i] = ents[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...entity.Lobby) error); ok {
		r0 = rf(ctx, ents...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: ctx, ent
func (_m *Lobby) Save(ctx context.Context, ent entity.Lobby) error {
	ret := _m.Called(ctx, ent)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Lobby) error); ok {
		r0 = rf(ctx, ent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetField provides a mock function with given fields: ctx, id, fieldName, value
func (_m *Lobby) SetField(ctx context.Context, id entity.ID, fieldName string, value interface{}) error {
	ret := _m.Called(ctx, id, fieldName, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.ID, string, interface{}) error); ok {
		r0 = rf(ctx, id, fieldName, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserState provides a mock function with given fields: ctx, lobbyId, userId, key, val
func (_m *Lobby) UpdateUserState(ctx context.Context, lobbyId string, userId int64, key string, val interface{}) error {
	ret := _m.Called(ctx, lobbyId, userId, key, val)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, string, interface{}) error); ok {
		r0 = rf(ctx, lobbyId, userId, key, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewLobby creates a new instance of Lobby. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLobby(t interface {
	mock.TestingT
	Cleanup(func())
}) *Lobby {
	mock := &Lobby{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
