// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	mto_shipment "github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/mto_shipment"
)

// MTOShipmentStatusUpdater is an autogenerated mock type for the MTOShipmentStatusUpdater type
type MTOShipmentStatusUpdater struct {
	mock.Mock
}

// UpdateMTOShipmentStatus provides a mock function with given fields: payload
func (_m *MTOShipmentStatusUpdater) UpdateMTOShipmentStatus(payload mto_shipment.PatchMTOShipmentStatusParams) (*models.MTOShipment, error) {
	ret := _m.Called(payload)

	var r0 *models.MTOShipment
	if rf, ok := ret.Get(0).(func(mto_shipment.PatchMTOShipmentStatusParams) *models.MTOShipment); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MTOShipment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mto_shipment.PatchMTOShipmentStatusParams) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
