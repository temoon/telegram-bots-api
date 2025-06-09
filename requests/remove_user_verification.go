package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type RemoveUserVerification struct {
	UserId int64
}

func (r *RemoveUserVerification) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "removeUserVerification", r, response)
	return
}

func (r *RemoveUserVerification) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *RemoveUserVerification) GetFiles() (files map[string]io.Reader) {
	return
}
