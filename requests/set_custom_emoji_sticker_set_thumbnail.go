package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetCustomEmojiStickerSetThumbnail struct {
	Name          string
	CustomEmojiId *string
}

func (r *SetCustomEmojiStickerSetThumbnail) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setCustomEmojiStickerSetThumbnail", r, response)
	return
}

func (r *SetCustomEmojiStickerSetThumbnail) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["name"] = r.Name

	if r.CustomEmojiId != nil {
		values["custom_emoji_id"] = *r.CustomEmojiId
	}

	return
}

func (r *SetCustomEmojiStickerSetThumbnail) GetFiles() (files map[string]io.Reader) {
	return
}
