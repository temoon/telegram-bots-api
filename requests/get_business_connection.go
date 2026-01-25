package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetBusinessConnection struct {
	BusinessConnectionId string
}

func (r *GetBusinessConnection) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.BusinessConnection)
	err = b.CallMethod(ctx, "getBusinessConnection", r, response)
	return
}

func (r *GetBusinessConnection) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	return
}

func (r *GetBusinessConnection) GetFiles() (files map[string]io.Reader) {
	return
}
