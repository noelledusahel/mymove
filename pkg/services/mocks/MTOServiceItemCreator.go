// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	validate "github.com/gobuffalo/validate"
)

// MTOServiceItemCreator is an autogenerated mock type for the MTOServiceItemCreator type
type MTOServiceItemCreator struct {
	mock.Mock
}

// CreateMTOServiceItem provides a mock function with given fields: serviceItem
func (_m *MTOServiceItemCreator) CreateMTOServiceItem(serviceItem *models.MtoServiceItem) (*models.MtoServiceItem, *validate.Errors, error) {
	ret := _m.Called(serviceItem)

	var r0 *models.MtoServiceItem
	if rf, ok := ret.Get(0).(func(*models.MtoServiceItem) *models.MtoServiceItem); ok {
		r0 = rf(serviceItem)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MtoServiceItem)
		}
	}

	var r1 *validate.Errors
	if rf, ok := ret.Get(1).(func(*models.MtoServiceItem) *validate.Errors); ok {
		r1 = rf(serviceItem)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*validate.Errors)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*models.MtoServiceItem) error); ok {
		r2 = rf(serviceItem)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
