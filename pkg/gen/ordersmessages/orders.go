// Code generated by go-swagger; DO NOT EDIT.

package ordersmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Orders orders
// swagger:model Orders
type Orders struct {

	// Electronic Data Interchange Personal Identifier, AKA the 10 digit DoD ID Number of the member.
	// Required: true
	// Pattern: ^\d{10}$
	Edipi string `json:"edipi"`

	// issuer
	// Required: true
	Issuer Issuer `json:"issuer"`

	// Orders number. Supposed to be unique, but in practice uniqueness is not guaranteed for all branches of service.
	// # Army
	// Typically found in the upper-left hand corner of printed orders. For example, "030-00362". At this time, there is no common format for Orders numbers between Army installations.
	// # Navy
	// Corresponds to the CT (Commercial Travel) SDN. On printed orders, this is found on the SDN line in the `------- ACCOUNTING DATA -------` section in the `PCS ACCOUNTING DATA` paragraph. For example, "N0001234ABC5XYZ".
	// The BUPERS Orders number is not suitable, because it includes the sailor's full SSN, and the included four digit date code could repeat for a sailor if he or she gets orders exactly 10 years apart.
	// No-cost moves do not have a CT SDN, because they involve no travel. Without a CT SDN, USN Orders have nothing to use for the Orders number. Such Orders won't authorize any PCS expenses either, so they do not need to be submitted to this API.
	// # Marine Corps
	// Corresponds to the CT (Commercial Travel) SDN. On Web Orders, the CT SDN is found in the table at the bottom, in the last column of the row that begins with "Travel". For example, "M7000213CTB28DZ".
	// No-cost moves do not have a CT SDN, because they involve no travel. Without a CT SDN, USMC Orders have nothing to use for the Orders number. Such Orders won't authorize any PCS expenses either, so they do not need to be submitted to this API.
	// # Air Force
	// Corresponds to the SPECIAL ORDER NO, found in box 27 on AF Form 899. For example, "AJ-063322".
	// # Coast Guard
	// Corresponds to the Travel Order No. For example, "1214G85PRAAGK000".
	// # Civilian
	// Corresponds to the Travel Authorization Number. For example, "PS8D000025".
	//
	// Required: true
	OrdersNum string `json:"ordersNum"`

	// revisions
	// Required: true
	Revisions []*Revision `json:"revisions"`

	// Universally Unique IDentifier. Generated internally.
	// Format: uuid
	UUID strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this orders
func (m *Orders) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdipi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrdersNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevisions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Orders) validateEdipi(formats strfmt.Registry) error {

	if err := validate.RequiredString("edipi", "body", string(m.Edipi)); err != nil {
		return err
	}

	if err := validate.Pattern("edipi", "body", string(m.Edipi), `^\d{10}$`); err != nil {
		return err
	}

	return nil
}

func (m *Orders) validateIssuer(formats strfmt.Registry) error {

	if err := m.Issuer.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("issuer")
		}
		return err
	}

	return nil
}

func (m *Orders) validateOrdersNum(formats strfmt.Registry) error {

	if err := validate.RequiredString("ordersNum", "body", string(m.OrdersNum)); err != nil {
		return err
	}

	return nil
}

func (m *Orders) validateRevisions(formats strfmt.Registry) error {

	if err := validate.Required("revisions", "body", m.Revisions); err != nil {
		return err
	}

	for i := 0; i < len(m.Revisions); i++ {
		if swag.IsZero(m.Revisions[i]) { // not required
			continue
		}

		if m.Revisions[i] != nil {
			if err := m.Revisions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("revisions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Orders) validateUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Orders) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Orders) UnmarshalBinary(b []byte) error {
	var res Orders
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}