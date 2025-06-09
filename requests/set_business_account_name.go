package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetBusinessAccountName struct {
	BusinessConnectionId string
	FirstName            string
	LastName             *string
}

func (r *SetBusinessAccountName) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setBusinessAccountName", r, response)
	return
}

func (r *SetBusinessAccountName) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	values["first_name"] = r.FirstName

	if r.LastName != nil {
		values["last_name"] = *r.LastName
	}

	return
}

func (r *SetBusinessAccountName) GetFiles() (files map[string]io.Reader) {
	return
}
