package requests

import (
	"context"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type SetChatStickerSet struct {
	ChatId         interface{}
	StickerSetName string
}

func (r *SetChatStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatStickerSet", r, response)
	return
}

func (r *SetChatStickerSet) IsMultipart() (multipart bool) {
	return false
}

func (r *SetChatStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["sticker_set_name"] = r.StickerSetName

	return
}
