// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/kat6123/tournament/model"

// TourProvider is an autogenerated mock type for the TourProvider type
type TourProvider struct {
	mock.Mock
}

// CreateTournament provides a mock function with given fields: t
func (_m *TourProvider) CreateTournament(t *model.Tournament) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Tournament) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadTournament provides a mock function with given fields: tourID
func (_m *TourProvider) LoadTournament(tourID int) (*model.Tournament, error) {
	ret := _m.Called(tourID)

	var r0 *model.Tournament
	if rf, ok := ret.Get(0).(func(int) *model.Tournament); ok {
		r0 = rf(tourID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Tournament)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(tourID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveTournament provides a mock function with given fields: t
func (_m *TourProvider) SaveTournament(t *model.Tournament) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Tournament) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
