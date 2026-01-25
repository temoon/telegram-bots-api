package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetAvailableGifts struct {
}

func (r *GetAvailableGifts) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Gifts)
	err = b.CallMethod(ctx, "getAvailableGifts", r, response)
	return
}

func (r *GetAvailableGifts) GetValues() (values map[string]string, err error) {
	return
}

func (r *GetAvailableGifts) GetFiles() (files map[string]io.Reader) {
	return
}
