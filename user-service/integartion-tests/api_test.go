//go:build integration

package integration_tests

import (
	"encoding/json"
	"net/http"

	"github.com/jarcoal/httpmock"
	"github.com/underbek/integration-test-the-best/testutils"
	"github.com/underbek/integration-test-the-best/testutils/fixtureloader"
	"github.com/underbek/integration-test-the-best/user-service/api"
	"github.com/underbek/integration-test-the-best/user-service/internal/repository/billing"
)

func (s *TestSuite) TestCreateUser() {
	// создаем реквест
	buf := s.loader.LoadFile(s.T(), "fixtures/api/create_user_request.json")

	// вызываем апи создания пользователя
	res, err := s.server.Client().Post(s.server.URL+"/users", "", buf)
	s.Require().NoError(err)

	defer res.Body.Close()

	s.Require().Equal(http.StatusOK, res.StatusCode)

	// получаем response
	response := api.CreateUserResponse{}
	err = json.NewDecoder(res.Body).Decode(&response)
	s.Require().NoError(err)

	// проверяем
	expected := s.loader.LoadTemplate(s.T(), "fixtures/api/create_user_response.json.temp", map[string]interface{}{
		"id": response.ID,
	})
	testutils.JSONEq(s.T(), expected, response)
}

func (s *TestSuite) TestGetUser() {
	// вызываем апи получения пользователя
	res, err := s.server.Client().Get(s.server.URL + "/users/1")
	s.Require().NoError(err)

	defer res.Body.Close()

	s.Require().Equal(http.StatusOK, res.StatusCode)

	// проверяем
	expected := s.loader.LoadFile(s.T(), "fixtures/api/get_user_response.json")
	testutils.JSONEq(s.T(), expected, res.Body)
}

func (s *TestSuite) TestDepositBalance() {
	// создаем реквест
	buf := s.loader.LoadFile(s.T(), "fixtures/api/deposit_user_request.json")

	billingResponse := fixtureloader.LoadAPIFixture[billing.UserDepositResponse](
		s.T(), s.loader, "fixtures/billing/user_deposit_response.json",
	)

	responder, err := httpmock.NewJsonResponder(http.StatusOK, billingResponse)
	s.Require().NoError(err)

	httpmock.RegisterResponder(
		http.MethodPost,
		billingDSN+"/deposit",
		responder,
	)

	// вызываем апи депозита
	res, err := s.server.Client().Post(s.server.URL+"/users/deposit", "", buf)
	s.Require().NoError(err)

	defer res.Body.Close()

	s.Require().Equal(http.StatusOK, res.StatusCode)

	expected := s.loader.LoadFile(s.T(), "fixtures/api/deposit_user_response.json")
	testutils.JSONEq(s.T(), expected, res.Body)
}
