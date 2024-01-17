package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetStickerEmojiList struct {
	Sticker   string
	EmojiList []string
}

func (r *SetStickerEmojiList) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerEmojiList", r, response)
	return
}

func (r *SetStickerEmojiList) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["sticker"] = r.Sticker

	var dataEmojiList []byte
	if dataEmojiList, err = json.Marshal(r.EmojiList); err != nil {
		return
	}

	values["emoji_list"] = string(dataEmojiList)

	return
}

func (r *SetStickerEmojiList) GetFiles() (files map[string]io.Reader) {
	return
}
