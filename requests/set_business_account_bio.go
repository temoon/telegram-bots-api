package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type SetBusinessAccountBio struct {
	BusinessConnectionId string
	Bio                  *string
}

func (r *SetBusinessAccountBio) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setBusinessAccountBio", r, response)
	return
}

func (r *SetBusinessAccountBio) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	if r.Bio != nil {
		values["bio"] = *r.Bio
	}

	return
}

func (r *SetBusinessAccountBio) GetFiles() (files map[string]io.Reader) {
	return
}
