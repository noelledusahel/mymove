package models

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// ReContractYear represents a single "year" of a contract
type ReContractYear struct {
	ID         uuid.UUID `json:"id" db:"id"`
	ContractID uuid.UUID `json:"contract_id" db:"contract_id"`
	Name       string    `json:"name" db:"name"`
	StartDate  time.Time `json:"start_date" db:"start_date"`
	EndDate    time.Time `json:"end_date" db:"end_date"`
	Escalation float64   `json:"escalation" db:"escalation"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`

	// Associations
	ReContract ReContract `belongs_to:"re_contract"`
}

// ReContractYears is not required by pop and may be deleted
type ReContractYears []ReContractYear

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *ReContractYear) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: r.ContractID, Name: "ContractID"},
		&validators.StringIsPresent{Field: r.Name, Name: "Name"},
		&validators.TimeIsPresent{Field: r.StartDate, Name: "StartDate"},
		&validators.TimeIsPresent{Field: r.EndDate, Name: "EndDate"},
		&validators.TimeAfterTime{FirstTime: r.EndDate, FirstName: "EndDate", SecondTime: r.StartDate, SecondName: "StartDate"},
		&Float64IsGreaterThan{Field: r.Escalation, Name: "Escalation", Compared: 0},
	), nil
}
