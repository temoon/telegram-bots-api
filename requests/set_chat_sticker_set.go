package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetChatStickerSet struct {
	ChatId         telegram.ChatId
	StickerSetName string
}

func (r *SetChatStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatStickerSet", r, response)
	return
}

func (r *SetChatStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["sticker_set_name"] = r.StickerSetName

	return
}

func (r *SetChatStickerSet) GetFiles() (files map[string]io.Reader) {
	return
}
