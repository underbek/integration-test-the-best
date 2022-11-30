//go:build integration

package integration_tests

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/underbek/integration-test-the-best/user-service/api"
)

func (s *TestSuite) TestCreateUser() {
	// создаем реквест
	request := api.CreateUserRequest{}
	buf := bytes.NewBufferString("")
	err := json.NewEncoder(buf).Encode(request)
	s.Require().NoError(err)

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
	// TODO: добавить проверки
}

func (s *TestSuite) TestGetUser() {
	// вызываем апи получения пользователя
	res, err := s.server.Client().Get(s.server.URL + "/users/1")
	s.Require().NoError(err)

	defer res.Body.Close()

	s.Require().Equal(http.StatusOK, res.StatusCode)

	// получаем response
	response := api.GetUserResponse{}
	err = json.NewDecoder(res.Body).Decode(&response)
	s.Require().NoError(err)

	// проверяем
	// TODO: добавить проверки
}

func (s *TestSuite) TestDepositBalance() {
	// создаем реквест
	request := api.DepositBalanceRequest{}
	buf := bytes.NewBufferString("")
	err := json.NewEncoder(buf).Encode(request)
	s.Require().NoError(err)

	// вызываем апи депозита
	res, err := s.server.Client().Post(s.server.URL+"/users/deposit", "", buf)
	s.Require().NoError(err)

	defer res.Body.Close()

	s.Require().Equal(http.StatusOK, res.StatusCode)

	// получаем response
	response := api.DepositBalanceResponse{}
	err = json.NewDecoder(res.Body).Decode(&response)
	s.Require().NoError(err)

	// проверяем
	// TODO: добавить проверки
}
