## Последовательность
1. Показать что из себя пребставляет сервис
2. Запустить и показать, что он заработал
3. suite_test.go
4. Зачатки тестов
5. Пишем TestCreateUser
6. Пытаемся запустить
7. 

## Вставки

### httpmock
```go
	httpmock.RegisterResponder(
		http.MethodPost,
		billingAddr+"/deposit",
		httpmock.NewStringResponder(http.StatusOK, ""),
	)
```

### Фикстуры
```go

```

```go
requestBody := s.loader.LoadString("fixtures/api/deposit_user_request.json")
expected := s.loader.LoadString("fixtures/api/deposit_user_response.json")
```

```go
bytes.NewBufferString(requestBody)
```
