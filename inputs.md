## Вставки
### postgres
```go
PostgresDSN: "postgresql://user:password@localhost/postgres?sslmode=disable",
```

### Проверки
```go
	s.Assert().NotEqual(0, response.ID)
	s.Assert().Equal(request.Name, response.Name)
	s.Assert().Equal("0", response.Balance.String())
```

### conatiner
```go
	s.postgresContainer, err = testcontainer.NewPostgresContainer(ctx)
	s.Require().NoError(err)
```

```go
	err := s.postgresContainer.Terminate(ctx)
	s.Require().NoError(err)
```

```go
PostgresDSN: s.postgresContainer.GetDSN(),
```

### DB Фикстуры
```go
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
```

```go
	s.Assert().Equal(1, response.ID)
	s.Assert().Equal("test_name", response.Name)
	s.Assert().Equal("12.81", response.Balance.String())
```

### httpmock
```go
	httpmock.ActivateNonDefault(app.BillingHttpClient)
```

```go
	httpmock.DeactivateAndReset()
```

```go
	responder, err := httpmock.NewJsonResponder(http.StatusOK, billing.UserDepositResponse{
        Amount: decimal.NewFromInt(5),
	})
    s.Require().NoError(err)
    
    httpmock.RegisterResponder(
        http.MethodPost,
        billingDSN+"/deposit",
        responder,
    )
```

```go
	s.Assert().Equal(1, response.ID)
	s.Assert().Equal("test_name", response.Name)
	s.Assert().Equal("17.81", response.Balance.String())
```

### API Фикстуры
```go
	s.loader = fixtureloader.NewLoader(testutils.Fixtures)
```

```go
	billingResponse := fixtureloader.LoadAPIFixture[billing.UserDepositResponse](
		s.T(), s.loader, "fixtures/billing/user_deposit_response.json",
	)
```

```go
	expected := s.loader.LoadFile(s.T(), "fixtures/api/deposit_user_response.json")
	testutils.JSONEq(s.T(), expected, res.Body)
```


```go
	expected := s.loader.LoadTemplate(s.T(), "fixtures/api/create_user_response.json.temp",
        map[string]interface{}{
            "id": response.ID,
        },
    )

    testutils.JSONEq(s.T(), expected, response)
```
