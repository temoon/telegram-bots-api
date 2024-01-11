package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
)

type GetForumTopicIconStickers struct {
}

func (r *GetForumTopicIconStickers) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.Sticker)
	err = b.CallMethod(ctx, "getForumTopicIconStickers", nil, response)
	return
}

func (r *GetForumTopicIconStickers) IsMultipart() bool {
	return false
}

func (r *GetForumTopicIconStickers) GetValues() (values map[string]interface{}, err error) {

	return
}
