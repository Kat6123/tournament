// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/kat6123/tournament/model"

// TourService is an autogenerated mock type for the TourService type
type TourService struct {
	mock.Mock
}

// AnnounceTournament provides a mock function with given fields: tourID, deposit
func (_m *TourService) AnnounceTournament(tourID int, deposit float32) error {
	ret := _m.Called(tourID, deposit)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, float32) error); ok {
		r0 = rf(tourID, deposit)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Balance provides a mock function with given fields: playerID
func (_m *TourService) Balance(playerID int) (float32, error) {
	ret := _m.Called(playerID)

	var r0 float32
	if rf, ok := ret.Get(0).(func(int) float32); ok {
		r0 = rf(playerID)
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(playerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndTournament provides a mock function with given fields: tourID
func (_m *TourService) EndTournament(tourID int) error {
	ret := _m.Called(tourID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(tourID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fund provides a mock function with given fields: playerID, points
func (_m *TourService) Fund(playerID int, points float32) error {
	ret := _m.Called(playerID, points)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, float32) error); ok {
		r0 = rf(playerID, points)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// JoinTournament provides a mock function with given fields: tourID, playerID
func (_m *TourService) JoinTournament(tourID int, playerID int) error {
	ret := _m.Called(tourID, playerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(tourID, playerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResultTournament provides a mock function with given fields: tourID
func (_m *TourService) ResultTournament(tourID int) (*model.Winner, error) {
	ret := _m.Called(tourID)

	var r0 *model.Winner
	if rf, ok := ret.Get(0).(func(int) *model.Winner); ok {
		r0 = rf(tourID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Winner)
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

// Take provides a mock function with given fields: playerID, points
func (_m *TourService) Take(playerID int, points float32) error {
	ret := _m.Called(playerID, points)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, float32) error); ok {
		r0 = rf(playerID, points)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
