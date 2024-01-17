package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetCustomEmojiStickerSetThumbnail struct {
	CustomEmojiId *string
	Name          string
}

func (r *SetCustomEmojiStickerSetThumbnail) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setCustomEmojiStickerSetThumbnail", r, response)
	return
}

func (r *SetCustomEmojiStickerSetThumbnail) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.CustomEmojiId != nil {
		values["custom_emoji_id"] = *r.CustomEmojiId
	}

	values["name"] = r.Name

	return
}

func (r *SetCustomEmojiStickerSetThumbnail) GetFiles() (files map[string]io.Reader) {
	return
}
