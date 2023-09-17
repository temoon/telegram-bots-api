package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
)

type SetStickerEmojiList struct {
	EmojiList []string
	Sticker   string
}

func (r *SetStickerEmojiList) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerEmojiList", r, response)
	return
}

func (r *SetStickerEmojiList) IsMultipart() (multipart bool) {
	return false
}

func (r *SetStickerEmojiList) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	var dataEmojiList []byte
	if dataEmojiList, err = json.Marshal(r.EmojiList); err != nil {
		return
	}

	values["emoji_list"] = string(dataEmojiList)

	values["sticker"] = r.Sticker

	return
}
