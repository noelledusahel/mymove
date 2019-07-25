package migrate

import (
	"bufio"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *MigrateSuite) TestExecWithDeleteUsersSQL() {

	// Load the fixture with the sql example
	f, err := os.Open("./fixtures/deleteUsers.sql")
	suite.NoError(err)

	errTransaction := suite.DB().Transaction(func(tx *pop.Connection) error {
		wait := 10 * time.Millisecond
		err := Exec(f, tx, wait)
		suite.NoError(err)
		return err
	})
	suite.NoError(errTransaction)
}

func (suite *MigrateSuite) TestExecWithCopyFromStdinSQL() {

	// Create common TSP
	testdatagen.MakeTSP(suite.DB(), testdatagen.Assertions{
		TransportationServiceProvider: models.TransportationServiceProvider{
			ID: uuid.Must(uuid.FromString("231a7b21-346c-4e94-b6bc-672413733f77")),
		},
	})

	// Create TDLs
	testdatagen.MakeTDL(suite.DB(), testdatagen.Assertions{
		TrafficDistributionList: models.TrafficDistributionList{
			ID: uuid.Must(uuid.FromString("27f1fbeb-090c-4a91-955c-67899de4d6d6")),
		},
	})

	// Load the fixture with the sql example
	f, err := os.Open("./fixtures/copyFromStdin.sql")
	suite.NoError(err)

	errTransaction := suite.DB().Transaction(func(tx *pop.Connection) error {
		wait := 10 * time.Millisecond
		err := Exec(f, tx, wait)
		suite.NoError(err)
		return err
	})
	suite.NoError(errTransaction)
}

func (suite *MigrateSuite) TestExecWithLoopSQL() {

	// Load the fixture with the sql example
	f, err := os.Open("./fixtures/loop.sql")
	suite.NoError(err)

	errTransaction := suite.DB().Transaction(func(tx *pop.Connection) error {
		wait := 10 * time.Millisecond
		err := Exec(f, tx, wait)
		suite.NoError(err)
		return err
	})
	suite.NoError(errTransaction)
}

func (suite *MigrateSuite) TestExecWithUpdateFromSetSQL() {

	// Load the fixture with the sql example
	f, err := os.Open("./fixtures/update_from_set.sql")
	suite.NoError(err)

	errTransaction := suite.DB().Transaction(func(tx *pop.Connection) error {
		wait := 10 * time.Millisecond
		err := Exec(f, tx, wait)
		suite.NoError(err)
		return err
	})
	suite.NoError(errTransaction)
}

func (suite *MigrateSuite) TestExecWithInsertConflictSQL() {

	// Create Transportation Office
	testdatagen.MakeTransportationOffice(suite.DB(), testdatagen.Assertions{
		TransportationOffice: models.TransportationOffice{
			ID: uuid.Must(uuid.FromString("c219d9e5-2659-427d-be33-bf439251b7f3")),
		},
	})

	// Load the fixture with the sql example
	f, err := os.Open("./fixtures/insert_conflict.sql")
	suite.NoError(err)

	errTransaction := suite.DB().Transaction(func(tx *pop.Connection) error {
		wait := 10 * time.Millisecond
		err := Exec(f, tx, wait)
		suite.NoError(err)
		return err
	})
	suite.NoError(errTransaction)
}

func TestExtractCopyDataRows(t *testing.T) {

	// Load the fixture with the sql example
	fixture := "./fixtures/copyMultiFromStdin.sql"
	f, err := os.Open(fixture)
	defer f.Close()
	require.Nil(t, err)

	lines := make(chan string, 1000)
	dropComments := true
	dropSearchPath := true
	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines <- ReadInSQLLine(scanner.Text(), dropComments, dropSearchPath)
		}
		close(lines)
	}()

	wait := 10 * time.Millisecond
	statements := make(chan string, 1000)
	go SplitStatements(lines, statements, wait)

	// Now extract the copy data rows
	res, _ := extractCopyDataRows(statements)

	// length of data rows should be 2 based on the fixture file
	require.Equal(t, len(res), 2)
}
