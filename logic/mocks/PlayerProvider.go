// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/kat6123/tournament/model"

// PlayerProvider is an autogenerated mock type for the PlayerProvider type
type PlayerProvider struct {
	mock.Mock
}

// ByID provides a mock function with given fields: playerID
func (_m *PlayerProvider) ByID(playerID int) (*model.Player, error) {
	ret := _m.Called(playerID)

	var r0 *model.Player
	if rf, ok := ret.Get(0).(func(int) *model.Player); ok {
		r0 = rf(playerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Player)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(playerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: p
func (_m *PlayerProvider) Save(p *model.Player) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Player) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
