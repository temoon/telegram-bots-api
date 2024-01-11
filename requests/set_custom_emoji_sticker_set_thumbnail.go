package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
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

func (r *SetCustomEmojiStickerSetThumbnail) IsMultipart() bool {
	return false
}

func (r *SetCustomEmojiStickerSetThumbnail) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.CustomEmojiId != nil {
		values["custom_emoji_id"] = *r.CustomEmojiId
	}

	values["name"] = r.Name

	return
}
