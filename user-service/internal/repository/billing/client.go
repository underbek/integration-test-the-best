package billing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"
	"gopkg.in/resty.v1"
)

type client struct {
	client *resty.Client
	addr   string
}

func New(hc *http.Client, addr string) *client {
	return &client{
		client: resty.NewWithClient(hc),
		addr:   addr,
	}
}

func (c *client) UserDeposit(ctx context.Context, userID int, amount decimal.Decimal) (decimal.Decimal, error) {
	request := UserDepositRequest{
		ID:     userID,
		Amount: amount,
	}

	depositResponse := &UserDepositResponse{}

	res, err := c.client.R().
		SetContext(ctx).
		SetBody(request).
		SetResult(depositResponse).
		Post(c.addr + "/deposit")

	if err != nil {
		return decimal.Zero, err
	}

	if res.IsError() {
		return decimal.Zero, fmt.Errorf("code %d, response: %s", res.StatusCode(), res.String())
	}

	return depositResponse.Amount, nil
}
