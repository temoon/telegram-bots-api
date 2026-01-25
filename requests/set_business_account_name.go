package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
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

func (r *SetBusinessAccountName) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

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
