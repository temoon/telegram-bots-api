package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type GetForumTopicIconStickers struct {
}

func (r *GetForumTopicIconStickers) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.Sticker)
	err = b.CallMethod(ctx, "getForumTopicIconStickers", nil, response)
	return
}

func (r *GetForumTopicIconStickers) GetValues() (values map[string]interface{}, err error) {

	return
}

func (r *GetForumTopicIconStickers) GetFiles() (files map[string]io.Reader) {
	return
}
