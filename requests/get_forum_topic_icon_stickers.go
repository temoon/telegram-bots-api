package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetForumTopicIconStickers struct {
}

func (r *GetForumTopicIconStickers) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.Sticker)
	err = b.CallMethod(ctx, "getForumTopicIconStickers", r, response)
	return
}

func (r *GetForumTopicIconStickers) GetValues() (values map[string]string, err error) {
	return
}

func (r *GetForumTopicIconStickers) GetFiles() (files map[string]io.Reader) {
	return
}
