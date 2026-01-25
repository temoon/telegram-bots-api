package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type SetBusinessAccountUsername struct {
	BusinessConnectionId string
	Username             *string
}

func (r *SetBusinessAccountUsername) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setBusinessAccountUsername", r, response)
	return
}

func (r *SetBusinessAccountUsername) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	if r.Username != nil {
		values["username"] = *r.Username
	}

	return
}

func (r *SetBusinessAccountUsername) GetFiles() (files map[string]io.Reader) {
	return
}
