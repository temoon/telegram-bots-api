package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type DeleteChatStickerSet struct {
	ChatId telegram.ChatId
}

func (r *DeleteChatStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteChatStickerSet", r, response)
	return
}

func (r *DeleteChatStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *DeleteChatStickerSet) GetFiles() (files map[string]io.Reader) {
	return
}
