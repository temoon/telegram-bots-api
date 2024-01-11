package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
)

type GetMe struct {
}

func (r *GetMe) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.User)
	err = b.CallMethod(ctx, "getMe", nil, response)
	return
}

func (r *GetMe) IsMultipart() bool {
	return false
}

func (r *GetMe) GetValues() (values map[string]interface{}, err error) {

	return
}
