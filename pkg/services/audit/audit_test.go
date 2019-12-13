package audit

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/auth"
	"github.com/transcom/mymove/pkg/models"
)

func TestCapture(t *testing.T) {
	uuidString := "77c9922f-58c7-45cd-8c10-48f2a52bb55d"
	officeUserID, _ := uuid.FromString(uuidString)
	model := models.OfficeUser{
		ID:        officeUserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	logger := zap.NewNop()

	t.Run("success", func(t *testing.T) {
		uuidString := "88c9922f-58c7-45cd-8c10-48f2a52bbabc"
		adminUserID, _ := uuid.FromString(uuidString)

		session := auth.Session{
			AdminUserID: adminUserID,
		}

		zapFields, _ := Capture(&model, nil, logger, &session, "create_office_user")
		var eventType string
		for _, field := range zapFields {
			if field.Key == "event_type" {
				eventType = field.String
			}
		}

		if assert.NotEmpty(t, zapFields) {
			assert.Equal(t, "record_id", zapFields[0].Key)
			assert.Contains(t, eventType, "audit_")
		}
	})

	t.Run("success with optional patch payload", func(t *testing.T) {
		uuidString := "88c9922f-58c7-45cd-8c10-48f2a52bbabc"
		adminUserID, _ := uuid.FromString(uuidString)

		type fakePatchPayload struct {
			Active         bool    `json:"active,omitempty"`
			FirstName      string  `json:"first_name,omitempty"`
			LastName       string  `json:"last_name,omitempty"`
			MiddleInitials *string `json:"middle_initials,omitempty"`
			Telephone      string  `json:"telephone,omitempty"`
		}

		payload := fakePatchPayload{
			Active:    true,
			FirstName: "Leo",
			LastName:  "Spaceman",
			Telephone: "800-588-2300",
		}

		session := auth.Session{
			AdminUserID: adminUserID,
		}

		zapFields, _ := Capture(&model, &payload, logger, &session, "create_office_user")

		var fieldsChanged string
		for _, field := range zapFields {
			if field.Key == "fields_changed" {
				fieldsChanged = field.String
			}
		}

		if assert.NotEmpty(t, zapFields) {
			assert.Equal(t, "active,first_name,last_name,telephone", fieldsChanged)
		}
	})

	t.Run("service member session should not include names", func(t *testing.T) {
		uuidString := "88c9922f-58c7-45cd-8c10-48f2a52bbabc"
		serviceMemberID, _ := uuid.FromString(uuidString)

		session := auth.Session{
			ServiceMemberID: serviceMemberID,
		}

		zapFields, _ := Capture(&model, nil, logger, &session, "create_office_user")

		if assert.NotEmpty(t, zapFields) {
			var keys []string
			for _, field := range zapFields {
				keys = append(keys, field.Key)
			}

			assert.NotContains(t, "responsible_user_name", keys)
		}
	})

	t.Run("failure when a non-pointer is passed in", func(t *testing.T) {
		session := auth.Session{}
		_, err := Capture(model, nil, logger, &session, "create_office_user")

		assert.Equal(t, "must pass a pointer to a struct", err.Error())
	})

	t.Run("failure when a non-struct is passed in", func(t *testing.T) {
		session := auth.Session{}
		invalidArg := 5
		_, err := Capture(&invalidArg, nil, logger, &session, "create_office_user")

		assert.Equal(t, "must pass a pointer to a struct", err.Error())
	})
}
