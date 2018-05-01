package testdatagen

import (
	"log"

	"github.com/gobuffalo/pop"

	"github.com/transcom/mymove/pkg/models"
)

// MakeDocument creates a single Document.
func MakeDocument(db *pop.Connection, serviceMember *models.ServiceMember) (models.Document, error) {
	if serviceMember == nil {
		newServiceMember, err := MakeServiceMember(db)
		if err != nil {
			log.Panic(err)
		}
		serviceMember = &newServiceMember
	}

	document := models.Document{
		UploaderID:      serviceMember.UserID,
		ServiceMemberID: serviceMember.ID,
		ServiceMember:   *serviceMember,
	}

	verrs, err := db.ValidateAndSave(&document)
	if err != nil {
		log.Panic(err)
	}
	if verrs.Count() != 0 {
		log.Panic(verrs.Error())
	}

	return document, err
}
