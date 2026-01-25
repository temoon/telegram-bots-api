package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type RemoveUserVerification struct {
	UserId int64
}

func (r *RemoveUserVerification) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "removeUserVerification", r, response)
	return
}

func (r *RemoveUserVerification) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *RemoveUserVerification) GetFiles() (files map[string]io.Reader) {
	return
}
