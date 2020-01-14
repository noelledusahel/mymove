// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MoveDocumentPayload move document payload
// swagger:model MoveDocumentPayload
type MoveDocumentPayload struct {

	// document
	// Required: true
	Document *DocumentPayload `json:"document"`

	// Empty Weight
	// Minimum: 0
	EmptyWeight *int64 `json:"empty_weight,omitempty"`

	// missing empty weight ticket
	EmptyWeightTicketMissing *bool `json:"empty_weight_ticket_missing,omitempty"`

	// Full Weight
	// Minimum: 0
	FullWeight *int64 `json:"full_weight,omitempty"`

	// missing full weight ticket
	FullWeightTicketMissing *bool `json:"full_weight_ticket_missing,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// move document type
	// Required: true
	MoveDocumentType MoveDocumentType `json:"move_document_type"`

	// move id
	// Required: true
	// Format: uuid
	MoveID *strfmt.UUID `json:"move_id"`

	// moving expense type
	MovingExpenseType MovingExpenseType `json:"moving_expense_type,omitempty"`

	// Notes
	Notes *string `json:"notes,omitempty"`

	// Payment Method
	// Enum: [OTHER GTCC]
	PaymentMethod string `json:"payment_method,omitempty"`

	// personally procured move id
	// Format: uuid
	PersonallyProcuredMoveID *strfmt.UUID `json:"personally_procured_move_id,omitempty"`

	// missing expense receipt
	ReceiptMissing *bool `json:"receipt_missing,omitempty"`

	// Requested Amount
	//
	// unit is cents
	// Minimum: 1
	RequestedAmountCents int64 `json:"requested_amount_cents,omitempty"`

	// status
	// Required: true
	Status MoveDocumentStatus `json:"status"`

	// End date of storage for storage expenses
	// Format: date
	StorageEndDate *strfmt.Date `json:"storage_end_date,omitempty"`

	// Start date of storage for storage expenses
	// Format: date
	StorageStartDate *strfmt.Date `json:"storage_start_date,omitempty"`

	// Document Title
	// Required: true
	Title *string `json:"title"`

	// missing trailer ownership documentation
	TrailerOwnershipMissing *bool `json:"trailer_ownership_missing,omitempty"`

	// Vehicle nickname (ex. 'My car')
	VehicleNickname string `json:"vehicle_nickname,omitempty"`

	// Weight ticket date
	// Format: date
	WeightTicketDate *strfmt.Date `json:"weight_ticket_date,omitempty"`

	// weight ticket set type
	WeightTicketSetType *WeightTicketSetType `json:"weight_ticket_set_type,omitempty"`
}

// Validate validates this move document payload
func (m *MoveDocumentPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmptyWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveDocumentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMovingExpenseType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonallyProcuredMoveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedAmountCents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeightTicketDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeightTicketSetType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveDocumentPayload) validateDocument(formats strfmt.Registry) error {

	if err := validate.Required("document", "body", m.Document); err != nil {
		return err
	}

	if m.Document != nil {
		if err := m.Document.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("document")
			}
			return err
		}
	}

	return nil
}

func (m *MoveDocumentPayload) validateEmptyWeight(formats strfmt.Registry) error {

	if swag.IsZero(m.EmptyWeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("empty_weight", "body", int64(*m.EmptyWeight), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validateFullWeight(formats strfmt.Registry) error {

	if swag.IsZero(m.FullWeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("full_weight", "body", int64(*m.FullWeight), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validateMoveDocumentType(formats strfmt.Registry) error {

	if err := m.MoveDocumentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("move_document_type")
		}
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validateMoveID(formats strfmt.Registry) error {

	if err := validate.Required("move_id", "body", m.MoveID); err != nil {
		return err
	}

	if err := validate.FormatOf("move_id", "body", "uuid", m.MoveID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validateMovingExpenseType(formats strfmt.Registry) error {

	if swag.IsZero(m.MovingExpenseType) { // not required
		return nil
	}

	if err := m.MovingExpenseType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("moving_expense_type")
		}
		return err
	}

	return nil
}

var moveDocumentPayloadTypePaymentMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OTHER","GTCC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moveDocumentPayloadTypePaymentMethodPropEnum = append(moveDocumentPayloadTypePaymentMethodPropEnum, v)
	}
}

const (

	// MoveDocumentPayloadPaymentMethodOTHER captures enum value "OTHER"
	MoveDocumentPayloadPaymentMethodOTHER string = "OTHER"

	// MoveDocumentPayloadPaymentMethodGTCC captures enum value "GTCC"
	MoveDocumentPayloadPaymentMethodGTCC string = "GTCC"
)

// prop value enum
func (m *MoveDocumentPayload) validatePaymentMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, moveDocumentPayloadTypePaymentMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MoveDocumentPayload) validatePaymentMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentMethodEnum("payment_method", "body", m.PaymentMethod); err != nil {
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validatePersonallyProcuredMoveID(formats strfmt.Registry) error {

	if swag.IsZero(m.PersonallyProcuredMoveID) { // not required
		return nil
	}

	if err := validate.FormatOf("personally_procured_move_id", "body", "uuid", m.PersonallyProcuredMoveID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validateRequestedAmountCents(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedAmountCents) { // not required
		return nil
	}

	if err := validate.MinimumInt("requested_amount_cents", "body", int64(m.RequestedAmountCents), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validateStatus(formats strfmt.Registry) error {

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validateStorageEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("storage_end_date", "body", "date", m.StorageEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validateStorageStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("storage_start_date", "body", "date", m.StorageStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validateWeightTicketDate(formats strfmt.Registry) error {

	if swag.IsZero(m.WeightTicketDate) { // not required
		return nil
	}

	if err := validate.FormatOf("weight_ticket_date", "body", "date", m.WeightTicketDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveDocumentPayload) validateWeightTicketSetType(formats strfmt.Registry) error {

	if swag.IsZero(m.WeightTicketSetType) { // not required
		return nil
	}

	if m.WeightTicketSetType != nil {
		if err := m.WeightTicketSetType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight_ticket_set_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveDocumentPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveDocumentPayload) UnmarshalBinary(b []byte) error {
	var res MoveDocumentPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
