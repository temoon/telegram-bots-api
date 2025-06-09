package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type GetBusinessAccountStarBalance struct {
	BusinessConnectionId string
}

func (r *GetBusinessAccountStarBalance) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.StarAmount)
	err = b.CallMethod(ctx, "getBusinessAccountStarBalance", r, response)
	return
}

func (r *GetBusinessAccountStarBalance) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	return
}

func (r *GetBusinessAccountStarBalance) GetFiles() (files map[string]io.Reader) {
	return
}
