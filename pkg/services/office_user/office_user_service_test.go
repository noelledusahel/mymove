package officeuser

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/testingsuite"
)

type OfficeUserServiceSuite struct {
	testingsuite.PopTestSuite
	logger Logger
}

func TestUserSuite(t *testing.T) {

	hs := &OfficeUserServiceSuite{
		PopTestSuite: testingsuite.NewPopTestSuite(testingsuite.CurrentPackage()),
		logger:       zap.NewNop(), // Use a no-op logger during testing
	}
	suite.Run(t, hs)
}
