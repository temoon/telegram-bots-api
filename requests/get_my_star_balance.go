package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetMyStarBalance struct {
}

func (r *GetMyStarBalance) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.StarAmount)
	err = b.CallMethod(ctx, "getMyStarBalance", r, response)
	return
}

func (r *GetMyStarBalance) GetValues() (values map[string]string, err error) {
	return
}

func (r *GetMyStarBalance) GetFiles() (files map[string]io.Reader) {
	return
}
