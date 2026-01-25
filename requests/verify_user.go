package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
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

func (r *VerifyUser) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.CustomDescription != nil {
		values["custom_description"] = *r.CustomDescription
	}

	return
}

func (r *VerifyUser) GetFiles() (files map[string]io.Reader) {
	return
}
