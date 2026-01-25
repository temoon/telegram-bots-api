package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type GetStarTransactions struct {
	Limit  *int64
	Offset *int64
}

func (r *GetStarTransactions) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.StarTransactions)
	err = b.CallMethod(ctx, "getStarTransactions", r, response)
	return
}

func (r *GetStarTransactions) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	if r.Limit != nil {
		values["limit"] = strconv.FormatInt(*r.Limit, 10)
	}

	if r.Offset != nil {
		values["offset"] = strconv.FormatInt(*r.Offset, 10)
	}

	return
}

func (r *GetStarTransactions) GetFiles() (files map[string]io.Reader) {
	return
}
