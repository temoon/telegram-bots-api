package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
)

type SetStickerKeywords struct {
	Keywords []string
	Sticker  string
}

func (r *SetStickerKeywords) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setStickerKeywords", r, response)
	return
}

func (r *SetStickerKeywords) IsMultipart() (multipart bool) {
	return false
}

func (r *SetStickerKeywords) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	var dataKeywords []byte
	if dataKeywords, err = json.Marshal(r.Keywords); err != nil {
		return
	}

	values["keywords"] = string(dataKeywords)

	values["sticker"] = r.Sticker

	return
}
