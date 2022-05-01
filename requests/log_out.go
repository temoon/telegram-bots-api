package requests

import (
	"context"

	"github.com/temoon/telegram-bots-api"
)

type LogOut struct {
}

func (r *LogOut) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "logOut", nil, response)
	return
}

func (r *LogOut) IsMultipart() (multipart bool) {
	return
}

func (r *LogOut) GetValues() (values map[string]interface{}, err error) {
	return
}
