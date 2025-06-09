package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type VerifyUser struct {
	UserId            int64
	CustomDescription *string
}

func (r *VerifyUser) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "verifyUser", r, response)
	return
}

func (r *VerifyUser) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.CustomDescription != nil {
		values["custom_description"] = *r.CustomDescription
	}

	return
}

func (r *VerifyUser) GetFiles() (files map[string]io.Reader) {
	return
}
