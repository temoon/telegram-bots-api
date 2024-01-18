package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type LogOut struct {
}

func (r *LogOut) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "logOut", r, response)
	return
}

func (r *LogOut) GetValues() (values map[string]interface{}, err error) {
	return
}

func (r *LogOut) GetFiles() (files map[string]io.Reader) {
	return
}
