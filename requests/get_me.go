package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetMe struct {
}

func (r *GetMe) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.User)
	err = b.CallMethod(ctx, "getMe", r, response)
	return
}

func (r *GetMe) GetValues() (values map[string]string, err error) {
	return
}

func (r *GetMe) GetFiles() (files map[string]io.Reader) {
	return
}
