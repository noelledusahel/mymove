// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	validate "github.com/gobuffalo/validate"
)

// OfficeUserUpdater is an autogenerated mock type for the OfficeUserUpdater type
type OfficeUserUpdater struct {
	mock.Mock
}

// UpdateOfficeUser provides a mock function with given fields: user
func (_m *OfficeUserUpdater) UpdateOfficeUser(user *models.OfficeUser) (*models.OfficeUser, *validate.Errors, error) {
	ret := _m.Called(user)

	var r0 *models.OfficeUser
	if rf, ok := ret.Get(0).(func(*models.OfficeUser) *models.OfficeUser); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OfficeUser)
		}
	}

	var r1 *validate.Errors
	if rf, ok := ret.Get(1).(func(*models.OfficeUser) *validate.Errors); ok {
		r1 = rf(user)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*validate.Errors)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*models.OfficeUser) error); ok {
		r2 = rf(user)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}