package requests

import (
	"context"

	"github.com/temoon/telegram-bots-api"
)

type Close struct {
}

func (r *Close) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "close", nil, response)
	return
}

func (r *Close) IsMultipart() (multipart bool) {
	return
}

func (r *Close) GetValues() (values map[string]interface{}, err error) {
	return
}
