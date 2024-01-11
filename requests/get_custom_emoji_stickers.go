package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
)

type GetCustomEmojiStickers struct {
	CustomEmojiIds []string
}

func (r *GetCustomEmojiStickers) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.Sticker)
	err = b.CallMethod(ctx, "getCustomEmojiStickers", r, response)
	return
}

func (r *GetCustomEmojiStickers) IsMultipart() bool {
	return false
}

func (r *GetCustomEmojiStickers) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	var dataCustomEmojiIds []byte
	if dataCustomEmojiIds, err = json.Marshal(r.CustomEmojiIds); err != nil {
		return
	}

	values["custom_emoji_ids"] = string(dataCustomEmojiIds)

	return
}
