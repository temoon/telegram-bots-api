package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type DeleteChatStickerSet struct {
	ChatId telegram.ChatId
}

func (r *DeleteChatStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteChatStickerSet", r, response)
	return
}

func (r *DeleteChatStickerSet) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *DeleteChatStickerSet) GetFiles() (files map[string]io.Reader) {
	return
}
