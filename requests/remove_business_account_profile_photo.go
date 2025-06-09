package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type RemoveBusinessAccountProfilePhoto struct {
	BusinessConnectionId string
	IsPublic             *bool
}

func (r *RemoveBusinessAccountProfilePhoto) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "removeBusinessAccountProfilePhoto", r, response)
	return
}

func (r *RemoveBusinessAccountProfilePhoto) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	if r.IsPublic != nil {
		if *r.IsPublic {
			values["is_public"] = "1"
		} else {
			values["is_public"] = "0"
		}
	}

	return
}

func (r *RemoveBusinessAccountProfilePhoto) GetFiles() (files map[string]io.Reader) {
	return
}
