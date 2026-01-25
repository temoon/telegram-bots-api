package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetBusinessAccountStarBalance struct {
	BusinessConnectionId string
}

func (r *GetBusinessAccountStarBalance) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.StarAmount)
	err = b.CallMethod(ctx, "getBusinessAccountStarBalance", r, response)
	return
}

func (r *GetBusinessAccountStarBalance) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["business_connection_id"] = r.BusinessConnectionId

	return
}

func (r *GetBusinessAccountStarBalance) GetFiles() (files map[string]io.Reader) {
	return
}
