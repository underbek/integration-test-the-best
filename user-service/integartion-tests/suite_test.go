//go:build integration

// для go test ./... --tags=integration

package integration_tests

import (
	"context"
	"database/sql"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"github.com/underbek/integration-test-the-best/testutils"
	"github.com/underbek/integration-test-the-best/testutils/fixtureloader"
	"github.com/underbek/integration-test-the-best/testutils/testcontainer"
	"github.com/underbek/integration-test-the-best/user-service/internal/app"
	"github.com/underbek/integration-test-the-best/user-service/internal/config"
)

const billingDSN = "http://localhost:8085"

type TestSuite struct {
	suite.Suite
	app               *app.App
	server            *httptest.Server
	postgresContainer *testcontainer.PostgresContainer
	loader            *fixtureloader.Loader
}

func (s *TestSuite) SetupSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer ctxCancel()

	var err error

	s.postgresContainer, err = testcontainer.NewPostgresContainer(ctx)
	s.Require().NoError(err)

	s.app, err = app.New(config.Config{
		PostgresDSN: s.postgresContainer.GetDSN(),
		BillingDSN:  billingDSN,
	})
	s.Require().NoError(err)

	s.server = httptest.NewServer(s.app.HttpServer.Handler)

	s.loader = fixtureloader.NewLoader(testutils.Fixtures)

	httpmock.ActivateNonDefault(s.app.BillingHttpClient)
}

func (s *TestSuite) SetupTest() {
	db, err := sql.Open("postgres", s.postgresContainer.GetDSN())
	s.Require().NoError(err)

	defer db.Close()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.FS(testutils.Fixtures),
		testfixtures.Directory("fixtures/storage"),
	)
	s.Require().NoError(err)
	s.Require().NoError(fixtures.Load())
}

func (s *TestSuite) TearDownSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	err := s.postgresContainer.Terminate(ctx)
	s.Require().NoError(err)

	httpmock.DeactivateAndReset()
}

func TestSuite_Run(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
