package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type Close struct {
}

func (r *Close) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "close", nil, response)
	return
}

func (r *Close) GetValues() (values map[string]interface{}, err error) {

	return
}

func (r *Close) GetFiles() (files map[string]io.Reader) {
	return
}
