package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type DeleteChatStickerSet struct {
	ChatId interface{}
}

func (r *DeleteChatStickerSet) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteChatStickerSet", r, response)
	return
}

func (r *DeleteChatStickerSet) IsMultipart() bool {
	return false
}

func (r *DeleteChatStickerSet) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	default:
		err = errors.New("invalid chat_id field type")
		return
	}

	return
}
