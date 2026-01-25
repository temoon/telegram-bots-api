package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type Close struct {
}

func (r *Close) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "close", r, response)
	return
}

func (r *Close) GetValues() (values map[string]string, err error) {
	return
}

func (r *Close) GetFiles() (files map[string]io.Reader) {
	return
}
