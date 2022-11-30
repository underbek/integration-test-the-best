//go:build integration

// для go test ./... --tags=integration

package integration_tests

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/underbek/integration-test-the-best/user-service/internal/app"
	"github.com/underbek/integration-test-the-best/user-service/internal/config"
)

const billingDSN = "http://localhost:8085"

type TestSuite struct {
	suite.Suite
	app    *app.App
	server *httptest.Server
}

func (s *TestSuite) SetupSuite() {
	_, ctxCancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer ctxCancel()

	var err error
	s.app, err = app.New(config.Config{})
	s.Require().NoError(err)

	s.server = httptest.NewServer(s.app.HttpServer.Handler)
}

func (s *TestSuite) TearDownSuite() {
	_, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()
}

func TestSuite_Run(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
