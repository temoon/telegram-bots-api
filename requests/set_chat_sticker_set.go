package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetChatStickerSet struct {
	StickerSetName string
	ChatId         telegram.ChatId
}

func (r *SetChatStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatStickerSet", r, response)
	return
}

func (r *SetChatStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["sticker_set_name"] = r.StickerSetName

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *SetChatStickerSet) GetFiles() (files map[string]io.Reader) {
	return
}
